package aes

import (
	"encoding/hex"
	"testing"
)

var shiftRowsTests = []struct {
	before string
	after  string
}{
	{"d42711aee0bf98f1b8b45de51e415230", "d4bf5d30e0b452aeb84111f11e2798e5"},
	{"49ded28945db96f17f39871a7702533b", "49db873b453953897f02d2f177de961a"},
	{"ac73cf7befc111df13b5d6b545235ab8", "acc1d6b8efb55a7b1323cfdf457311b5"},
}

func TestShiftRows(t *testing.T) {
	for _, test := range shiftRowsTests {
		state, _ := hex.DecodeString(test.before)
		shiftRows(state)
		shifted := hex.EncodeToString(state)
		if shifted != test.after {
			t.Error(test.before, "should have been shifted to", test.after, "but was instead", shifted)
		}
	}
}

func TestShiftRowsReverse(t *testing.T) {
	for _, test := range shiftRowsTests {
		state, _ := hex.DecodeString(test.after)
		shiftRowsReverse(state)
		unshifted := hex.EncodeToString(state)
		if unshifted != test.before {
			t.Error(test.after, "should have been unshifted to", test.before, "but was instead", unshifted)
		}
	}
}
