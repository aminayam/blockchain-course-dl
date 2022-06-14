package main

import (
	"Cryptographic_algorithm/DigitalSignature"
	"Cryptographic_algorithm/VigenereCipher"
	"fmt"
)

func main() {
	//Vigenere
	//please, enter your plaintext and key ONLY lowercase letters :)
	//1st word in args is plaintext, 2nd - key
	encodeResult := VigenereCipher.MyChipher("helloworld", "blockchain")
	fmt.Println("Vigenere Ciphertext:", encodeResult)

	//RSA
	//it works now only for "small" integers
	//You can put your own prime numbers in func KeyGen() and your own message to encrypt/decrypt by this generated keys
	d, e, n := DigitalSignature.KeyGen(3, 11)
	encr := DigitalSignature.Encrypt(2, d, n)
	decr := DigitalSignature.Decrypt(29, e, n)
	fmt.Println("Cipher:", encr, "Plain:", decr)

}
