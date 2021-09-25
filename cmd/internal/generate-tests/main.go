package main

import (
	"bufio"
	"embed"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/butuzov/mirror/internal/rules"
)

//go:embed templates
var content embed.FS

var templates = template.Must(template.ParseFS(content, "templates/*.tmpl"))

func main() {
	testdata := os.Args[len(os.Args)-1]
	if !testdataExistenceCheck(testdata) {
		log.Printf("testdata directory doesn't exists: %s", testdata)
		return
	}

	{ // regexp

		tests := []string{}
		tests = append(tests, generateTests("regexp", rules.RegexpFunctions)...)
		tests = append(tests, generateTests("regexp", rules.RegexpRegexpMethods)...)

		GenerateTestFile(filepath.Join(testdata, "regexp.go"), "regexp", "Regexp", tests)
	}

	{ // strings

		tests := []string{}
		tests = append(tests, generateTests("strings", rules.StringFunctions)...)
		tests = append(tests, generateTests("strings", rules.StringsBuilderMethods)...)

		GenerateTestFile(filepath.Join(testdata, "strings.go"), "strings", "Builder", tests)
	}

	{ // bytes

		tests := []string{}
		tests = append(tests, generateTests("bytes", rules.BytesFunctions)...)
		tests = append(tests, generateTests("bytes", rules.BytesBufferMethods)...)

		GenerateTestFile(filepath.Join(testdata, "bytes.go"), "bytes", "Buffer", tests)
	}
}

func testdataExistenceCheck(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GenerateTestFile(file string, Package string, Struct string, Tests []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	return templates.ExecuteTemplate(w, "file.tmpl", struct {
		Package string
		Struct  string
		Tests   string
	}{
		Package,
		Struct,
		strings.Join(Tests, "\n"),
	})
}
