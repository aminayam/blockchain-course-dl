package VigenereCipher

func MyChipher(plaintext string, key string) string { //VigenereCipher realization
	m := len(plaintext)
	Key := NewKey(m, key)
	var ciphertext string
	for i := 0; i < m; i++ {
		res := (plaintext[i] - 92 + Key[i] - 102) % 26 //this peculiar numbers implements magic with ascii in golang
		ciphertext += string(res + 97)
	}
	return ciphertext
}

func NewKey(m int, key string) string { //generates new keyword based on key that repeats until it matches the length of the plaintext
	i := 0
	for {
		if m == i {
			i = 0
		} else if len(key) == m {
			break
		}
		key += string(key[i])
		i++
	}
	return key
}
