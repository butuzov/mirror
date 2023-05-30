package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func bornToRune() {
	fmt.Printf("%T\n", 'п')
	fmt.Printf("%T\n", `п`)
	fmt.Printf("%T\n", "p"[0])

	var bbuf bytes.Buffer
	var sbuf strings.Builder
	buf := bufio.NewWriter(io.Discard)
	var r rune = 'п'
	b := "p"[0]

	sbuf.WriteString(string([]byte("foobar"))) // want `avoid allocations with \(\*strings\.Builder\)\.Write`
	sbuf.WriteString(string('п'))              // want `avoid allocations with \(\*strings\.Builder\)\.WriteRune`
	sbuf.WriteString(string('r'))              // want `avoid allocations with \(\*strings\.Builder\)\.WriteRune`
	sbuf.WriteString(string(`п`))
	sbuf.WriteString(string(b))

	bbuf.WriteString(string([]byte("foobar"))) // want `avoid allocations with \(\*bytes\.Buffer\)\.Write`
	bbuf.WriteString(string('п'))              // want `avoid allocations with \(\*bytes\.Buffer\)\.WriteRune`
	bbuf.WriteString(string(r))                // want `avoid allocations with \(\*bytes\.Buffer\)\.WriteRune`
	bbuf.WriteString(string(`п`))
	bbuf.WriteString(string(b))

	buf.WriteString(string([]byte("foobar"))) // want `avoid allocations with \(\*bufio\.Writer\)\.Write`
	buf.WriteString(string('п'))              // want `avoid allocations with \(\*bufio\.Writer\)\.WriteRune`
	buf.WriteString(string(r))                // want `avoid allocations with \(\*bufio\.Writer\)\.WriteRune`
	buf.WriteString(string(`п`))
	buf.WriteString(string(b))

	fmt.Println("strings.Builder:", sbuf.String())
	fmt.Println("  bytes.Buffer:", bbuf.String())
}
