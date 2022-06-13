package DigitalSignature

import (
	"fmt"
	"math/big"
)

func KeyGen() { //P, Q string
	var p, q, n, phi *big.Int //p,q - primary
	p = big.NewInt(10001)
	q = big.NewInt(253)
	fmt.Println(p, q)

	//A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself
	//n.Mul(big.NewInt(-1), big.NewInt(0))
	p.Add(p, big.NewInt(-10))
	q.Add(q, big.NewInt(-1))
	fmt.Println(p, q, n)
	phi = phi.Mul(p, q) //Euler totient function
	e := getE(phi)
	d := getD(e, phi)
	fmt.Println(d, n, e, n)
	//PrivateKey(d, n)
	//PublicKey(e, n)
}

func getE(phi *big.Int) *big.Int {
	// Выбирается целое число e ( 1 < e < t ) // взаимно простое со значением функции Эйлера (t)

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
func Coprime(a, b *big.Int) bool {
	// Set a to gcd(a,b)
	var t, x, m *big.Int
	for {
		t = b
		x.DivMod(a, b, m)
		b = m
		a = t
		if b.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	// By definition, a and b are coprime if gcd(a,b) == 1
	return a.Cmp(big.NewInt(1)) == 0
}

func getD(e, phi *big.Int) *big.Int { //bruteforce
	// Вычисляется число d, мультипликативно обратное к числу e по модулю φ(n), то есть число, удовлетворяющее сравнению:
	//    d ⋅ e ≡ 1 (mod φ(n))
	x := big.NewInt(1)
	var z, m1, m2, m *big.Int
	for x.Cmp(phi) <= 0 {
		//m1 - e%phi
		//m2 - x%phi
		//m -(m1*m2)%phi
		z.DivMod(e, phi, m1)
		z.DivMod(x, phi, m2)
		z.DivMod(m1, m2, m)

		if m.Cmp(big.NewInt(1)) == 0 {
			return x
		}
		x.Add(x, big.NewInt(1))
	}
	return big.NewInt(-1)
}

/*
func Encrypt(message string, PublicKey int) {
	return powmod(message, e, n)
}

func Decrypt(ciphertext string, PrivateKey int) {
	return powmod(ciphertext, d, n)
}


*/
