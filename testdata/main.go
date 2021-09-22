package main

import (
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
		_, _ = regexp.Match(`foo.*`, []byte(b1)) // want `you should be using regexp.MatchString`
	}
	{
		b1 := []byte("seafood")
		_, _ = regexp.MatchString(`foo.*`, string(b1)) // want `you should be using regexp.Match`
	}

	{
		foobar := `foobar`
		_, _ = Match(`foo.*`, []byte(foobar)) // want `you should be using regexp.Match`
	}
	{
		foobar := []byte(`foobar`)
		_, _ = MatchString(`foo.*`, string(foobar)) // want `you should be using regexp.Match`
	}

	{
		footbal := `football`
		_, _ = reg.Match(`foo.*`, []byte(footbal)) // want `you should be using regexp.MatchString`
	}
	{
		footbal := []byte(`football`)
		_, _ = reg.MatchString(`foo.*`, string(footbal)) // want `you should be using regexp\.Match`
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		str := "fool"
		_ = re1.Match([]byte(str)) // want `you should be using MatchString method`
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		str := []byte("fool")
		_ = re1.MatchString(string(str)) // want `you should be using Match method`
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		str := "fool"
		_ = re1.FindAllIndex([]byte(str), -1) // want `you should be using FindAllStringIndex method`
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		b := []byte("fool")
		_ = re1.FindAllStringIndex(string(b), -1) // want `you should be using FindAllIndex method`
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		s1 := "fool"
		s2 := "bool"
		_ = re1.ReplaceAll([]byte(s1), []byte(s2))
	}

	{
		re1, _ := regexp.Compile(`foo.*`)
		b1 := []byte("fool")
		b2 := []byte("bool")
		_ = re1.ReplaceAllString(string(b1), string(b2))
	}
}
