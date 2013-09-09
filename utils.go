package aes

func xor(word1 []byte, word2 []byte) {
	for i := range word1 {
		word1[i] ^= word2[i]
	}
}

func swap(state []byte, i, j int) {
	temp := state[i]
	state[i] = state[j]
	state[j] = temp
}
