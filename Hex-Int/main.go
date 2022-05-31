package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() { //test
	/*you can put your custom valid number
	Sorry, but you need to write only lowercase letters in hex number. Exp: "abc234d9", NOT "AbC23zD9".
	When I have more time, I will fix it <3
	*/
	littleNum := new(big.Int)
	littleNum.SetString("255", 10)
	bigNum := new(big.Int)
	bigNum.SetString("115339776388732929035197660848497720713218148788040405586178452820382218977280", 10)
	fmt.Println(LittleHex(littleNum), BigHex(bigNum))
	fmt.Println(HexLittle("ff00000000000000000000000000000000000000000000000000000000000000"), HexBig("ff00000000000000000000000000000000000000000000000000000000000000"))

}

var Hex = "0123456789abcdef"

func ConvByte(str1, str2 string) *big.Int { //func converts simple hex to one-byte-demical
	a := strings.Index(Hex, str1)*16 + strings.Index(Hex, str2)
	ans := big.NewInt(int64(a))
	return ans
}

func Reverse(str string) string { //func reverses number to read it according to the big principle
	var ans string
	for _, item := range str {
		ans = string(item) + ans
	}
	return ans
}

func HexLittle(str string) *big.Int { // Converting HEX values to Little Endian values
	d, ans, z := big.NewInt(1), big.NewInt(0), big.NewInt(0) //d = digit of the number
	num := big.NewInt(256)                                   //two hex digits
	for i := 0; i < len([]rune(str)); i += 2 {               //in each iteration we separate the byte and get it into a decimal number (in accordance with the digit)
		ans.Add(ans, z.Mul(d, ConvByte(string(str[i]), string(str[i+1]))))
		d.Mul(d, num) //increase digit
	}
	return ans
}

func HexBig(str string) *big.Int { // Converting HEX values to Big Endian values
	ans := big.NewInt(0) //func performs the same as little-algorithm, but with reverse order of reading
	st := Reverse(str)
	ans = HexLittle(st)
	return ans
}

func LittleHex(num *big.Int) string { //Converting Little Endian values to HEX values
	k, mod := big.NewInt(0), num //classical algorithm of converting a number to bigger number system
	var ans string
	for {
		k.DivMod(num, big.NewInt(16), mod)
		ans = ans + string(Hex[int(mod.Int64())])
		num = k
		if k.Cmp(big.NewInt(0)) <= 0 { //calculates mod while we can divide number
			break
		}
	}
	return ans
}

func BigHex(num *big.Int) string { //Converting Big Endian values to HEX values
	str := LittleHex(num) //func performs the same as little-algorithm, but with reverse order of reading
	ans := Reverse(str)
	return ans
}
