package DigitalSignature

import (
	"fmt"
	"math/big"
)

func KeyGen() (*big.Int, *big.Int, *big.Int) {
	var p, q *big.Int //p,q - prime
	p, _ = big.NewInt(0).SetString("3", 10)
	q, _ = big.NewInt(0).SetString("13", 10)

	n := big.NewInt(0).Mul(p, q)
	p.Add(p, big.NewInt(-1))
	q.Add(q, big.NewInt(-1))

	phi := big.NewInt(0).Mul(p, q) //Euler function

	e := getE(phi)
	d := getD(e, phi)

	fmt.Println("PrivateKey", d, n, "PublicKey", e, n, phi)
	return d, e, n
}

func getE(phi *big.Int) *big.Int {
	e := big.NewInt(2)
	for {
		if Coprime(e, phi) { //proverka na coprime
			return e
		}
		if e.Cmp(phi) >= 0 { //calculates while e<phi
			break
		}
		e.Add(e, big.NewInt(1))
	}

	return big.NewInt(-1)
}

func Coprime(a, b *big.Int) bool { //finds coprime number
	var t *big.Int
	for {
		t = b
		_, m := big.NewInt(0).DivMod(a, b, new(big.Int))
		b = m
		a = t
		if b.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return a.Cmp(big.NewInt(1)) == 0
}

func getD(e, phi *big.Int) *big.Int { //bruteforce  d ⋅ e ≡ 1 (mod φ(n))
	x := big.NewInt(1)
	for {
		ans, m := big.NewInt(0).DivMod(x, e, new(big.Int))

		if m.Cmp(big.NewInt(0)) == 0 {
			return ans
		}
		x.Add(x, phi)
	}
}

func Encrypt(message, e, n *big.Int) *big.Int {
	return powmod(message, e, n)
}

func Decrypt(ciphertext, d, n *big.Int) *big.Int {
	return powmod(ciphertext, d, n)
}

func powmod(message, a, b *big.Int) *big.Int {
	message.Exp(message, a, nil)
	_, m := big.NewInt(0).DivMod(message, b, new(big.Int))
	return m
}
