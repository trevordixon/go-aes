package aes

type AES struct {
	key []byte
}

func (a *AES) Encrypt(input []byte) []byte {
	state := make([]byte, 16)
	copy(state, input)

	roundKey := generateKey(a.key)
	numRounds := len(a.key)/4 + 5

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

func (a *AES) Decrypt(input []byte) []byte {
	state := make([]byte, 16)
	copy(state, input)

	numRounds := len(a.key)/4 + 5

	roundKey := generateKeyReverse(a.key)

	xor(state, roundKey())
	shiftRowsReverse(state)
	subBytesReverse(state)

	for i := 0; i < numRounds; i++ {
		xor(state, roundKey())
		mixColumnsReverse(state)
		shiftRowsReverse(state)
		subBytesReverse(state)
	}

	xor(state, roundKey())

	return state
}
