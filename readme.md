# `mirror` [![Code Coverage](https://coveralls.io/repos/github/butuzov/mirror/badge.svg?branch=main)](https://coveralls.io/github/butuzov/mirror?branch=main) [![build status](https://github.com/butuzov/mirror/actions/workflows/main.yaml/badge.svg?branch=main)]()

`mirror` help to suggest use of alternative methods, in order to avoid additional convertion of `[]byte`/`string`.

## 🇺🇦 PLEASE HELP ME 🇺🇦
Fundrise for scout drone **DJI Matrice 30T** See more details at [butuzov/README.md](https://github.com/butuzov/butuzov/)

## Examples

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
regexp.Match("foo", []byte("foobar1"))
```

</td><td>

```go
regexp.MatchString("foo", "foobar1")
```

</td></tr>
</tbody></table>

## Supported Checks
<table><tr>
<td><code>func (b *bufio.Writer) Write(p []byte) (int, error)</code></td>
<td><code>func (b *bufio.Writer) WriteString(s string) (int, error)</code></td>
</tr>
<tr>
<td><code>func (b *bytes.Buffer) Write(p []byte) (int, error)</code></td>
<td><code>func (b *bytes.Buffer) WriteString(s string) (int, error)</code></td>
</tr>
<tr>
<td><code>func bytes.Compare(a, b []byte) int</code></td>
<td><code>func strings.Compare(a, b string) int</code></td>
</tr>
<tr>
<td><code>func bytes.Contains(b, subslice []byte) bool</code></td>
<td><code>func strings.Contains(s, substr string) bool</code></td>
</tr>
<tr>
<td><code>func bytes.ContainsAny(b []byte, chars string) bool</code></td>
<td><code>func strings.ContainsAny(s, chars string) bool</code></td>
</tr>
<tr>
<td><code>func bytes.ContainsRune(b []byte, r rune) bool</code></td>
<td><code>func strings.ContainsRune(s string, r rune) bool</code></td>
</tr>
<tr>
<td><code>func bytes.Count(s, sep []byte) int</code></td>
<td><code>func strings.Count(s, substr string) int</code></td>
</tr>
<tr>
<td><code>func bytes.EqualFold(s, t []byte) bool</code></td>
<td><code>func strings.EqualFold(s, t string) bool</code></td>
</tr>
<tr>
<td><code>func bytes.HasPrefix(s, prefix []byte) bool</code></td>
<td><code>func strings.HasPrefix(s, prefix string) bool</code></td>
</tr>
<tr>
<td><code>func bytes.HasSuffix(s, suffix []byte) bool</code></td>
<td><code>func strings.HasSuffix(s, suffix string) bool</code></td>
</tr>
<tr>
<td><code>func bytes.Index(s, sep []byte) int</code></td>
<td><code>func strings.Index(s, substr string) int</code></td>
</tr>
<tr>
<td><code>func bytes.IndexAny(s []byte, chars string) int</code></td>
<td><code>func strings.IndexAny(s, chars string) int</code></td>
</tr>
<tr>
<td><code>func bytes.IndexByte(b []byte, c byte) int</code></td>
<td><code>func strings.IndexByte(s string, c byte) int</code></td>
</tr>
<tr>
<td><code>func bytes.IndexFunc(s []byte, f func(r rune) bool) int</code></td>
<td><code>func strings.IndexFunc(s string, f func(rune) bool) int</code></td>
</tr>
<tr>
<td><code>func bytes.IndexRune(s []byte, r rune) int</code></td>
<td><code>func strings.IndexRune(s string, r rune) int</code></td>
</tr>
<tr>
<td><code>func bytes.LastIndex(s, sep []byte) int</code></td>
<td><code>func strings.LastIndex(s, sep string) int</code></td>
</tr>
<tr>
<td><code>func bytes.LastIndexAny(s []byte, chars string) int</code></td>
<td><code>func strings.LastIndexAny(s, chars string) int</code></td>
</tr>
<tr>
<td><code>func bytes.LastIndexByte(s []byte, c byte) int</code></td>
<td><code>func strings.LastIndexByte(s string, c byte) int</code></td>
</tr>
<tr>
<td><code>func bytes.LastIndexFunc(s []byte, f func(r rune) bool) int</code></td>
<td><code>func strings.LastIndexFunc(s string, f func(rune) bool) int</code></td>
</tr>
<tr>
<td><code>bytes.NewBufferString</code></td>
<td><code>bytes.NewBuffer</code></td>
</tr>
<tr>
<td><code>func (h *maphash.Hash) Write(b []byte) (int, error)</code></td>
<td><code>func (h *maphash.Hash) WriteString(s string) (int, error)</code></td>
</tr>
<tr>
<td><code>func (rw *httptest.ResponseRecorder) Write(buf []byte) (int, error)</code></td>
<td><code>func (rw *httptest.ResponseRecorder) WriteString(str string) (int, error)</code></td>
</tr>
<tr>
<td><code>func (f *os.File) Write(b []byte) (n int, err error)</code></td>
<td><code>func (f *os.File) WriteString(s string) (n int, err error)</code></td>
</tr>
<tr>
<td><code>func regexp.Match(pattern string, b []byte) (bool, error)</code></td>
<td><code>func regexp.MatchString(pattern string, s string) (bool, error)</code></td>
</tr>
<tr>
<td><code>func (re *regexp.Regexp) FindAllIndex(b []byte, n int) [][]int</code></td>
<td><code>func (re *regexp.Regexp) FindAllStringIndex(s string, n int) [][]int</code></td>
</tr>
<tr>
<td><code>func (re *regexp.Regexp) FindAllSubmatch(b []byte, n int) [][][]byte</code></td>
<td><code>func (re *regexp.Regexp) FindAllStringSubmatch(s string, n int) [][]string</code></td>
</tr>
<tr>
<td><code>func (re *regexp.Regexp) FindIndex(b []byte) (loc []int)</code></td>
<td><code>func (re *regexp.Regexp) FindStringIndex(s string) (loc []int)</code></td>
</tr>
<tr>
<td><code>func (re *regexp.Regexp) FindSubmatchIndex(b []byte) []int</code></td>
<td><code>func (re *regexp.Regexp) FindStringSubmatchIndex(s string) []int</code></td>
</tr>
<tr>
<td><code>func (re *regexp.Regexp) Match(b []byte) bool</code></td>
<td><code>func (re *regexp.Regexp) MatchString(s string) bool</code></td>
</tr>
<tr>
<td><code>func (b *strings.Builder) Write(p []byte) (int, error)</code></td>
<td><code>func (b *strings.Builder) WriteByte(c byte) error</code></td>
</tr>
<tr>
<td><code>func utf8.Valid(p []byte) bool</code></td>
<td><code>func utf8.ValidString(s string) bool</code></td>
</tr>
</table>

## Install

### Compile From Source

You can get `mirror` with `go install` command.

```shell
go install github.com/butuzov/mirror/cmd/mirror@latest
```

## Usage

```shell
# include test files reports
mirror --with-tests ./...
```

