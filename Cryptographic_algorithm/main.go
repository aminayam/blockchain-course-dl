package main

import (
	"Cryptographic_algorithm/VigenereCipher"
	"fmt"
)

func main() {
	//please, enter your plaintext and key ONLY lowercase letters :)
	encodeResult := VigenereCipher.MyChipher("helloworld", "blockchain")
	fmt.Println("Ciphertext:", encodeResult)
}
