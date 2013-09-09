package aes

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

func shiftRowsReverse(state []byte) {
	swap(state, 9, 13)
	swap(state, 5, 9)
	swap(state, 1, 5)

	swap(state, 6, 10)
	swap(state, 2, 6)
	swap(state, 10, 14)
	swap(state, 6, 10)

	swap(state, 3, 7)
	swap(state, 7, 11)
	swap(state, 11, 15)
}
