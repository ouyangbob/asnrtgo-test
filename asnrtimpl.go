package main

import (
	. "github.com/ouyangbob/asnrtgo"
)

func Length(self BitString) uint {
	println("BitString.Length")
	return uint(len(self.Bytes)<<3) - uint(self.UnusedBits)
}

func Bit(self BitString, index uint) bool {
	println("BitString.Bit")
	return false
}

func SetBit(self BitString, index uint, b bool) {
	println("BitString.SetBit")
}

func AllocateBuffer(length int, encodingRules byte) (Buffer, error) {
	println("AllocateBuffer")
	return nil, nil
}

func Encode(b Buffer, v Dynamic, t AsnType) error {
	println("Encode")
	return nil
}

func Decode(b Buffer, v Dynamic, t AsnType) error {
	println("Decode")
	return nil
}

func DecodeMetadata(metadata []byte) map[int]AsnType {
	println("DecodeMetadata")
	types := make(map[int]AsnType)
	return types
}
