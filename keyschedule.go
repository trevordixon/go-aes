package aes

type expander func() []byte

func generateKey(key []byte) expander {
	nk := len(key) / 4
	q := make(chan []byte, nk)

	temp := make([]byte, 4)
	copy(temp, key[nk*4-4:nk*4])

	i := 0
	next := func() []byte {
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
			subBytes(temp)
		}

		xor(temp, <-q)
		save := make([]byte, 4)
		copy(save, temp)
		q <- save

		return save
	}

	return func() []byte {
		roundKey := make([]byte, 16)

		copy(roundKey[0:4], next())
		copy(roundKey[4:8], next())
		copy(roundKey[8:12], next())
		copy(roundKey[12:16], next())

		return roundKey
	}
}

func generateKeyReverse(key []byte) expander {
	roundKey := generateKey(key)
	numRounds := len(key)/4 + 7
	roundKeys := make([][]byte, numRounds)
	for i := numRounds - 1; i >= 0; i-- {
		roundKeys[i] = roundKey()
	}

	i := 0
	return func() []byte {
		defer (func() { i++ })()
		return roundKeys[i]
	}
}

func core(word []byte, i int) {
	rotWord(word)
	subBytes(word)
	rcon(word, i)
}

func rotWord(word []byte) {
	first := word[0]
	copy(word, word[1:])
	word[3] = first
}

func xor(word1 []byte, word2 []byte) {
	for i := range word1 {
		word1[i] ^= word2[i]
	}
}
