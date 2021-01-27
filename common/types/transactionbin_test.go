package types

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

// PrivKey = 616c69636500000000000000000000000000000000000000000000000000000038088e4c2ae82f5c45c6808a61a6490d3c612ce1da235714466fc748fbc4cbbb
// PubKey = 38088e4c2ae82f5c45c6808a61a6490d3c612ce1da235714466fc748fbc4cbbb

var txValueOldCoinTx1 = OldCoinTx{
	AccountNonce: 5,
	Recipient:    Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	GasLimit:     10,
	Fee:          1,
	Amount:       100,
}
var txCaseOldCoinTx1Ed = SignedTransaction{
	/*type:*/ 0x06, 0x00, 0x00, 0x00,
	/*sign:*/ 0x9b, 0x76, 0x94, 0xb8, 0x5a, 0x99, 0xb1, 0xf8, 0x4b, 0xe1, 0x0e, 0x18, 0x2e, 0x35, 0xc3, 0xf6, 0xb1, 0xcf, 0x6f, 0x5a, 0x91, 0x60, 0x99, 0xd9, 0x66, 0xab, 0x1b, 0x04, 0xbb, 0x27, 0x45, 0x5c, 0xbd, 0x95, 0x98, 0xb8, 0x82, 0x40, 0x08, 0x84, 0x0c, 0x28, 0x64, 0x0d, 0x0f, 0x2d, 0xfa, 0x37, 0x5b, 0xc3, 0xfa, 0xb8, 0xcc, 0x63, 0x4a, 0x9f, 0xb2, 0x03, 0xce, 0x25, 0x67, 0xd4, 0x3f, 0x0b,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseOldCoinTx1EdPlus = SignedTransaction{
	/*type:*/ 0x07, 0x00, 0x00, 0x00,
	/*sign:*/ 0xc7, 0x9a, 0x78, 0x6b, 0xd4, 0x9f, 0x96, 0xb0, 0x76, 0x82, 0x4e, 0x3d, 0x6e, 0x32, 0x88, 0x41, 0xc3, 0x44, 0xf1, 0xe6, 0xbf, 0xc2, 0x58, 0x0d, 0xdc, 0xf1, 0x1c, 0x4e, 0xf5, 0xd1, 0x86, 0xcf, 0x93, 0x85, 0x2f, 0xb2, 0x70, 0x92, 0x63, 0xa6, 0x73, 0x67, 0xfb, 0xd8, 0x9c, 0xc1, 0xe5, 0x40, 0x2e, 0x3f, 0x72, 0xfe, 0x0a, 0xbb, 0xf2, 0x8a, 0x9d, 0x27, 0xd9, 0x72, 0x86, 0x94, 0x1e, 0x0e,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	/*pkey:*/
}
var txValueOldCoinTx2 = OldCoinTx{
	AccountNonce: 6,
	Recipient:    Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	GasLimit:     10,
	Fee:          2,
	Amount:       101,
}
var txCaseOldCoinTx2Ed = SignedTransaction{
	/*type:*/ 0x06, 0x00, 0x00, 0x00,
	/*sign:*/ 0x99, 0x7b, 0x47, 0x0c, 0xfd, 0xee, 0x7f, 0xbf, 0x19, 0x85, 0x1f, 0xaf, 0xfb, 0xc8, 0x17, 0x0d, 0xff, 0x7a, 0x82, 0x94, 0xf8, 0xc2, 0x37, 0xf7, 0x4a, 0x39, 0x96, 0x2f, 0x91, 0x58, 0x0f, 0x68, 0x0e, 0x06, 0xc2, 0x69, 0x63, 0x61, 0x65, 0x7d, 0xd6, 0xa4, 0x31, 0xa2, 0xb7, 0xa3, 0x47, 0x5a, 0xde, 0xb7, 0x83, 0xcd, 0x22, 0xda, 0x7f, 0x0e, 0xf1, 0xe1, 0xd8, 0x22, 0x8e, 0xd0, 0x2f, 0x08,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseOldCoinTx2EdPlus = SignedTransaction{
	/*type:*/ 0x07, 0x00, 0x00, 0x00,
	/*sign:*/ 0x6a, 0xb5, 0xb4, 0x5a, 0xfe, 0xb6, 0x47, 0x52, 0xf9, 0x29, 0x12, 0x4e, 0x74, 0x26, 0x35, 0x56, 0x6c, 0xb3, 0x2d, 0xdd, 0x45, 0xcb, 0x06, 0x21, 0x97, 0x9d, 0x0e, 0x88, 0xf1, 0x83, 0xfb, 0xf2, 0x57, 0x24, 0x4d, 0x3b, 0xc2, 0x3e, 0x69, 0x95, 0x83, 0x80, 0xee, 0x02, 0x8c, 0x6a, 0xe2, 0x9f, 0x41, 0xbd, 0x4a, 0xd8, 0xae, 0x01, 0xa2, 0x21, 0xd2, 0x4b, 0x2b, 0x45, 0x90, 0x4f, 0x27, 0x00,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x65,
	/*pkey:*/
}
var txValueSimpleCoinTx1 = SimpleCoinTx{
	TTL:       0,
	Nonce:     5,
	Recipient: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:    102,
	GasLimit:  1,
	GasPrice:  10,
}
var txCaseSimpleCoinTx1Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x00,
	/*sign:*/ 0xdc, 0xe4, 0x44, 0x01, 0x2e, 0xd8, 0x41, 0xcf, 0xc8, 0x26, 0x62, 0x4d, 0xe9, 0x8c, 0x7e, 0x43, 0xc4, 0x26, 0x28, 0xa5, 0x15, 0x38, 0x5b, 0xc3, 0x4c, 0xc3, 0xaf, 0xc1, 0xc7, 0x40, 0x55, 0x01, 0x3d, 0x6c, 0xa7, 0x11, 0xde, 0x27, 0x3f, 0x3a, 0xe7, 0x26, 0x77, 0xd3, 0x2a, 0x6f, 0x22, 0x54, 0x52, 0x0b, 0x93, 0xaf, 0x1a, 0xa9, 0x81, 0x4a, 0x02, 0x57, 0x03, 0xcd, 0x23, 0x21, 0x3b, 0x0c,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSimpleCoinTx1EdPlus = SignedTransaction{
	/*type:*/ 0x01, 0x00, 0x00, 0x00,
	/*sign:*/ 0x25, 0x25, 0x46, 0x83, 0x08, 0xd6, 0xda, 0x59, 0x56, 0x08, 0x7b, 0x15, 0xbe, 0x6d, 0x42, 0xc8, 0x76, 0x07, 0x37, 0x98, 0xcd, 0xad, 0xfd, 0xbd, 0x3c, 0x36, 0xce, 0x03, 0x2d, 0x12, 0x5c, 0xd1, 0x59, 0x2c, 0xa2, 0x42, 0xd6, 0x88, 0x2f, 0x8a, 0xf0, 0xc7, 0x9a, 0x1a, 0xa3, 0x2e, 0x8a, 0x16, 0xeb, 0xe7, 0x41, 0xed, 0x40, 0xb1, 0x8c, 0x83, 0x97, 0x60, 0x02, 0xf4, 0x51, 0x20, 0x7f, 0x03,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a,
	/*pkey:*/
}
var txValueSimpleCoinTx2 = SimpleCoinTx{
	TTL:       0,
	Nonce:     8,
	Recipient: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:    103,
	GasLimit:  10,
	GasPrice:  1,
}
var txCaseSimpleCoinTx2Ed = SignedTransaction{
	/*type:*/ 0x00, 0x00, 0x00, 0x00,
	/*sign:*/ 0x7a, 0xa4, 0xe9, 0xd2, 0x25, 0x96, 0xf4, 0x39, 0x05, 0xf9, 0x18, 0x07, 0x18, 0xa0, 0x2b, 0x68, 0x48, 0x5e, 0x8d, 0xe2, 0xa7, 0x67, 0x1a, 0x6a, 0xe6, 0x0b, 0xef, 0x41, 0x88, 0x3e, 0xfc, 0x04, 0xf6, 0xe9, 0x2e, 0xaa, 0x5a, 0x81, 0x05, 0x4a, 0xf6, 0xbe, 0x30, 0x41, 0xe5, 0x55, 0x36, 0x2f, 0x3f, 0xf7, 0x5f, 0xe3, 0x43, 0xc3, 0x43, 0xd0, 0xcf, 0x38, 0xea, 0x4f, 0x6a, 0x87, 0xad, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSimpleCoinTx2EdPlus = SignedTransaction{
	/*type:*/ 0x01, 0x00, 0x00, 0x00,
	/*sign:*/ 0xad, 0x9e, 0xef, 0xfb, 0x28, 0x09, 0xb4, 0x0e, 0x64, 0x1c, 0x3c, 0xe1, 0xd7, 0xcf, 0x0d, 0x66, 0x4c, 0x87, 0xeb, 0x8e, 0xa5, 0x87, 0xa8, 0x30, 0x95, 0xf6, 0x22, 0xe1, 0xc7, 0xf6, 0x24, 0xf2, 0x49, 0x71, 0xde, 0x43, 0x16, 0xca, 0xcb, 0x68, 0x01, 0xa5, 0xa7, 0x90, 0x9e, 0x4c, 0x83, 0xe5, 0x00, 0xb1, 0xc8, 0xdc, 0x94, 0x4f, 0x08, 0x62, 0x56, 0xd6, 0xa6, 0x06, 0xfb, 0xe2, 0x6c, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	/*pkey:*/
}
var txValueCallAppTx1 = CallAppTx{
	TTL:        0,
	Nonce:      5,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     104,
	GasLimit:   1,
	GasPrice:   10,
	CallData:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
}
var txCaseCallAppTx1Ed = SignedTransaction{
	/*type:*/ 0x02, 0x00, 0x00, 0x00,
	/*sign:*/ 0x2d, 0x1a, 0xe0, 0x2f, 0x93, 0x8c, 0x69, 0x0e, 0xbd, 0x5c, 0x6b, 0xe2, 0xbf, 0x25, 0x74, 0x3a, 0x23, 0x33, 0x97, 0x79, 0xbe, 0x66, 0x8b, 0x3c, 0x5e, 0xc8, 0xb6, 0x0d, 0x71, 0xcd, 0xce, 0x52, 0xc5, 0xce, 0x94, 0x0a, 0xfa, 0x9d, 0x29, 0x76, 0x43, 0x7a, 0xdd, 0x03, 0xb9, 0xfa, 0xf7, 0x78, 0xce, 0x1b, 0x6d, 0x4d, 0x23, 0xc3, 0x0a, 0xf4, 0x2a, 0xd7, 0x5f, 0x82, 0x80, 0x24, 0x62, 0x0a,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x1a, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx1EdPlus = SignedTransaction{
	/*type:*/ 0x03, 0x00, 0x00, 0x00,
	/*sign:*/ 0xa0, 0x48, 0xea, 0xaf, 0x78, 0x4d, 0x0c, 0x6f, 0xeb, 0xe7, 0x97, 0x26, 0x52, 0x0d, 0x8c, 0x32, 0x70, 0xef, 0x54, 0x61, 0xca, 0x8c, 0xca, 0x5d, 0xc8, 0xa8, 0x48, 0x14, 0x05, 0x1e, 0x67, 0x51, 0x16, 0x6d, 0x0f, 0xd9, 0xe6, 0x7e, 0x72, 0x5b, 0x4a, 0x60, 0x77, 0xe3, 0x26, 0x36, 0xb9, 0x31, 0x9f, 0x9e, 0xbb, 0x5a, 0x1a, 0x77, 0x5a, 0x4d, 0x2b, 0x10, 0x44, 0x8b, 0xaf, 0x95, 0x0d, 0x0f,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x1a, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/
}
var txValueCallAppTx2 = CallAppTx{
	TTL:        0,
	Nonce:      8,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     105,
	GasLimit:   1,
	GasPrice:   5,
	CallData:   []byte{},
}
var txCaseCallAppTx2Ed = SignedTransaction{
	/*type:*/ 0x02, 0x00, 0x00, 0x00,
	/*sign:*/ 0x74, 0x99, 0xb2, 0xb7, 0x1b, 0xa8, 0x27, 0x12, 0x64, 0xa1, 0xff, 0x0b, 0x50, 0xde, 0x27, 0x9e, 0xbf, 0xd7, 0x69, 0x38, 0xc9, 0xe6, 0x31, 0xd2, 0x03, 0x31, 0x60, 0xd8, 0x6c, 0xea, 0xff, 0xbd, 0xc0, 0x99, 0xee, 0x14, 0x66, 0x1e, 0x71, 0x0a, 0xd8, 0x11, 0x23, 0xf6, 0x7b, 0x55, 0xc5, 0x37, 0xca, 0x32, 0xbc, 0x08, 0xb9, 0x64, 0x78, 0x14, 0xa5, 0x50, 0xdd, 0x69, 0xa6, 0x11, 0x0e, 0x09,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x14, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx2EdPlus = SignedTransaction{
	/*type:*/ 0x03, 0x00, 0x00, 0x00,
	/*sign:*/ 0xe2, 0xc5, 0xa3, 0x55, 0x18, 0x9e, 0xc6, 0x34, 0xf6, 0x42, 0x07, 0xfe, 0x2c, 0x2f, 0x31, 0xcc, 0xc4, 0xc2, 0x97, 0xa7, 0x91, 0x81, 0xb9, 0x51, 0xf8, 0x52, 0xc6, 0x29, 0x53, 0xb9, 0xc2, 0x79, 0xfb, 0xe0, 0xaa, 0x2e, 0xb8, 0xd3, 0xfc, 0x47, 0x7a, 0x8f, 0xaf, 0x5d, 0x00, 0x84, 0x3a, 0x47, 0xbe, 0xbd, 0x20, 0xc8, 0x92, 0x24, 0xc2, 0x5a, 0x7a, 0x61, 0xa7, 0x67, 0x41, 0xc2, 0x56, 0x06,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x14, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e,
	/*pkey:*/
}
var txValueCallAppTx3 = CallAppTx{
	TTL:        0,
	Nonce:      12,
	AppAddress: Address{0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb},
	Amount:     106,
	GasLimit:   1,
	GasPrice:   33,
	CallData:   []byte{},
}
var txCaseCallAppTx3Ed = SignedTransaction{
	/*type:*/ 0x02, 0x00, 0x00, 0x00,
	/*sign:*/ 0xa1, 0x87, 0x3f, 0x1f, 0x07, 0x20, 0xdd, 0x9b, 0x55, 0x3e, 0x7d, 0xeb, 0x4e, 0x08, 0x16, 0x84, 0x44, 0x73, 0xad, 0xf7, 0x7c, 0xca, 0x7f, 0xe1, 0x9c, 0xab, 0x6a, 0x68, 0xde, 0x6b, 0x16, 0x90, 0x79, 0x37, 0x4d, 0x22, 0xe7, 0xe3, 0x25, 0x70, 0x7a, 0x82, 0x28, 0x11, 0x74, 0xc3, 0xd3, 0xd2, 0xc5, 0xb7, 0x70, 0x48, 0x14, 0x51, 0x09, 0xbc, 0xae, 0x16, 0xe4, 0x5b, 0x63, 0xe4, 0xbc, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x14, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseCallAppTx3EdPlus = SignedTransaction{
	/*type:*/ 0x03, 0x00, 0x00, 0x00,
	/*sign:*/ 0x21, 0x1e, 0xcf, 0x47, 0xc6, 0x64, 0xf6, 0x21, 0xab, 0x9a, 0xe5, 0x32, 0x61, 0x8e, 0x7f, 0xb9, 0xe0, 0xe0, 0x88, 0x6e, 0x08, 0x22, 0x44, 0xed, 0x4b, 0xe7, 0xa8, 0xab, 0xa4, 0x22, 0x28, 0x65, 0x5f, 0x56, 0x52, 0x71, 0x84, 0x4f, 0xcf, 0x2f, 0x99, 0x74, 0x58, 0xc7, 0x6f, 0x9c, 0x77, 0x01, 0x40, 0xde, 0x6a, 0x0f, 0xb5, 0x5f, 0x2a, 0xa2, 0x89, 0x5e, 0x00, 0x91, 0x66, 0xd1, 0x8c, 0x04,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x14, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
	/*pkey:*/
}
var txValueSpawnAppTx1 = SpawnAppTx{
	TTL:        0,
	Nonce:      5,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     107,
	GasLimit:   1,
	GasPrice:   10,
	CallData:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
}
var txCaseSpawnAppTx1Ed = SignedTransaction{
	/*type:*/ 0x04, 0x00, 0x00, 0x00,
	/*sign:*/ 0x04, 0xff, 0x66, 0xcd, 0x9a, 0x45, 0x38, 0x06, 0x1e, 0x29, 0x78, 0xd4, 0x4a, 0x4b, 0x2f, 0x76, 0x10, 0x1e, 0x12, 0xf6, 0x91, 0x69, 0x42, 0x41, 0xb5, 0x71, 0xc2, 0xa1, 0xee, 0x42, 0x44, 0xbc, 0x2a, 0x8b, 0x6b, 0x2a, 0xf1, 0x34, 0xc5, 0x82, 0x98, 0xfe, 0x44, 0xcb, 0xe5, 0x00, 0x17, 0x2c, 0x5f, 0xeb, 0x2a, 0x53, 0x6e, 0x05, 0x3c, 0xd4, 0x0c, 0xbd, 0x88, 0x19, 0x10, 0xf0, 0x19, 0x05,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x1a, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx1EdPlus = SignedTransaction{
	/*type:*/ 0x05, 0x00, 0x00, 0x00,
	/*sign:*/ 0x55, 0x72, 0x7f, 0x91, 0x18, 0x32, 0xce, 0x59, 0xaa, 0xcc, 0x00, 0xce, 0xe0, 0xf8, 0x24, 0x8c, 0x34, 0xc3, 0x76, 0x23, 0x4d, 0xea, 0x74, 0xcf, 0x53, 0x0a, 0x07, 0x40, 0x6b, 0x55, 0x66, 0xfe, 0x06, 0xef, 0x7f, 0x50, 0x04, 0x34, 0x75, 0x77, 0xc4, 0x6f, 0xf0, 0x79, 0x7d, 0x65, 0x77, 0xcc, 0xa9, 0x80, 0x5b, 0xa3, 0xa5, 0x3a, 0x11, 0xeb, 0x91, 0x41, 0xb3, 0x1f, 0x36, 0xb3, 0x0c, 0x07,
	/*data:*/ 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x1a, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x00, 0x00,
	/*pkey:*/
}
var txValueSpawnAppTx2 = SpawnAppTx{
	TTL:        0,
	Nonce:      8,
	AppAddress: Address{0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e},
	Amount:     108,
	GasLimit:   1,
	GasPrice:   5,
	CallData:   []byte{},
}
var txCaseSpawnAppTx2Ed = SignedTransaction{
	/*type:*/ 0x04, 0x00, 0x00, 0x00,
	/*sign:*/ 0x85, 0x24, 0xe7, 0x6a, 0x4f, 0xb4, 0xfd, 0x72, 0xd8, 0x2a, 0x66, 0x3f, 0x9a, 0x2f, 0x90, 0x1d, 0x18, 0x88, 0x20, 0xc5, 0x7c, 0x1d, 0x08, 0x2d, 0x9b, 0xd9, 0xda, 0x98, 0xb9, 0x4d, 0x4a, 0x51, 0x91, 0x55, 0x02, 0xc7, 0xd5, 0x75, 0x68, 0x4b, 0xc2, 0xc8, 0xcd, 0x0f, 0xb8, 0x53, 0xa7, 0xae, 0xc0, 0xb9, 0xef, 0xbe, 0x2e, 0x4b, 0xd6, 0xcc, 0xf8, 0x8f, 0x74, 0xe8, 0x6f, 0x80, 0xc7, 0x03,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x14, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx2EdPlus = SignedTransaction{
	/*type:*/ 0x05, 0x00, 0x00, 0x00,
	/*sign:*/ 0x98, 0x70, 0xb8, 0x5b, 0x2f, 0xa1, 0xf6, 0xae, 0xbc, 0x7a, 0xe2, 0xee, 0x7e, 0x32, 0x7e, 0x7b, 0x60, 0x6c, 0xf3, 0x77, 0x2d, 0xaa, 0x4f, 0xe2, 0x52, 0x61, 0x54, 0xf9, 0x72, 0x31, 0xa6, 0x3d, 0xc0, 0x04, 0x64, 0xc8, 0xca, 0xc1, 0x57, 0xbf, 0xc7, 0x62, 0x8c, 0x6c, 0x6c, 0x34, 0xf5, 0xd1, 0x51, 0x39, 0x81, 0xb1, 0xa4, 0xdb, 0x24, 0x86, 0x9b, 0x9b, 0x2a, 0x09, 0x38, 0x00, 0x6f, 0x08,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x14, 0x26, 0x4d, 0xd0, 0x80, 0xe4, 0xb3, 0xe7, 0x30, 0x16, 0xbd, 0xf6, 0x57, 0x4b, 0x4c, 0xd8, 0x7e, 0xaa, 0x69, 0x7f, 0x5e,
	/*pkey:*/
}
var txValueSpawnAppTx3 = SpawnAppTx{
	TTL:        0,
	Nonce:      12,
	AppAddress: Address{0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb},
	Amount:     109,
	GasLimit:   1,
	GasPrice:   33,
	CallData:   []byte{},
}
var txCaseSpawnAppTx3Ed = SignedTransaction{
	/*type:*/ 0x04, 0x00, 0x00, 0x00,
	/*sign:*/ 0x3e, 0x60, 0xbd, 0x79, 0x90, 0x49, 0x6c, 0xa2, 0x19, 0x31, 0xe3, 0x86, 0x1b, 0xde, 0xf0, 0x20, 0xd0, 0x6a, 0x11, 0x6b, 0xe8, 0x61, 0x12, 0x3f, 0x03, 0x71, 0xd5, 0x38, 0x12, 0x52, 0x9f, 0xaa, 0xbd, 0x13, 0xeb, 0x24, 0x63, 0x40, 0xf7, 0x79, 0x34, 0x8f, 0x9e, 0x46, 0x2e, 0x2c, 0x56, 0x34, 0x24, 0x19, 0x1b, 0xbf, 0x59, 0x89, 0x64, 0x87, 0x59, 0x99, 0xd8, 0x1e, 0xa5, 0x2a, 0x4c, 0x02,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x14, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
	/*pkey:*/ 0x00, 0x00, 0x00, 0x20, 0x38, 0x08, 0x8e, 0x4c, 0x2a, 0xe8, 0x2f, 0x5c, 0x45, 0xc6, 0x80, 0x8a, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
}
var txCaseSpawnAppTx3EdPlus = SignedTransaction{
	/*type:*/ 0x05, 0x00, 0x00, 0x00,
	/*sign:*/ 0x17, 0x79, 0xc4, 0xc1, 0xd3, 0xde, 0xc8, 0xa1, 0x45, 0xcc, 0x7b, 0x85, 0xe7, 0xe5, 0x7c, 0x11, 0x7e, 0x08, 0x4e, 0x58, 0x7c, 0xed, 0x72, 0x29, 0x01, 0xb5, 0x5a, 0xd3, 0xfd, 0x7f, 0x4f, 0x25, 0x40, 0x5c, 0x6e, 0xf7, 0xc5, 0xae, 0xcf, 0x28, 0xae, 0xdd, 0x1c, 0x35, 0x7c, 0x68, 0x8c, 0x39, 0x2c, 0xf3, 0xdf, 0xf5, 0x36, 0x81, 0x22, 0xa4, 0xd0, 0x29, 0x65, 0x31, 0xb6, 0x32, 0xa3, 0x01,
	/*data:*/ 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x14, 0x61, 0xa6, 0x49, 0x0d, 0x3c, 0x61, 0x2c, 0xe1, 0xda, 0x23, 0x57, 0x14, 0x46, 0x6f, 0xc7, 0x48, 0xfb, 0xc4, 0xcb, 0xbb,
	/*pkey:*/
}
var txCases = []struct {
	Tx     interface{}
	Signed SignedTransaction
}{
	{txValueOldCoinTx1, txCaseOldCoinTx1Ed},
	{txValueOldCoinTx1, txCaseOldCoinTx1EdPlus},
	{txValueOldCoinTx2, txCaseOldCoinTx2Ed},
	{txValueOldCoinTx2, txCaseOldCoinTx2EdPlus},
	{txValueSimpleCoinTx1, txCaseSimpleCoinTx1Ed},
	{txValueSimpleCoinTx1, txCaseSimpleCoinTx1EdPlus},
	{txValueSimpleCoinTx2, txCaseSimpleCoinTx2Ed},
	{txValueSimpleCoinTx2, txCaseSimpleCoinTx2EdPlus},
	{txValueCallAppTx1, txCaseCallAppTx1Ed},
	{txValueCallAppTx1, txCaseCallAppTx1EdPlus},
	{txValueCallAppTx2, txCaseCallAppTx2Ed},
	{txValueCallAppTx2, txCaseCallAppTx2EdPlus},
	{txValueCallAppTx3, txCaseCallAppTx3Ed},
	{txValueCallAppTx3, txCaseCallAppTx3EdPlus},
	{txValueSpawnAppTx1, txCaseSpawnAppTx1Ed},
	{txValueSpawnAppTx1, txCaseSpawnAppTx1EdPlus},
	{txValueSpawnAppTx2, txCaseSpawnAppTx2Ed},
	{txValueSpawnAppTx2, txCaseSpawnAppTx2EdPlus},
	{txValueSpawnAppTx3, txCaseSpawnAppTx3Ed},
	{txValueSpawnAppTx3, txCaseSpawnAppTx3EdPlus},
}

func TestBinaryTransactions(t *testing.T) {
	for _, v := range txCases {
		tx, err := v.Signed.Decode()
		require.NoError(t, err)
		val := reflect.ValueOf(v.Tx)
		b := reflect.New(val.Type())
		ok := tx.Extract(b.Interface())
		require.True(t, ok)
		require.Equal(t, v.Tx, b.Elem().Interface())
	}
}
