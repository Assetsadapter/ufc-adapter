package addrdec

import (
	"github.com/blocktree/go-owaddress"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	alphabet = addressEncoder.BTCAlphabet

	UFCAddressPrefix_mainnet = "UFC"
	UFCAddressPrefix_testnet = "UFCT"
)

var (
	UFC_mainnetAddress = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "ripemd160", HashType: "sha512_ripemd160", HashLen: 20, Prefix: []byte{0x35}, Suffix: nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
	IsTestNet bool
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2() *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := UFC_mainnetAddress

	addreesPrefix := UFCAddressPrefix_mainnet
	if dec.IsTestNet {
		addreesPrefix = UFCAddressPrefix_testnet
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	base := addr[len(addreesPrefix):]
	//hash, _ := base58.Decode(base)
	//log.Infof("hash: %s", hex.EncodeToString(hash))
	return addressEncoder.AddressDecode(base, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := UFC_mainnetAddress

	addressPrefix := UFCAddressPrefix_mainnet
	if dec.IsTestNet {
		addressPrefix = UFCAddressPrefix_testnet
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	if len(hash) != cfg.HashLen {
		//公钥hash处理

		//UFC 地址生成文档
		//1， 获取到二进制公钥。
		//2， 对公钥进行sha512 获取sha512_data
		//3， 对sha512_data进行ripemd160 获取到ripe_data
		//4， Ripe_data 长度为20位。通过 0x35+ripe_data 组成新的长度为21位的base_data
		//5， 对base_data 进行ripemd160 获取到check_data
		//6， 将21位的base_data 和check_data的前4位拼接在一起，组成25位的end_data (base_data+check_data[:4])
		//7， 对end_data进行base58转码 获取到base58_data
		//8， 最终地址为固定字符串”UFC”+base58_data

		sha512hash := owcrypt.Hash(hash, 0, owcrypt.HASH_ALG_SHA512)
		//log.Infof("sha512hash = %s", hex.EncodeToString(sha512hash))
		ripe := owcrypt.Hash(sha512hash, 0, owcrypt.HASH_ALG_RIPEMD160)
		//log.Infof("ripe = %s", hex.EncodeToString(ripe))
		hash = ripe
	}

	//base := append(cfg.Prefix, hash...)
	//checksum := owcrypt.Hash(base, 0, owcrypt.HASH_ALG_RIPEMD160)[:4]
	//end := append(base, checksum...)
	////log.Infof("end = %s", hex.EncodeToString(end))
	//address := base58.Encode(end)

	address := addressEncoder.AddressEncode(hash, cfg)
	address = addressPrefix + address
	return address, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	valid, err := owaddress.Verify("ufc", address)
	if err != nil {
		return false
	}
	return valid
}
