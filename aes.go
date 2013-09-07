package aes

type AES struct {
	key []byte
}

func (a *AES) Encode(input []byte) []byte {
	state := make([]byte, 16)
	copy(state, input)

	numRounds := len(a.key)/4 + 5

	var roundKey expander
	switch len(a.key) {
	case 16:
		roundKey = generate128(a.key)
	case 24:
		roundKey = generate192(a.key)
	case 32:
		roundKey = generate256(a.key)
	}

	xor(state, roundKey())

	for i := 0; i < numRounds; i++ {
		subBytes(state)
		shiftRows(state)
		mixColumns(state)
		xor(state, roundKey())
	}

	subBytes(state)
	shiftRows(state)
	xor(state, roundKey())

	return state
}

func (a *AES) Decode(input []byte) []byte {
	return input
}

func swap(state []byte, i, j int) {
	temp := state[i]
	state[i] = state[j]
	state[j] = temp
}

func shiftRows(state []byte) {
	swap(state, 1, 5)
	swap(state, 5, 9)
	swap(state, 9, 13)

	swap(state, 6, 10)
	swap(state, 10, 14)
	swap(state, 2, 6)
	swap(state, 6, 10)

	swap(state, 11, 15)
	swap(state, 7, 11)
	swap(state, 3, 7)
}
