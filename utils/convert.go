package utils

import (
	"bytes"
	"strings"
)

const (
	ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"
	BASE     = len(ALPHABET)
)

func reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Encodes an ID back to its key
func Encode(id int) string {
	var out bytes.Buffer

	// Special case
	if id == 0 {
		return string(ALPHABET[0])
	}

	for id > 0 {
		out.WriteByte(ALPHABET[id%BASE])
		id = id / BASE
	}

	return reverse(out.String())
}

// Decodes a key back to its ID
func Decode(key string) int {
	var id int

	for _, char := range key {
		id = id*BASE + strings.IndexRune(ALPHABET, char)
	}

	return id
}
