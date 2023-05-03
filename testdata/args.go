package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const s3 string = "a"

// strins are ot string
func v(n int) (err error) {
	n += 1
	_, err = fmt.Fprintf(os.Stdout, "sss: %d", n)
	return err
}

func m() {
	golden := "foobar"

	// pointer to ...
	// {
	// 	s0 := []byte("foobar")
	// 	s1 := &s0
	// 	_ = strings.IndexAny(string(*s1), "foobar")  want `avoid allocations with strings\.IndexAny`
	// }
	{
		const s0 = "a"
		_ = bytes.IndexAny([]byte(s0), golden) // want `avoid allocations with strings\.IndexAny`
	}
	{
		const s0 string = "a"
		_ = bytes.IndexAny([]byte(s0), golden) // want `avoid allocations with strings\.IndexAny`
	}
	{
		_ = bytes.IndexAny([]byte(s3), golden) // want `avoid allocations with strings\.IndexAny`
	}

	// constant (string)
	{
		var s0 string = "a"
		_ = bytes.IndexAny([]byte(s0), golden) // want `avoid allocations with strings\.IndexAny`
	}
	{
		const s0 = "foo"
		_ = bytes.IndexAny([]byte(s0), golden) // want `avoid allocations with strings\.IndexAny`
	}
	// function return string
	{
		s0 := func() string { return golden }
		_ = bytes.IndexAny([]byte(s0()), golden) // want `avoid allocations with strings\.IndexAny`
	}
	// function return byte
	{
		s0 := func() []byte { return []byte(golden) }
		_ = strings.IndexAny(string(s0()), golden) // want `avoid allocations with bytes\.IndexAny`
	}
	{
		s0 := [][]byte{[]byte("foobar")}
		_ = strings.IndexAny(string(s0[0]), golden) // want `avoid allocations with bytes\.IndexAny`
	}
	{
		s0 := "foo"
		s1 := "foo"
		_ = bytes.IndexAny([]byte(s0+s1), golden) // want `avoid allocations with strings\.IndexAny`
	}
	{
		s0 := []string{"foobar"}
		_ = bytes.IndexAny([]byte(s0[0]), golden) // want `avoid allocations with strings\.IndexAny`
	}
}
