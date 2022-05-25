package password

import (
	"bytes"
	"strings"
)

var Hex = []byte("0123456789abcdef")

func nextByte(p string) byte { //generates next byte
	if p == "f" {
		return 102
	} else {
		idx := strings.Index(string(Hex), p)
		return Hex[idx+1]
	}
}

func nextPass(pass []byte, custom []byte) { //generates next password
	for i := 0; i < len(pass); i++ {
		if pass[i] == custom[i] {
			continue
		} else {
			pass[i] = nextByte(string(pass[i]))
		}
	}
}

func Bruteforce(pass []byte, custom []byte) string { //compares generated password with our password
	for !bytes.Equal(pass, custom) {
		nextPass(pass, custom)
	}
	return string(pass)
}
