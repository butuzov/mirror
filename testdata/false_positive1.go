package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// using function that accepts string and process it and then return []bytes
func S1() {
	h := sha256.New()
	if _, err := h.Write([]byte("foobar")); err != nil {
		log.Fatal(err)
	}

	contentBuf := bytes.NewBufferString("foo-bar\n")
	contentBuf.WriteString(hex.EncodeToString(h.Sum(nil))) // false postitve
}
