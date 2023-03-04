package variant

import (
	"math/rand"
)

type Option struct {
	name  string
	value string
}

func RandomCode(length int) string {
	const (
		letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digitBytes  = "0123456789"
	)

	i := 0
	code := make([]byte, length)

	for ; i < length/2; i++ {
		code[i] = randomChar(digitBytes)
	}

	for ; i < length; i++ {
		code[i] = randomChar(letterBytes)
	}

	return string(code)
}

func randomChar(chars string) byte {
	return chars[rand.Intn(len(chars))]
}
