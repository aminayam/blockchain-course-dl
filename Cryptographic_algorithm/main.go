package main

import (
	"Cryptographic_algorithm/DigitalSignature"
	"Cryptographic_algorithm/VigenereCipher"
	"fmt"
	"math/big"
)

func main() {
	//Vigenere
	//please, enter your plaintext and key ONLY lowercase letters :)
	//1st word in args is plaintext, 2nd - key
	encodeResult := VigenereCipher.MyChipher("helloworld", "blockchain")
	fmt.Println("Vigenere Ciphertext:", encodeResult)

	//RSA
	d, e, n := DigitalSignature.KeyGen()
	encr := DigitalSignature.Encrypt(big.NewInt(123), e, n)
	decr := DigitalSignature.Decrypt(big.NewInt(15), d, n)
	fmt.Println("Cipher:", encr, "Plain:", decr)

}
