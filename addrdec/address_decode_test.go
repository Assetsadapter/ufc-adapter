package addrdec

import (
	"encoding/hex"
	"testing"

	"github.com/blocktree/go-owcrypt"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	Default.IsTestNet = false

	p2pk, _ := hex.DecodeString("cf036690a6fbbcdd9dfdd6249e1d121bec1eacd8")
	p2pkAddr, _ := Default.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)
}

func TestAddressDecoder_PubKeyEncode(t *testing.T) {
	Default.IsTestNet = false

	p2pk, _ := hex.DecodeString("02c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cfeb05f9d2")
	p2pkAddr, _ := Default.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)
	//UFCTNa5ZMhvFYXSYN4E2sAKqDVBKZgU9AGEBfZ

	p2pk2, _ := hex.DecodeString("0282ddc719bdb0f82ddb184de7e00ce1f8d01e1abf4e91dd6b419f4c548725c11241efb876")
	p2pkAddr2, _ := Default.AddressEncode(p2pk2)
	t.Logf("p2pkAddr2: %s", p2pkAddr2)
	//UFCNenZ8kD5AHghHhYRB3PWTABZjGFHTVGG4C
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	Default.IsTestNet = false

	p2pkAddr := "UFC56mC7HSbM37S7RoG6sktavZYymSRqhSiazctQp8DWW21jNWB6c"
	p2pkHash, _ := Default.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))
}

func TestDecompressPubKey(t *testing.T) {
	pub, _ := hex.DecodeString("03a2147994c34ec6ac3ba4d0737e672002a832f2cf050f0d2cffc7906ad8a1e7b2")
	uncompessedPublicKey := owcrypt.PointDecompress(pub, owcrypt.ECC_CURVE_SECP256K1)
	t.Logf("pub: %s", hex.EncodeToString(uncompessedPublicKey))
}
