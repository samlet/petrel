package alfin

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/magefile/mage/sh"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}

type Creator struct {
	InputPath    string
	TemplatePath string
	PackageName  string
	TargetName   string
	Extra        []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExistsPath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func EnsureFile(path string) error {
	f, err := os.Stat(path)
	if err == nil {
		if !f.IsDir() {
			return nil
		}
		return errors.New(path + " is not a valid file")
	}
	return err
}

func EnsureDir(path string) error {
	// Fast path: if we can tell whether path is a directory or file, stop with success or error.
	dir, err := os.Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
		return &os.PathError{Op: "mkdir", Path: path, Err: syscall.ENOTDIR}
	}

	return os.MkdirAll(path, 0755)
}

func WriteContent(file string, text string) error {
	f, err := os.Create(file)
	w := bufio.NewWriter(f)
	bytes, err := w.WriteString(text)
	if err != nil {
		return err
	}
	log.Printf("wrote %d bytes to file %s\n", bytes, file)

	return w.Flush()
}

func PyGen(script string, arg ...string) (string, error) {
	args := append([]string{script}, arg...)
	binpath := fmt.Sprintf("%s/miniconda3/envs/bigdata/bin/python",
		os.Getenv("HOME"))
	cmd := exec.Command(binpath, args...)
	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	errstr := errout.String()
	if errstr != "" {
		logger.Error(errstr)
	}
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (c *Creator) GetTarget(suffix string) string {
	if c.PackageName != "" {
		return filepath.Join(c.PackageName, c.TargetName+suffix)
	} else {
		return c.TargetName + suffix
	}
}

// ReadTemplate read from file path at first, if not exists,
// read template from rice.box
func (c *Creator) ReadTemplate() (string, error) {
	err := EnsureFile(c.TemplatePath)
	//if err != nil {
	//	dir, file:=filepath.Split(c.TemplatePath)
	//	box:=strings.TrimRight(dir, "/")
	//	println(box, file)
	//	// get template from a rice.Box
	//	return rice.MustFindBox(box).String(file)  // err: argument must be a string literal
	//}

	if err != nil {
		return "", err
	}

	t, err := ioutil.ReadFile(c.TemplatePath)
	if err != nil {
		logger.Fatal(err.Error())
		return "", err
	}
	return string(t), nil
}

func (c *Creator) Execute(suffix string) error {
	err := EnsureFile(c.InputPath)
	if err != nil {
		return err
	}


	if c.PackageName != "" {
		err := EnsureDir(c.PackageName)
		if err != nil {
			return err
		}
	}

	generator := NewGenHelper(logger)
	f, err := os.Create(c.GetTarget(suffix))
	if err != nil {
		return err
	}
	defer f.Close()


	d, err := ioutil.ReadFile(c.InputPath)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}

	t,err:=c.ReadTemplate()
	if err != nil {
		return err
	}

	generator.GenTemplate(string(d), t, f)

	return nil
}

type MainEnt struct {
	Name   string   `yaml:"name"`
	Extras []string `yaml:"extras"`
}

type MainEntConf struct {
	MainEnts []MainEnt `yaml:"mainEnts"`
}

func ReadMaintConf(configFile string) MainEntConf {
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v, Error: %v", configFile, err))
	}

	var maintConf MainEntConf
	if err := yaml.Unmarshal(configData, &maintConf); err != nil {
		panic(fmt.Sprintf("Error initializing configuration: %v", err))
	}
	return maintConf
}

func mkdir(dir string){
	err:=os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err)
	}
}

func chdir(dir string){
	err:=os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func CreateMod(modName string){
	//f:=fmt.Sprintf
	modDir:=filepath.Join("modules", modName)
	mkdir(modDir)
	chdir(modDir)
	err:=sh.RunV("ent", "init", "ModEnt")
	if err != nil {
		panic(err)
	}
	mkdir("helper")

	//targetFile:=filepath.Join("ent", "schema", "modent.go")
}

func WriteSchemas(pkg string, onlyWrite bool) {
	//pkg:="workload"
	// write modent
	path:=filepath.Join("modules", pkg, "ent", "schema", "modent.go")
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	mani, err:=GenSchemas(pkg)
	if err != nil {
		panic(err)
	}
	err=mani.GenSchemaWithFiles(f)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
	f.Sync()
	println("write modent to ", path)

	// write helper
	path=filepath.Join("modules", pkg, "helper", pkg+"helper.go")
	f, err = os.Create(path)
	check(err)
	defer f.Close()
	mani.GenHelpers(f)
	f.Sync()
	println("write helper to ", path)

	if !onlyWrite {
		// invoke generate
		genPath := filepath.Join(".", "modules", pkg, "ent")
		println("execute go generate:", genPath)
		sh.RunV("go", "generate", "./"+genPath)
		println("done.")
	}
}

