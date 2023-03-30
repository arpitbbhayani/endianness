package main

import (
	"fmt"
)

func ToByteArrayFromUInt32LE(x uint32) []byte {
	var barray []byte = make([]byte, 4)
	for i := 0; i < 4; i++ {
		b := x & 0x000000FF
		barray[i] = byte(b)
		x >>= 8
	}
	return barray
}

func ToUInt32FromByteArrayLE(barray []byte) uint32 {
	var number uint32 = 0
	for i := 0; i < 4; i++ {
		number = number | (uint32(barray[i]) << (8 * i))
	}
	return number
}

func ToUInt32FromByteArrayBE(barray []byte) uint32 {
	var number uint32 = 0
	for i := 0; i < 4; i++ {
		number = (number << 8) | (uint32(barray[i]))
	}
	return number
}

func main() {
	var number uint32 = 0x10203040

	var barrayLE []byte = ToByteArrayFromUInt32LE(number)
	fmt.Println("number", number, "byte array as LE", barrayLE)
	fmt.Printf("LE interpretation: 0x%x, int value %d\n", ToUInt32FromByteArrayLE(barrayLE), ToUInt32FromByteArrayLE(barrayLE))
	fmt.Printf("BE interpretation: 0x%x, int value %d\n", ToUInt32FromByteArrayBE(barrayLE), ToUInt32FromByteArrayBE(barrayLE))
}
