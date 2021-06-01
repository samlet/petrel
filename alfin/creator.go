package alfin

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
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

func (c *Creator) Execute(suffix string) error {
	err := EnsureFile(c.InputPath)
	if err != nil {
		return err
	}
	err = EnsureFile(c.TemplatePath)
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

	t, err := ioutil.ReadFile(c.TemplatePath)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}
	d, err := ioutil.ReadFile(c.InputPath)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}
	generator.GenTemplate(string(d), string(t), f)

	return nil
}
