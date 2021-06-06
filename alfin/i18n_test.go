package alfin

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"testing"
)

// ref: https://phrase.com/blog/posts/internationalization-i18n-go/
func TestEntries(t *testing.T) {
	p := message.NewPrinter(language.Greek)
	p.Printf("Hello World")
	p.Println()
	p.Printf("%d task(s) remaining!", 2)
	p.Println()

	p = message.NewPrinter(language.English)
	p.Printf("Hello World")
	p.Println()
	// 字符串必须完全匹配, 即使只加入'\n'后辍也以原格式输出
	p.Printf("%d task(s) remaining!", 2)
}
