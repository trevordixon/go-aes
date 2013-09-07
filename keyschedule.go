package aes

type expander func() []byte

func generateKey(key []byte, nk int) expander {
	q := make(chan []byte, nk)

	temp := make([]byte, 4)
	copy(temp, key[nk*4-4:nk*4])

	i := 0
	return func() []byte {
		defer (func() { i++ })()

		if i < nk {
			next := key[i*4 : i*4+4]
			q <- next
			return next
		}

		if i%nk == 0 {
			core(temp, i/nk)
		}

		if nk == 8 && (i-4)%nk == 0 {
			subWord(temp)
		}

		xor(temp, <-q)
		save := make([]byte, 4)
		copy(save, temp)
		q <- save

		return save
	}
}

func generate128(key []byte) expander {
	return generateKey(key, 4)
}

func generate192(key []byte) expander {
	return generateKey(key, 6)
}

func generate256(key []byte) expander {
	return generateKey(key, 8)
}

func core(word []byte, i int) {
	rotWord(word)
	subWord(word)
	rcon(word, i)
}

func rotWord(word []byte) {
	first := word[0]
	copy(word, word[1:])
	word[3] = first
}

func subWord(word []byte) {
	for i, b := range word {
		word[i] = sbox[b]
	}
}

func rcon(word []byte, i int) {
	word[0] ^= rcons[i]
}

func xor(word1 []byte, word2 []byte) {
	for i, _ := range word1 {
		word1[i] ^= word2[i]
	}
}
