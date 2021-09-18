package main

import (
	"fmt"
	"regexp"
	. "regexp"
	reg "regexp"
)

type ( // for sake of keeping imports active, while dis/enabling test cases
	s1 regexp.Regexp
	s2 reg.Regexp
	s3 Regexp
)

func main() {
	{
		b1 := "seafood"
		matched, _ := regexp.Match(`foo.*`, []byte(b1)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("regexp - %#v\n", matched)
	}
	{
		b1 := []byte("seafood")
		matched, _ := regexp.MatchString(`foo.*`, string(b1)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("regexp - %#v\n", matched)
	}

	{
		foobar := `foobar`
		matched, _ := Match(`foo.*`, []byte(foobar)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("dot import - %#v\n", matched)
	}
	{
		foobar := []byte(`foobar`)
		matched, _ := MatchString(`foo.*`, string(foobar)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("dot import - %#v\n", matched)
	}

	{
		footbal := `football`
		matched, _ := reg.Match(`foo.*`, []byte(footbal)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("named - %#v\n", matched)
	}
	{
		footbal := []byte(`football`)
		matched, _ := reg.MatchString(`foo.*`, string(footbal)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("named - %#v\n", matched)
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		str := "fool"
		matched := re1.Match([]byte(str)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("regexp.Regexp - %#v\n", matched)
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		str := []byte("fool")
		matched := re1.MatchString(string(str)) // want `Match(String)? can be \((.*)\)`
		fmt.Printf("regexp.Regexp - %#v\n", matched)
	}
}
