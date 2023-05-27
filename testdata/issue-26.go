package main

import (
	"fmt"
	"strings"
)

func foobar_byte() {
	var strBuilder strings.Builder
	var text string = "text"
	var b byte = text[0]

	fmt.Printf("%T\n", b)

	strBuilder.WriteString(string(b))
	fmt.Println(strBuilder.String())
}
