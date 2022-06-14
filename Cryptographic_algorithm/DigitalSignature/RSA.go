package DigitalSignature

import (
	"fmt"
	"math/big"
)

func KeyGen(p, q int) (int, int, int) {
	//p,q - prime

	n := p * q

	phi := (p - 1) * (q - 1) //Euler function

	e := getE(phi)
	d := getD(e, phi)

	fmt.Println("PrivateKey", d, n, "PublicKey", e, n)
	return d, e, n
}

func getE(phi int) int {
	for e := 2; e < phi; e++ {
		if Coprime(e, phi) { //proverka na coprime
			return e
		}
	}
	return -1
}

func Coprime(e, phi int) bool {
	for e > 0 {
		var temp int
		temp = e
		e = phi % e
		phi = temp
	}
	return phi == 1
}

func getD(e, phi int) int { //bruteforce  d ⋅ e ≡ 1 (mod φ(n))
	x := 1
	for {
		x += phi
		m := x % e
		if m == 0 {
			return x / e
		}
	}
}

func Encrypt(message, e, n int) int {
	return powmod(message, e, n)
}

func Decrypt(ciphertext, d, n int) int {
	return powmod(ciphertext, d, n)
}

func powmod(message, a, b int) int { //using big.Int because of going out of limit
	Message := big.NewInt(int64(message))
	A := big.NewInt(int64(a))
	B := big.NewInt(int64(b))

	Message.Exp(Message, A, nil)
	_, m := big.NewInt(0).DivMod(Message, B, new(big.Int))
	ans := m.Int64()
	return int(ans)
}
