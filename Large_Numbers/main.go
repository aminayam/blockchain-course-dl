package main

import (
	"Large_Numbers/password"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var numbers = [10]int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) //randomizer by time config
	Numbers()
	t := time.Now()
	var exmplPass = []byte("7b3b9f1d2c") //you can enter your custom valid hexadecimal password or choose generated one
	var pass = make([]byte, len(exmplPass))
	for i := range pass {
		pass[i] = 48
	}
	fmt.Println("Example pass is: ", password.Bruteforce(pass, exmplPass))
	fmt.Println("Time: ", time.Since(t))
}

func Numbers() {
	var i int64 = 0
	for ; i < int64(len(numbers)); i++ {
		fmt.Printf("sequence length: %d, max key: %d, rand key:", numbers[i], maxKey(numbers[i]))
		fmt.Println(randKey(numbers[i]))
	}
}

func maxKey(pow int64) *big.Int {
	ans := new(big.Int).Exp(big.NewInt(2), big.NewInt(pow), nil)
	return ans
}

func randKey(l int64) string { //generates each digit of hex number
	b := make([]byte, l/4)
	for i := range b {
		b[i] = password.Hex[rand.Intn(len(password.Hex))]
	}
	return string(b)
}
