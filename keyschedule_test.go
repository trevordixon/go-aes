package aes

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func printWord(word []byte) {
	fmt.Print(hex.EncodeToString(word))
}

func Example128KeyExpansion() {
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	next := generate128(key)

	for i := 0; i < 11; i++ {
		printWord(next())
	}

	// Output:
	// 2b7e151628aed2a6abf7158809cf4f3ca0fafe1788542cb123a339392a6c7605f2c295f27a96b9435935807a7359f67f3d80477d4716fe3e1e237e446d7a883bef44a541a8525b7fb671253bdb0bad00d4d1c6f87c839d87caf2b8bc11f915bc6d88a37a110b3efddbf98641ca0093fd4e54f70e5f5fc9f384a64fb24ea6dc4fead27321b58dbad2312bf5607f8d292fac7766f319fadc2128d12941575c006ed014f9a8c9ee2589e13f0cc8b6630ca6
}

func Example192KeyExpansion() {
	key := []byte{0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5, 0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b}
	next := generate192(key)

	for i := 0; i < 13; i++ {
		printWord(next())
	}

	// Output:
	// 8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7bfe0c91f72402f5a5ec12068e6c827f6b0e7a95b95c56fec24db7b4bd69b5411885a74796e92538fde75fad44bb095386485af05721efb14fa448f6d94d6dce24aa326360113b30e6a25e7ed583b1cf9a27f939436a94f767c0a69407d19da4e1ec1786eb6fa64971485f703222cb8755e26d135233f0b7b340beeb282f18a2596747d26b458c553ea7e1466c9411f1df821f750aad07d753ca4005388fcc5006282d166abc3ce7b5e98ba06f448c773c8ecc720401002202
}

func Example256KeyExpansion() {
	key := []byte{0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81, 0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4}
	next := generate256(key)

	for i := 0; i < 15; i++ {
		printWord(next())
	}

	// Output:
	// 603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff49ba354118e6925afa51a8b5f2067fcdea8b09c1a93d194cdbe49846eb75d5b9ad59aecb85bf3c917fee94248de8ebe96b5a9328a2678a647983122292f6c79b3812c81addadf48ba24360af2fab8b46498c5bfc9bebd198e268c3ba709e0421468007bacb2df331696e939e46c518d80c814e20476a9fb8a5025c02d59c58239de1369676ccc5a71fa2563959674ee155886ca5d2e2f31d77e0af1fa27cf73c3749c47ab18501ddae2757e4f7401905acafaaae3e4d59b349adf6acebd10190dfe4890d1e6188d0b046df344706c631e
}

/////////////////////////////////////////
////////// Rotate Word Tests ////////////
/////////////////////////////////////////

var rotateTests = []struct {
	word string
	out  string
}{
	{"09cf4f3c", "cf4f3c09"},
	{"2a6c7605", "6c76052a"},
	{"7359f67f", "59f67f73"},
	{"6d7a883b", "7a883b6d"},
	{"db0bad00", "0bad00db"},
	{"11f915bc", "f915bc11"},
	{"ca0093fd", "0093fdca"},
	{"4ea6dc4f", "a6dc4f4e"},
	{"7f8d292f", "8d292f7f"},
	{"575c006e", "5c006e57"},
	{"522c6b7b", "2c6b7b52"},
	{"5c56fec2", "56fec25c"},
	{"bb095386", "095386bb"},
	{"113b30e6", "3b30e611"},
}

func TestRotateWord(t *testing.T) {
	for _, test := range rotateTests {
		word, _ := hex.DecodeString(test.word)
		rotWord(word)
		rotatedWord := hex.EncodeToString(word)
		if rotatedWord != test.out {
			t.Error(test.word, "should have been rotated to", test.out, "but was instead", rotatedWord)
		}
	}
}

/////////////////////////////////////////
///////////// Sbox Tests ////////////////
/////////////////////////////////////////

var sboxTests = []struct {
	word string
	out  string
}{
	{"cf4f3c09", "8a84eb01"},
	{"6c76052a", "50386be5"},
	{"59f67f73", "cb42d28f"},
	{"7a883b6d", "dac4e23c"},
	{"0bad00db", "2b9563b9"},
	{"f915bc11", "99596582"},
	{"0093fdca", "63dc5474"},
	{"a6dc4f4e", "2486842f"},
	{"8d292f7f", "5da515d2"},
	{"5c006e57", "4a639f5b"},
}

func TestSubWord(t *testing.T) {
	for _, test := range sboxTests {
		word, _ := hex.DecodeString(test.word)
		subWord(word)
		sWord := hex.EncodeToString(word)
		if sWord != test.out {
			t.Error(test.word, "should have been substituted to", test.out, "but was instead", sWord)
		}
	}
}

/////////////////////////////////////////
/////////// Rcon Word Tests /////////////
/////////////////////////////////////////

var rconTests = []struct {
	word string
	i    int
	out  string
}{
	{"8a84eb01", 1, "8b84eb01"},
	{"50386be5", 2, "52386be5"},
	{"cb42d28f", 3, "cf42d28f"},
	{"dac4e23c", 4, "d2c4e23c"},
	{"2b9563b9", 5, "3b9563b9"},
	{"99596582", 6, "b9596582"},
	{"63dc5474", 7, "23dc5474"},
	{"2486842f", 8, "a486842f"},
	{"5da515d2", 9, "46a515d2"},
	{"4a639f5b", 10, "7c639f5b"},
}

func TestRcon(t *testing.T) {
	for _, test := range rconTests {
		word, _ := hex.DecodeString(test.word)
		rcon(word, test.i)
		xorWord := hex.EncodeToString(word)
		if xorWord != test.out {
			t.Error(test.word, "should have been xored to", test.out, "but was instead", xorWord)
		}
	}
}
