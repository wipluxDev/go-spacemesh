package types
import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

// PrivKey = 616c69636500000000000000000000000000000000000000000000000000000038088e4c2ae82f5c45c6808a61a6490d3c612ce1da235714466fc748fbc4cbbb
// PubKey = 38088e4c2ae82f5c45c6808a61a6490d3c612ce1da235714466fc748fbc4cbbb

var txValueOldCoinTx11 = OldCoinTx{
	AccountNonce: 5,
	Recipient: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	GasLimit: 10,
	Fee: 1,
	Amount: 100,
}
var txCaseOldCoinTx11Ed = SignedTransaction{
/*type:*/ 0x06,0x00,0x00,0x00,
/*sign:*/ 0x5c,0x50,0xa9,0x24,0xe0,0x44,0xea,0x5b,0x3f,0x8d,0x84,0x80,0x99,0x5d,0x85,0xce,0x25,0x69,0x06,0xfc,0xc3,0xce,0x06,0xf8,0xc6,0x82,0x76,0xd9,0xa7,0xb5,0xce,0x40,0x16,0xea,0xe7,0x6a,0xbb,0x30,0x8a,0x13,0x17,0x10,0xca,0x11,0x14,0x56,0x02,0x7c,0xbe,0x1b,0xa2,0x30,0x3d,0x99,0xa8,0x67,0x49,0x42,0x0d,0xc5,0xa2,0x7e,0xdb,0x0a,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x64,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseOldCoinTx11EdPlus = SignedTransaction{
/*type:*/ 0x07,0x00,0x00,0x00,
/*sign:*/ 0x06,0x52,0x61,0xbd,0x72,0xb2,0x80,0x02,0x7f,0x6d,0xb1,0xfc,0x1b,0xb7,0x15,0xce,0x60,0x9b,0x06,0x85,0x22,0xf3,0x36,0x03,0x7b,0x68,0xe6,0x1d,0x4e,0x51,0x74,0xd8,0xed,0x35,0x7c,0xbc,0xff,0x32,0x88,0x3b,0x21,0x6d,0x72,0xbe,0x77,0x2b,0x31,0xca,0x64,0xad,0xed,0x9c,0xfe,0x26,0x9b,0x94,0x96,0xdd,0xc5,0x45,0x30,0xab,0x0c,0x05,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x64,
/*pkey:*/ 
}
var txValueOldCoinTx21 = OldCoinTx{
	AccountNonce: 6,
	Recipient: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	GasLimit: 10,
	Fee: 2,
	Amount: 101,
}
var txCaseOldCoinTx21Ed = SignedTransaction{
/*type:*/ 0x06,0x00,0x00,0x00,
/*sign:*/ 0xad,0xe1,0xd0,0x1b,0x91,0x02,0xd3,0x20,0x4d,0xb0,0x3b,0xb1,0xd1,0x4a,0xf9,0x96,0x07,0x05,0xe2,0xbb,0x92,0x93,0xab,0x17,0x18,0x23,0xe7,0xdd,0x54,0xf1,0xce,0xbb,0x54,0x38,0x7a,0xaa,0x48,0x28,0xcf,0x4f,0xd9,0xd9,0xc7,0x4a,0xc2,0xb6,0x9e,0xb9,0x19,0x6e,0x29,0xcb,0x57,0x2b,0x77,0x4f,0x75,0xd1,0xb2,0x7e,0xed,0x6c,0x8b,0x0a,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x65,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseOldCoinTx21EdPlus = SignedTransaction{
/*type:*/ 0x07,0x00,0x00,0x00,
/*sign:*/ 0x44,0x5e,0xdf,0xa6,0x90,0x31,0xb4,0xe6,0xbc,0xba,0x90,0xed,0x9c,0xc9,0x15,0x46,0x24,0x1b,0x42,0x17,0x16,0x6d,0xd8,0x9d,0x42,0x53,0xa2,0x85,0xfa,0x9c,0xff,0x97,0x43,0xd1,0xd2,0xe9,0xa2,0x78,0xf1,0x3a,0xc5,0xe1,0x41,0xab,0xcc,0x02,0xeb,0x0c,0xb6,0x00,0xe8,0xb2,0xfa,0xa8,0x7d,0x99,0x7e,0x18,0x28,0x77,0x63,0xba,0xc3,0x04,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x65,
/*pkey:*/ 
}
var txValueSimpleCoinTx11 = SimpleCoinTx{
	TTL: 0,
	Nonce: 5,
	Recipient: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 102,
	GasLimit: 1,
	GasPrice: 10,
}
var txCaseSimpleCoinTx11Ed = SignedTransaction{
/*type:*/ 0x00,0x00,0x00,0x00,
/*sign:*/ 0xb0,0x5b,0x3e,0x68,0x73,0x57,0xba,0x93,0x17,0x80,0x6b,0xd4,0x81,0x43,0x84,0x1f,0xc4,0xaf,0x3a,0xb0,0x75,0xcb,0xa4,0xba,0xe5,0xe9,0xb0,0x74,0x51,0x0e,0x29,0xc3,0x6f,0x9c,0xc2,0x5a,0x33,0x52,0xab,0xaf,0xa2,0x9e,0x96,0x9b,0x51,0xf5,0x18,0xe7,0x4a,0x83,0xda,0x5b,0xee,0xef,0x5f,0x52,0x91,0x3f,0x0d,0x2c,0x34,0x2c,0x81,0x07,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x66,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseSimpleCoinTx11EdPlus = SignedTransaction{
/*type:*/ 0x01,0x00,0x00,0x00,
/*sign:*/ 0xc8,0x4a,0x06,0x90,0x0a,0x04,0xec,0x7f,0x54,0x46,0x52,0xa4,0xba,0xd3,0x88,0xfb,0x56,0x3b,0x21,0xb5,0x5d,0x68,0x33,0x24,0xb3,0xd3,0xfb,0xdc,0x82,0xa9,0x91,0x56,0x58,0x8c,0x45,0x32,0xc8,0x40,0xac,0x3c,0xb4,0xb5,0x59,0xec,0x52,0x17,0x6e,0x30,0x0c,0x10,0x8a,0x6b,0x21,0x4d,0x7c,0xd1,0xcc,0xb0,0x92,0x21,0x0b,0x1f,0xda,0x0f,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x66,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,
/*pkey:*/ 
}
var txValueSimpleCoinTx21 = SimpleCoinTx{
	TTL: 0,
	Nonce: 8,
	Recipient: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 103,
	GasLimit: 10,
	GasPrice: 1,
}
var txCaseSimpleCoinTx21Ed = SignedTransaction{
/*type:*/ 0x00,0x00,0x00,0x00,
/*sign:*/ 0x49,0x7f,0x2a,0x0b,0xba,0x6f,0x14,0xa7,0x39,0xec,0x0e,0x66,0x14,0x2d,0x5f,0xda,0xb1,0xa1,0x2e,0xd1,0xba,0x7b,0x76,0x4d,0x0a,0xcd,0x2a,0x9d,0xee,0x05,0x83,0xf9,0xee,0x6b,0x87,0x2a,0x2e,0x9c,0x06,0x67,0x55,0xf3,0x1e,0x14,0x8a,0xf9,0x33,0x3a,0xb8,0x2c,0xac,0x9b,0x83,0xc2,0xec,0xfe,0x6d,0x85,0xc4,0x62,0x06,0x04,0x98,0x02,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x67,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseSimpleCoinTx21EdPlus = SignedTransaction{
/*type:*/ 0x01,0x00,0x00,0x00,
/*sign:*/ 0x67,0x21,0x92,0x57,0x76,0x96,0x90,0x70,0xfd,0x71,0xb9,0x0b,0x2c,0x14,0xb6,0xe4,0xfb,0xda,0x60,0x86,0xa4,0x11,0xaf,0xb2,0xba,0x8d,0xb4,0x8d,0x6d,0xc5,0x41,0xe8,0x6d,0xc9,0xef,0x35,0x50,0xf4,0x80,0x33,0xea,0x27,0x3e,0xe1,0x11,0x1a,0x3a,0x98,0x54,0xec,0xaa,0x7d,0x7c,0x51,0x56,0xb8,0xac,0x8e,0xf4,0xcc,0xc8,0xe8,0xff,0x05,
/*data:*/ 0x00,0x00,0x00,0x34,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x67,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,
/*pkey:*/ 
}
var txValueCallAppTx11 = CallAppTx{
	TTL: 0,
	Nonce: 5,
	AppAddress: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 104,
	GasLimit: 1,
	GasPrice: 10,
	CallData: []byte{0x01,0x02,0x03,0x04,0x05,0x06,},
}
var txCaseCallAppTx11Ed = SignedTransaction{
/*type:*/ 0x02,0x00,0x00,0x00,
/*sign:*/ 0x27,0xdf,0xfc,0xcc,0x7a,0x9d,0x63,0xb5,0x91,0xb8,0xe6,0x42,0x75,0x12,0xaf,0x11,0xfa,0x9c,0x19,0x1f,0x7f,0xa1,0x07,0xd0,0xda,0x48,0x5c,0xa5,0x9f,0x5d,0xba,0x79,0xff,0xa3,0xa7,0x43,0xba,0x01,0x32,0xd2,0x79,0x2b,0xce,0x37,0x9d,0xa8,0x2d,0x94,0xc0,0x31,0x1d,0x1c,0xda,0x96,0x67,0x70,0xc3,0x4d,0x22,0x58,0x31,0xe2,0xbc,0x0c,
/*data:*/ 0x00,0x00,0x00,0x40,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x68,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x06,0x01,0x02,0x03,0x04,0x05,0x06,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseCallAppTx11EdPlus = SignedTransaction{
/*type:*/ 0x03,0x00,0x00,0x00,
/*sign:*/ 0x3c,0x2e,0x2a,0x7b,0x8d,0xe5,0xc9,0x6f,0x91,0x00,0x07,0x23,0xf5,0xaf,0x29,0x8b,0x4b,0x58,0x48,0x73,0xc9,0xaa,0xd2,0xe7,0x4a,0xb2,0xd0,0x62,0x35,0x08,0x79,0x0f,0x43,0x80,0xd1,0x3b,0x79,0x84,0x61,0x45,0xcd,0x19,0x16,0x9f,0x4f,0x4c,0x0d,0x80,0x7e,0xf4,0x3e,0xc1,0x71,0x2d,0xae,0x65,0xe0,0xf6,0x08,0x29,0x11,0xd3,0x9c,0x06,
/*data:*/ 0x00,0x00,0x00,0x40,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x68,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x06,0x01,0x02,0x03,0x04,0x05,0x06,0x00,0x00,
/*pkey:*/ 
}
var txValueCallAppTx21 = CallAppTx{
	TTL: 0,
	Nonce: 8,
	AppAddress: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 105,
	GasLimit: 1,
	GasPrice: 5,
	CallData: nil,
}
var txCaseCallAppTx21Ed = SignedTransaction{
/*type:*/ 0x02,0x00,0x00,0x00,
/*sign:*/ 0xf9,0xb1,0x0e,0x18,0x89,0xdb,0x78,0x08,0x13,0x99,0xb8,0x93,0x31,0x1d,0xed,0x64,0xd7,0x16,0x4c,0x0a,0xe7,0x46,0x58,0x43,0xb6,0x81,0xc6,0xb5,0x8c,0x6f,0x05,0x32,0xa0,0xe5,0x0f,0x42,0xb1,0xc0,0xde,0x33,0xc8,0x64,0x03,0xf8,0xd6,0x80,0xc1,0x77,0xf4,0x56,0x93,0xa2,0x31,0xe0,0x8e,0xe3,0x08,0x1e,0x71,0x82,0x91,0xea,0x62,0x0f,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseCallAppTx21EdPlus = SignedTransaction{
/*type:*/ 0x03,0x00,0x00,0x00,
/*sign:*/ 0x36,0x6b,0xc3,0x28,0x63,0x3a,0xc1,0xd1,0xb4,0xd5,0xb7,0x2e,0xbc,0xf4,0xf3,0x50,0x3b,0x92,0x4f,0x76,0x2d,0xb9,0x46,0x59,0x80,0xa5,0x6d,0xde,0x5c,0xb6,0x6e,0xa9,0x9d,0xf0,0xc2,0x71,0x68,0x42,0x41,0x3e,0x47,0x12,0x90,0x97,0x10,0xa3,0xbd,0x45,0xce,0x03,0xc4,0x8b,0x81,0x28,0xe3,0x82,0xde,0x3c,0xb8,0x43,0xb6,0x93,0x74,0x07,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x00,
/*pkey:*/ 
}
var txValueCallAppTx31 = CallAppTx{
	TTL: 0,
	Nonce: 12,
	AppAddress: Address{0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,},
	Amount: 106,
	GasLimit: 1,
	GasPrice: 33,
	CallData: nil,
}
var txCaseCallAppTx31Ed = SignedTransaction{
/*type:*/ 0x02,0x00,0x00,0x00,
/*sign:*/ 0xa2,0x12,0x5e,0x48,0x1c,0x5a,0x14,0x3d,0x3d,0x32,0x70,0xba,0xb1,0x1d,0x8f,0x28,0x56,0x8b,0x3f,0x31,0xd2,0x55,0x8b,0xe3,0x7d,0xd2,0x80,0x11,0x5d,0x71,0x67,0x62,0x4b,0x42,0x0e,0xc9,0xb6,0xa4,0x5e,0xde,0x32,0x09,0xf6,0xc3,0x3b,0x36,0xf1,0x9d,0xc4,0x40,0xb8,0xbc,0xb0,0x92,0x90,0x03,0xe2,0x04,0x29,0x53,0x92,0xb5,0xd9,0x08,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x0c,0x00,0x00,0x00,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x21,0x00,0x00,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseCallAppTx31EdPlus = SignedTransaction{
/*type:*/ 0x03,0x00,0x00,0x00,
/*sign:*/ 0xa8,0xe7,0x0c,0x76,0x33,0x15,0xe8,0x09,0xd7,0x41,0x46,0x2e,0xd6,0xf5,0xfc,0xb5,0x77,0x00,0xc7,0x61,0x85,0x1c,0x71,0x98,0xd1,0x68,0x43,0x2b,0x45,0xba,0x9a,0x71,0xee,0xf3,0xe0,0x6a,0x0b,0x81,0xd1,0xb4,0x63,0x40,0xcf,0xa6,0xfb,0x73,0x14,0xd4,0x87,0xe9,0xfa,0x9a,0x02,0x8a,0xe8,0xb3,0x00,0xc4,0x57,0x69,0x1e,0x3e,0x2a,0x06,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x0c,0x00,0x00,0x00,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6a,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x21,0x00,0x00,0x00,0x00,
/*pkey:*/ 
}
var txValueSpawnAppTx11 = SpawnAppTx{
	TTL: 0,
	Nonce: 5,
	AppAddress: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 107,
	GasLimit: 1,
	GasPrice: 10,
	CallData: []byte{0x01,0x02,0x03,0x04,0x05,0x06,},
}
var txCaseSpawnAppTx11Ed = SignedTransaction{
/*type:*/ 0x04,0x00,0x00,0x00,
/*sign:*/ 0x7f,0xff,0x9c,0xcd,0xdf,0xe8,0xce,0x54,0x41,0x7b,0xdb,0xb2,0x7d,0xdd,0x5c,0x5e,0x04,0xc2,0xdf,0x87,0x7e,0xa5,0x8f,0x2d,0x7a,0xca,0x92,0xfc,0xe5,0xe5,0xa5,0x8b,0xee,0xc0,0xb1,0xf0,0x8c,0xac,0x25,0x27,0x12,0x09,0x2a,0xb4,0xfd,0x87,0xa3,0xd1,0xf9,0x0c,0x5d,0x5e,0x97,0x2c,0x31,0x1f,0x62,0x35,0x2a,0xbf,0x55,0xda,0x3d,0x02,
/*data:*/ 0x00,0x00,0x00,0x40,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6b,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x06,0x01,0x02,0x03,0x04,0x05,0x06,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseSpawnAppTx11EdPlus = SignedTransaction{
/*type:*/ 0x05,0x00,0x00,0x00,
/*sign:*/ 0x32,0xbb,0xd7,0xd8,0x12,0xc5,0x69,0x6d,0x48,0x4b,0x9f,0x3c,0x5e,0xbc,0x9c,0x83,0xf8,0xc4,0x2d,0x34,0x26,0x22,0xae,0x9b,0x37,0xb4,0x59,0x9f,0xb0,0xc8,0xb1,0x5b,0x24,0x8c,0x21,0xd7,0xf6,0xde,0x7e,0xd9,0x7f,0x7a,0xab,0x51,0x31,0xe2,0x01,0x45,0xf4,0x06,0x43,0xb6,0x2f,0x0c,0xce,0x6d,0x30,0x36,0x60,0x46,0x13,0xbd,0x70,0x03,
/*data:*/ 0x00,0x00,0x00,0x40,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6b,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a,0x00,0x00,0x00,0x06,0x01,0x02,0x03,0x04,0x05,0x06,0x00,0x00,
/*pkey:*/ 
}
var txValueSpawnAppTx21 = SpawnAppTx{
	TTL: 0,
	Nonce: 8,
	AppAddress: Address{0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,},
	Amount: 108,
	GasLimit: 1,
	GasPrice: 5,
	CallData: nil,
}
var txCaseSpawnAppTx21Ed = SignedTransaction{
/*type:*/ 0x04,0x00,0x00,0x00,
/*sign:*/ 0xf0,0x4a,0x03,0xb7,0x11,0xd2,0xaf,0x4f,0xc6,0xb3,0x61,0x46,0x96,0x86,0x41,0x58,0x63,0xa0,0x34,0x44,0x56,0xb7,0x3e,0x77,0x75,0x19,0xe5,0x9a,0xf5,0xe9,0xbd,0xe9,0x6a,0xab,0x3b,0xa7,0x1a,0xc3,0x5f,0xa5,0x59,0x83,0xb1,0x4d,0xa4,0x73,0x3a,0xd7,0x61,0x52,0xde,0xf0,0xcb,0x4c,0xd3,0x8b,0x8c,0x7d,0x69,0xbe,0x48,0x48,0xeb,0x09,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6c,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseSpawnAppTx21EdPlus = SignedTransaction{
/*type:*/ 0x05,0x00,0x00,0x00,
/*sign:*/ 0xf4,0xfc,0x47,0xd7,0x28,0x3d,0xfe,0x3c,0x21,0xb6,0x2a,0x39,0x64,0x25,0xb0,0xd7,0x7d,0x63,0x4d,0x40,0xea,0x0e,0xa0,0xef,0x1f,0x1d,0xa9,0xd6,0xef,0x24,0xc0,0x48,0x5f,0xdb,0x3b,0x4c,0xf3,0x61,0xf9,0x77,0x38,0x13,0xd2,0xeb,0x82,0x22,0x52,0xae,0x83,0x8c,0xe5,0xf6,0x05,0xe3,0x4e,0xfe,0xc5,0x85,0x09,0xcf,0x74,0xb4,0x2b,0x00,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x08,0x00,0x00,0x00,0x26,0x4d,0xd0,0x80,0xe4,0xb3,0xe7,0x30,0x16,0xbd,0xf6,0x57,0x4b,0x4c,0xd8,0x7e,0xaa,0x69,0x7f,0x5e,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6c,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05,0x00,0x00,0x00,0x00,
/*pkey:*/ 
}
var txValueSpawnAppTx31 = SpawnAppTx{
	TTL: 0,
	Nonce: 12,
	AppAddress: Address{0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,},
	Amount: 109,
	GasLimit: 1,
	GasPrice: 33,
	CallData: nil,
}
var txCaseSpawnAppTx31Ed = SignedTransaction{
/*type:*/ 0x04,0x00,0x00,0x00,
/*sign:*/ 0x47,0xb1,0x76,0xa4,0x66,0xb0,0x91,0x69,0xc4,0xe0,0x14,0xc5,0x3d,0x75,0x1c,0xbc,0x04,0x07,0x21,0xa7,0xae,0xc6,0xf1,0xb2,0x62,0x30,0x5c,0x1a,0x3a,0x9d,0x5b,0xea,0xa0,0xd6,0xaa,0x72,0xb9,0xb3,0xb0,0xf3,0xe8,0xa4,0xaf,0xe3,0xdd,0x18,0xc0,0x89,0xb2,0x61,0x57,0x75,0x18,0x8c,0xe1,0x58,0x54,0x58,0x3a,0xe8,0x0b,0xc2,0xdb,0x00,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x0c,0x00,0x00,0x00,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6d,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x21,0x00,0x00,0x00,0x00,
/*pkey:*/ 0x00,0x00,0x00,0x20,0x38,0x08,0x8e,0x4c,0x2a,0xe8,0x2f,0x5c,0x45,0xc6,0x80,0x8a,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,
}
var txCaseSpawnAppTx31EdPlus = SignedTransaction{
/*type:*/ 0x05,0x00,0x00,0x00,
/*sign:*/ 0xb5,0x1f,0x0f,0xc0,0xf8,0x40,0xc2,0x4d,0xed,0xdc,0x07,0x18,0xc3,0xb9,0x88,0x2c,0x15,0xde,0x98,0xf1,0xf6,0x89,0x31,0x7e,0xfc,0xab,0x2f,0x84,0xf6,0xe2,0x77,0x34,0x3a,0xbc,0x39,0xcc,0x47,0xe6,0x9b,0xd8,0x7c,0xd0,0xfe,0x7e,0xb5,0x78,0xf1,0x8a,0xbc,0x23,0x89,0x65,0x0c,0x0a,0x5d,0xa3,0xa1,0xea,0xef,0x3e,0xba,0xca,0x20,0x02,
/*data:*/ 0x00,0x00,0x00,0x38,0x00,0x00,0x00,0x00,0x0c,0x00,0x00,0x00,0x61,0xa6,0x49,0x0d,0x3c,0x61,0x2c,0xe1,0xda,0x23,0x57,0x14,0x46,0x6f,0xc7,0x48,0xfb,0xc4,0xcb,0xbb,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x6d,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x21,0x00,0x00,0x00,0x00,
/*pkey:*/ 
}
var txCases1 = []struct{Tx interface{}; Signed SignedTransaction}{
	{txValueOldCoinTx11,txCaseOldCoinTx11Ed},
	{txValueOldCoinTx11,txCaseOldCoinTx11EdPlus},
	{txValueOldCoinTx21,txCaseOldCoinTx21Ed},
	{txValueOldCoinTx21,txCaseOldCoinTx21EdPlus},
	{txValueSimpleCoinTx11,txCaseSimpleCoinTx11Ed},
	{txValueSimpleCoinTx11,txCaseSimpleCoinTx11EdPlus},
	{txValueSimpleCoinTx21,txCaseSimpleCoinTx21Ed},
	{txValueSimpleCoinTx21,txCaseSimpleCoinTx21EdPlus},
	{txValueCallAppTx11,txCaseCallAppTx11Ed},
	{txValueCallAppTx11,txCaseCallAppTx11EdPlus},
	{txValueCallAppTx21,txCaseCallAppTx21Ed},
	{txValueCallAppTx21,txCaseCallAppTx21EdPlus},
	{txValueCallAppTx31,txCaseCallAppTx31Ed},
	{txValueCallAppTx31,txCaseCallAppTx31EdPlus},
	{txValueSpawnAppTx11,txCaseSpawnAppTx11Ed},
	{txValueSpawnAppTx11,txCaseSpawnAppTx11EdPlus},
	{txValueSpawnAppTx21,txCaseSpawnAppTx21Ed},
	{txValueSpawnAppTx21,txCaseSpawnAppTx21EdPlus},
	{txValueSpawnAppTx31,txCaseSpawnAppTx31Ed},
	{txValueSpawnAppTx31,txCaseSpawnAppTx31EdPlus},
}
	
func TestBinaryTransactions1(t *testing.T) {
	var old bool
	old,EnableTransactionPruning = EnableTransactionPruning, false
	defer func(){
		EnableTransactionPruning = old
	}()
	for _,v := range txCases1 {
		tx, err := v.Signed.Decode()
		require.NoError(t, err)
		val := reflect.ValueOf(v.Tx)
		b := reflect.New(val.Type())
		ok := tx.Extract(b.Interface())
		require.True(t, ok)
		require.Equal(t, v.Tx, b.Elem().Interface())
	}
}	