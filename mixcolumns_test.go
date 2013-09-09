package aes

import (
	"encoding/hex"
	"testing"
)

var mixColumnsTests = []struct {
	before string
	after  string
}{
	{"d4bf5d30e0b452aeb84111f11e2798e5", "046681e5e0cb199a48f8d37a2806264c"},
	{"49db873b453953897f02d2f177de961a", "584dcaf11b4b5aacdbe7caa81b6bb0e5"},
	{"acc1d6b8efb55a7b1323cfdf457311b5", "75ec0993200b633353c0cf7cbb25d0dc"},
}

func TestMixColumns(t *testing.T) {
	for _, test := range mixColumnsTests {
		state, _ := hex.DecodeString(test.before)
		mixColumns(state)
		mixed := hex.EncodeToString(state)
		if mixed != test.after {
			t.Error(test.before, "should have been mixed to", test.after, "but was instead", mixed)
		}
	}
}

func TestMixColumnsReverse(t *testing.T) {
	for _, test := range mixColumnsTests {
		state, _ := hex.DecodeString(test.after)
		mixColumnsReverse(state)
		unmixed := hex.EncodeToString(state)
		if unmixed != test.before {
			t.Error(test.after, "should have been unmixed to", test.before, "but was instead", unmixed)
		}
	}
}
