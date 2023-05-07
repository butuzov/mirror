package main

import (
	"bufio"
	"embed"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/butuzov/mirror"
)

//go:embed templates
var content embed.FS

var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"basename": func(absPath string) string {
		return path.Base(absPath)
	},
}).ParseFS(content, "templates/*.tmpl"))

func main() {
	testdata := os.Args[len(os.Args)-1]
	if !testdataExistenceCheck(testdata) {
		log.Printf("testdata directory doesn't exists: %s", testdata)
		return
	}

	{ // regexp

		tests := []string{}
		tests = append(tests, generateTests("regexp", mirror.RegexpFunctions)...)
		tests = append(tests, generateTests("regexp", mirror.RegexpRegexpMethods)...)

		GenerateTestFile(filepath.Join(testdata, "regexp.go"), "regexp", tests)
	}

	{ // strings

		tests := []string{}
		tests = append(tests, generateTests("strings", mirror.StringFunctions)...)
		tests = append(tests, generateTests("strings", mirror.StringsBuilderMethods)...)

		GenerateTestFile(filepath.Join(testdata, "strings.go"), "strings", tests)
	}

	{ // bytes

		tests := []string{}
		tests = append(tests, generateTests("bytes", mirror.BytesFunctions)...)
		tests = append(tests, generateTests("bytes", mirror.BytesBufferMethods)...)

		GenerateTestFile(filepath.Join(testdata, "bytes.go"), "bytes", tests)
	}
	{ // hash/maphash

		tests := []string{}
		tests = append(tests, generateTests("maphash", mirror.MaphashMethods)...)

		GenerateTestFile(filepath.Join(testdata, "maphash.go"), "hash/maphash", tests)
	}

	{ // unicode/utf8

		tests := []string{}
		tests = append(tests, generateTests("utf8", mirror.UTF8Functions)...)

		GenerateTestFile(filepath.Join(testdata, "utf8.go"), "unicode/utf8", tests)
	}

	{ // bufio

		tests := []string{}
		tests = append(tests, generateTests("bufio", mirror.BufioMethods)...)

		GenerateTestFile(filepath.Join(testdata, "bufio.go"), "bufio", tests)
	}

	{ // bufio

		tests := []string{}
		tests = append(tests, generateTests("os", mirror.OsFileMethods)...)

		GenerateTestFile(filepath.Join(testdata, "os.go"), "os", tests)
	}
}

func testdataExistenceCheck(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GenerateTestFile(file string, pkgName string, Tests []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	return templates.ExecuteTemplate(w, "file.tmpl", struct {
		Package string
		Tests   string
	}{
		pkgName,
		strings.Join(Tests, "\n"),
	})
}
