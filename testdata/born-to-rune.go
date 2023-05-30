package main

import (
	"bytes"
	"fmt"
	"strings"
)

func bornToRune() {
	fmt.Printf("%T\n", 'п')
	fmt.Printf("%T\n", `п`)
	fmt.Printf("%T\n", "p"[0])

	var bbuf bytes.Buffer
	var sbuf strings.Builder
	var r rune = 'п'
	b := "p"[0]

	sbuf.WriteString(string([]byte("foobar"))) // want `avoid allocations with \(\*strings\.Builder\)\.Write`
	sbuf.WriteString(string('п'))              // want `avoid allocations with \(\*strings\.Builder\)\.WriteRune`
	sbuf.WriteString(string('r'))              // want `avoid allocations with \(\*strings\.Builder\)\.WriteRune`
	sbuf.WriteString(string(`п`))
	sbuf.WriteString(string(b)) // want `avoid allocations with \(\*strings\.Builder\)\.WriteByte`

	bbuf.WriteString(string([]byte("foobar"))) // want `avoid allocations with \(\*bytes\.Buffer\)\.Write`
	bbuf.WriteString(string('п'))              // want `avoid allocations with \(\*bytes\.Buffer\)\.WriteRune`
	bbuf.WriteString(string(r))                // want `avoid allocations with \(\*bytes\.Buffer\)\.WriteRune`
	bbuf.WriteString(string(`п`))
	bbuf.WriteString(string(b)) // want `avoid allocations with \(\*bytes\.Buffer\)\.WriteByte`

	fmt.Println("strings.Builder:", sbuf.String())
	fmt.Println("  bytes.Buffer:", bbuf.String())
}
