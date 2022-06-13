package main

import (
	"fmt"
	"math/big"
	"math/bits"
	"strconv"
)

func main() {
	GetSHA1Hash("abc")
}

func binary(s string) string { //converts string message to binary bits representation
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func parseBinToHex(s string) uint32 {
	ui, _ := strconv.ParseUint(s, 2, 64)
	ans := fmt.Sprintf("%x", ui)
	res, _ := strconv.Atoi(ans)
	return uint32(res)
}

func GetSHA1Hash(message string) {

	var (
		h0 uint32 = 0x67452301
		h1 uint32 = 0xEFCDAB89
		h2 uint32 = 0x98BADCFE
		h3 uint32 = 0x10325476
		h4 uint32 = 0xC3D2E1F0
	) //variables initialization

	//Padding the Message
	//ex: message = "abc"
	ml := len(binary(message)) //length of the massage in bits = int ex: 24
	ost := ml % 512
	bitsMl := strconv.FormatInt(int64(ml), 2) //(length of the massage in bits) in bits = bits ex: 11000
	Ml := len(bitsMl)                         //length ((length of the massage in bits) in bits )= int ex:5

	byteArr := binary(message)
	byteArr += "1"
	for i := 0; i < (448 - ost - 1); i++ {
		byteArr += "0"
	} //fills 448 bits
	for i := 0; i < (64 - Ml); i++ {
		byteArr += "0"
	} //fills length to 64 bits
	byteArr += bitsMl

	//break message into 512-bit chunks
	w := make([][]uint32, 80, 80)

	for i := 0; i < 80; i++ {
		w[i] = make([]uint32, 80)
		for j := 0; j < 80; j++ {
			w[i][j] = 0
		}
	}

	block := 0
	for k := 0; k < len(byteArr); k += 512 {
		for i, j := 0, 0; i < k+512 && j < 16; j++ {
			w[block][j] = parseBinToHex(byteArr[k+i : k+i+32])
			i += 32
		}
		for j := 16; j < 80; j++ {
			w[block][j] = bits.RotateLeft32(w[block][j-3]^w[block][j-8]^w[block][j-14]^w[block][j-16], 5)
		}
		block++
	}
	var (
		a = h0
		b = h1
		c = h2
		d = h3
		e = h4
	)

	var f, k, temp uint32

	for iter := 0; iter < block; iter++ {
		for i := 0; i < 80; i++ {
			if 0 <= i && i <= 19 {
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			} else if 20 <= i && i <= 39 {
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			} else if 40 <= i && i <= 59 {
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			} else if 60 <= i && i <= 79 {
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}
			temp = bits.RotateLeft32(a, 5) + f + e + k + w[iter][i]
			e = d
			d = c
			c = bits.RotateLeft32(b, 30)
			b = a
			a = temp
			//fmt.Println(a, b, c, d, e)
		}
		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
	}

	H0 := big.NewInt(int64(h0))
	H1 := big.NewInt(int64(h1))
	H2 := big.NewInt(int64(h2))
	H3 := big.NewInt(int64(h3))
	H4 := big.NewInt(int64(h4))

	hash := big.NewInt(0)

	hash.Or(H0.Lsh(H0, 128), H1.Lsh(H1, 96))
	hash.Or(hash, H2.Lsh(H2, 64))
	hash.Or(hash, H3.Lsh(H3, 32))
	hash.Or(hash, H4)
	// concat

	result := hash.Text(16)
	fmt.Println(result)

}
