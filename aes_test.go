package aes

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Example128() {
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	input := []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34}

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: 3925841d02dc09fbdc118597196a0b32
}

func ExampleAnother128() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	input, _ := hex.DecodeString("00112233445566778899aabbccddeeff")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: 69c4e0d86a7b0430d8cdb78070b4c55a
}

func Example192() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f1011121314151617")
	input, _ := hex.DecodeString("00112233445566778899aabbccddeeff")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: dda97ca4864cdfe06eaf70a0ec0d7191
}

func ExampleAnother192() {
	key, _ := hex.DecodeString("04050607090A0B0C0E0F10111314151618191A1B1D1E1F20")
	input, _ := hex.DecodeString("76777475F1F2F3F4F8F9E6E777707172")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: 5d1ef20dced6bcbc12131ac7c54788aa
}

func Example256() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	input, _ := hex.DecodeString("00112233445566778899aabbccddeeff")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: 8ea2b7ca516745bfeafc49904b496089
}

func ExampleAnother256() {
	key, _ := hex.DecodeString("08090A0B0D0E0F10121314151718191A1C1D1E1F21222324262728292B2C2D2E")
	input, _ := hex.DecodeString("069A007FC76A459F98BAF917FEDF9521")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Encode(input)))

	// Output: 080e9517eb1677719acf728086040ae3
}

func Example128Decode() {
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	input, _ := hex.DecodeString("3925841d02dc09fbdc118597196a0b32")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 3243f6a8885a308d313198a2e0370734
}

func ExampleAnother128Decode() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	input, _ := hex.DecodeString("69c4e0d86a7b0430d8cdb78070b4c55a")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 00112233445566778899aabbccddeeff
}

func Example192Decode() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f1011121314151617")
	input, _ := hex.DecodeString("dda97ca4864cdfe06eaf70a0ec0d7191")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 00112233445566778899aabbccddeeff
}

func ExampleAnother192Decode() {
	key, _ := hex.DecodeString("04050607090A0B0C0E0F10111314151618191A1B1D1E1F20")
	input, _ := hex.DecodeString("5d1ef20dced6bcbc12131ac7c54788aa")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 76777475f1f2f3f4f8f9e6e777707172
}

func Example256Decode() {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	input, _ := hex.DecodeString("8ea2b7ca516745bfeafc49904b496089")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 00112233445566778899aabbccddeeff
}

func ExampleAnother256Decode() {
	key, _ := hex.DecodeString("08090a0b0d0e0f10121314151718191a1c1d1e1f21222324262728292b2c2d2e")
	input, _ := hex.DecodeString("080e9517eb1677719acf728086040ae3")

	aes := &AES{key}
	fmt.Println(hex.EncodeToString(aes.Decode(input)))

	// Output: 069a007fc76a459f98baf917fedf9521
}

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
