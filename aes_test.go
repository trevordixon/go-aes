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

var mixColumnsTests = []struct {
	before string
	after  string
}{
	{},
	{},
	{},
	{},
}

func TestMixColumns(t *testing.T) {

}
