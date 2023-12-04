// Code generated by generate-tests; DO NOT EDIT.

package main

import (
	"strings"
	. "strings"
	pkg "strings"
)


func main_strings() {
	{
		
		_ = strings.Compare("foobar","foobar") 
	}

	{
		
		_ = strings.Compare("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.Compare(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.Compare(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Compare`
	}

	{
		
		_ = Compare("foobar","foobar") 
	}

	{
		
		_ = Compare("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = Compare(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = Compare(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Compare`
	}

	{
		
		_ = pkg.Compare("foobar","foobar") 
	}

	{
		
		_ = pkg.Compare("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.Compare(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.Compare(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Compare`
	}

	{
		
		_ = strings.Contains("foobar","foobar") 
	}

	{
		
		_ = strings.Contains("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.Contains(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.Contains(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Contains`
	}

	{
		
		_ = Contains("foobar","foobar") 
	}

	{
		
		_ = Contains("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = Contains(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = Contains(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Contains`
	}

	{
		
		_ = pkg.Contains("foobar","foobar") 
	}

	{
		
		_ = pkg.Contains("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.Contains(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.Contains(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Contains`
	}

	{
		
		_ = strings.ContainsAny(string([]byte{'f','o','o','b','a','r'}),"foobar") // want `avoid allocations with bytes\.ContainsAny`
	}

	{
		
		_ = strings.ContainsAny("foobar","foobar") 
	}

	{
		
		_ = ContainsAny(string([]byte{'f','o','o','b','a','r'}),"foobar") // want `avoid allocations with bytes\.ContainsAny`
	}

	{
		
		_ = ContainsAny("foobar","foobar") 
	}

	{
		
		_ = pkg.ContainsAny(string([]byte{'f','o','o','b','a','r'}),"foobar") // want `avoid allocations with bytes\.ContainsAny`
	}

	{
		
		_ = pkg.ContainsAny("foobar","foobar") 
	}

	{
		
		_ = strings.ContainsRune(string([]byte{'f','o','o','b','a','r'}),'ф') // want `avoid allocations with bytes\.ContainsRune`
	}

	{
		
		_ = strings.ContainsRune("foobar",'ф') 
	}

	{
		
		_ = ContainsRune(string([]byte{'f','o','o','b','a','r'}),'ф') // want `avoid allocations with bytes\.ContainsRune`
	}

	{
		
		_ = ContainsRune("foobar",'ф') 
	}

	{
		
		_ = pkg.ContainsRune(string([]byte{'f','o','o','b','a','r'}),'ф') // want `avoid allocations with bytes\.ContainsRune`
	}

	{
		
		_ = pkg.ContainsRune("foobar",'ф') 
	}

	{
		
		_ = strings.Count("foobar", "foobar") 
	}

	{
		
		_ = strings.Count("foobar", string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.Count(string([]byte{'f','o','o','b','a','r'}), "foobar") 
	}

	{
		
		_ = strings.Count(string([]byte{'f','o','o','b','a','r'}), string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Count`
	}

	{
		
		_ = Count("foobar", "foobar") 
	}

	{
		
		_ = Count("foobar", string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = Count(string([]byte{'f','o','o','b','a','r'}), "foobar") 
	}

	{
		
		_ = Count(string([]byte{'f','o','o','b','a','r'}), string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Count`
	}

	{
		
		_ = pkg.Count("foobar", "foobar") 
	}

	{
		
		_ = pkg.Count("foobar", string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.Count(string([]byte{'f','o','o','b','a','r'}), "foobar") 
	}

	{
		
		_ = pkg.Count(string([]byte{'f','o','o','b','a','r'}), string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Count`
	}

	{
		
		_ = strings.EqualFold("foobar","foobar") 
	}

	{
		
		_ = strings.EqualFold("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.EqualFold(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.EqualFold(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.EqualFold`
	}

	{
		
		_ = EqualFold("foobar","foobar") 
	}

	{
		
		_ = EqualFold("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = EqualFold(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = EqualFold(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.EqualFold`
	}

	{
		
		_ = pkg.EqualFold("foobar","foobar") 
	}

	{
		
		_ = pkg.EqualFold("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.EqualFold(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.EqualFold(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.EqualFold`
	}

	{
		
		_ = strings.HasPrefix("foobar","foobar") 
	}

	{
		
		_ = strings.HasPrefix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.HasPrefix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.HasPrefix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasPrefix`
	}

	{
		
		_ = HasPrefix("foobar","foobar") 
	}

	{
		
		_ = HasPrefix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = HasPrefix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = HasPrefix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasPrefix`
	}

	{
		
		_ = pkg.HasPrefix("foobar","foobar") 
	}

	{
		
		_ = pkg.HasPrefix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.HasPrefix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.HasPrefix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasPrefix`
	}

	{
		
		_ = strings.HasSuffix("foobar","foobar") 
	}

	{
		
		_ = strings.HasSuffix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.HasSuffix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.HasSuffix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasSuffix`
	}

	{
		
		_ = HasSuffix("foobar","foobar") 
	}

	{
		
		_ = HasSuffix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = HasSuffix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = HasSuffix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasSuffix`
	}

	{
		
		_ = pkg.HasSuffix("foobar","foobar") 
	}

	{
		
		_ = pkg.HasSuffix("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.HasSuffix(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.HasSuffix(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.HasSuffix`
	}

	{
		
		_ = strings.Index("foobar","foobar") 
	}

	{
		
		_ = strings.Index("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.Index(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.Index(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Index`
	}

	{
		
		_ = Index("foobar","foobar") 
	}

	{
		
		_ = Index("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = Index(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = Index(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Index`
	}

	{
		
		_ = pkg.Index("foobar","foobar") 
	}

	{
		
		_ = pkg.Index("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.Index(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.Index(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.Index`
	}

	{
		
		_ = strings.IndexAny(string([]byte{'f','o','o','b','a','r'}), "f") // want `avoid allocations with bytes\.IndexAny`
	}

	{
		
		_ = strings.IndexAny("foobar", "f") 
	}

	{
		
		_ = IndexAny(string([]byte{'f','o','o','b','a','r'}), "f") // want `avoid allocations with bytes\.IndexAny`
	}

	{
		
		_ = IndexAny("foobar", "f") 
	}

	{
		
		_ = pkg.IndexAny(string([]byte{'f','o','o','b','a','r'}), "f") // want `avoid allocations with bytes\.IndexAny`
	}

	{
		
		_ = pkg.IndexAny("foobar", "f") 
	}

	{
		
		_ = strings.IndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.IndexByte`
	}

	{
		
		_ = strings.IndexByte("foobar", byte('f')) 
	}

	{
		
		_ = IndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.IndexByte`
	}

	{
		
		_ = IndexByte("foobar", byte('f')) 
	}

	{
		
		_ = pkg.IndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.IndexByte`
	}

	{
		
		_ = pkg.IndexByte("foobar", byte('f')) 
	}

	{
		
		_ = strings.IndexFunc(string([]byte{'f','o','o','b','a','r'}),func(r rune) bool { return true }) // want `avoid allocations with bytes\.IndexFunc`
	}

	{
		
		_ = strings.IndexFunc("foobar",func(r rune) bool { return true }) 
	}

	{
		
		_ = IndexFunc(string([]byte{'f','o','o','b','a','r'}),func(r rune) bool { return true }) // want `avoid allocations with bytes\.IndexFunc`
	}

	{
		
		_ = IndexFunc("foobar",func(r rune) bool { return true }) 
	}

	{
		
		_ = pkg.IndexFunc(string([]byte{'f','o','o','b','a','r'}),func(r rune) bool { return true }) // want `avoid allocations with bytes\.IndexFunc`
	}

	{
		
		_ = pkg.IndexFunc("foobar",func(r rune) bool { return true }) 
	}

	{
		
		_ = strings.IndexRune(string([]byte{'f','o','o','b','a','r'}), rune('ф')) // want `avoid allocations with bytes\.IndexRune`
	}

	{
		
		_ = strings.IndexRune("foobar", rune('ф')) 
	}

	{
		
		_ = IndexRune(string([]byte{'f','o','o','b','a','r'}), rune('ф')) // want `avoid allocations with bytes\.IndexRune`
	}

	{
		
		_ = IndexRune("foobar", rune('ф')) 
	}

	{
		
		_ = pkg.IndexRune(string([]byte{'f','o','o','b','a','r'}), rune('ф')) // want `avoid allocations with bytes\.IndexRune`
	}

	{
		
		_ = pkg.IndexRune("foobar", rune('ф')) 
	}

	{
		
		_ = strings.LastIndex("foobar","foobar") 
	}

	{
		
		_ = strings.LastIndex("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = strings.LastIndex(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = strings.LastIndex(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.LastIndex`
	}

	{
		
		_ = LastIndex("foobar","foobar") 
	}

	{
		
		_ = LastIndex("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = LastIndex(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = LastIndex(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.LastIndex`
	}

	{
		
		_ = pkg.LastIndex("foobar","foobar") 
	}

	{
		
		_ = pkg.LastIndex("foobar",string([]byte{'f','o','o','b','a','r'})) 
	}

	{
		
		_ = pkg.LastIndex(string([]byte{'f','o','o','b','a','r'}),"foobar") 
	}

	{
		
		_ = pkg.LastIndex(string([]byte{'f','o','o','b','a','r'}),string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.LastIndex`
	}

	{
		
		_ = strings.LastIndexAny(string([]byte{'f','o','o','b','a','r'}),"f") // want `avoid allocations with bytes\.LastIndexAny`
	}

	{
		
		_ = strings.LastIndexAny("foobar","f") 
	}

	{
		
		_ = LastIndexAny(string([]byte{'f','o','o','b','a','r'}),"f") // want `avoid allocations with bytes\.LastIndexAny`
	}

	{
		
		_ = LastIndexAny("foobar","f") 
	}

	{
		
		_ = pkg.LastIndexAny(string([]byte{'f','o','o','b','a','r'}),"f") // want `avoid allocations with bytes\.LastIndexAny`
	}

	{
		
		_ = pkg.LastIndexAny("foobar","f") 
	}

	{
		
		_ = strings.LastIndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.LastIndexByte`
	}

	{
		
		_ = strings.LastIndexByte("foobar", byte('f')) 
	}

	{
		
		_ = LastIndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.LastIndexByte`
	}

	{
		
		_ = LastIndexByte("foobar", byte('f')) 
	}

	{
		
		_ = pkg.LastIndexByte(string([]byte{'f','o','o','b','a','r'}), byte('f')) // want `avoid allocations with bytes\.LastIndexByte`
	}

	{
		
		_ = pkg.LastIndexByte("foobar", byte('f')) 
	}

	{
		
		_ = strings.LastIndexFunc(string([]byte{'f','o','o','b','a','r'}), func(r rune) bool { return true }) // want `avoid allocations with bytes\.LastIndexFunc`
	}

	{
		
		_ = strings.LastIndexFunc("foobar", func(r rune) bool { return true }) 
	}

	{
		
		_ = LastIndexFunc(string([]byte{'f','o','o','b','a','r'}), func(r rune) bool { return true }) // want `avoid allocations with bytes\.LastIndexFunc`
	}

	{
		
		_ = LastIndexFunc("foobar", func(r rune) bool { return true }) 
	}

	{
		
		_ = pkg.LastIndexFunc(string([]byte{'f','o','o','b','a','r'}), func(r rune) bool { return true }) // want `avoid allocations with bytes\.LastIndexFunc`
	}

	{
		
		_ = pkg.LastIndexFunc("foobar", func(r rune) bool { return true }) 
	}

	{
		
		_ = strings.NewReader(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.NewReader`
	}

	{
		
		_ = strings.NewReader("foobar") 
	}

	{
		
		_ = NewReader(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.NewReader`
	}

	{
		
		_ = NewReader("foobar") 
	}

	{
		
		_ = pkg.NewReader(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with bytes\.NewReader`
	}

	{
		
		_ = pkg.NewReader("foobar") 
	}

	{
		builder := strings.Builder{}
		_,_ = builder.Write([]byte("foobar")) // want `avoid allocations with \(\*strings\.Builder\)\.WriteString`
	}

	{
		builder := strings.Builder{}
		_,_ = builder.Write([]byte{'f','o','o','b','a','r'}) 
	}

	{
		builder := Builder{}
		_,_ = builder.Write([]byte("foobar")) // want `avoid allocations with \(\*strings\.Builder\)\.WriteString`
	}

	{
		builder := Builder{}
		_,_ = builder.Write([]byte{'f','o','o','b','a','r'}) 
	}

	{
		builder := pkg.Builder{}
		_,_ = builder.Write([]byte("foobar")) // want `avoid allocations with \(\*strings\.Builder\)\.WriteString`
	}

	{
		builder := pkg.Builder{}
		_,_ = builder.Write([]byte{'f','o','o','b','a','r'}) 
	}

	{
		builder := strings.Builder{}
		_,_ = builder.WriteString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with \(\*strings\.Builder\)\.Write`
	}

	{
		builder := strings.Builder{}
		_,_ = builder.WriteString("foobar") 
	}

	{
		builder := Builder{}
		_,_ = builder.WriteString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with \(\*strings\.Builder\)\.Write`
	}

	{
		builder := Builder{}
		_,_ = builder.WriteString("foobar") 
	}

	{
		builder := pkg.Builder{}
		_,_ = builder.WriteString(string([]byte{'f','o','o','b','a','r'})) // want `avoid allocations with \(\*strings\.Builder\)\.Write`
	}

	{
		builder := pkg.Builder{}
		_,_ = builder.WriteString("foobar") 
	}

}
