package VigenereCipher

func VigenereCipher(message string, key string) string {
	m := len(message)
	Key := NewKey(m, key)
	var ciphertext string

	for i := 0; i < m; i++ {
		ciphertext += string((message[i] + Key[i]) % 26)
	}
	return ciphertext
}

func NewKey(m int, key string) string {
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
