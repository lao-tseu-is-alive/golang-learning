package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	text := "your text"
	hash := getHash(text)
	fmt.Printf("For text [%s] \nhash: %s\n", text, hash)
	text = "your text modified"
	hash = getHash(text)
	fmt.Printf("For text [%s] \nhash: %s\n", text, hash)
	text = "your text"
	hash = getHash(text)
	fmt.Printf("For text [%s] \nhash: %s\n", text, hash)
}

func getHash(s string) string {
	h := hmac.New(sha256.New, []byte("yourKey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
