package configformat

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/repr"
	"testing"
)

var parser = participle.MustBuild(&INI{},
	participle.Lexer(iniLexer),
	participle.Unquote("String"),
)

func TestIni(t *testing.T) {
	testText:=`age = 21
name = "Bob Smith"

[address]
city = "Beverly Hills"
postal_code = 90210
`
	ini := &INI{}
	err := parser.ParseString("", testText, ini)
	if err != nil {
		panic(err)
	}
	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
}

