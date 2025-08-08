package spsp_test

import (
	"fmt"

	"github.com/reiver/go-spsp"
)

func ExampleResponse_DecodeSharedSecret() {

	var response = spsp.Response{
		SharedSecret: "SGVsbG8gd29ybGQhDQpIb3cgYXJlIHlvdSB0b2RheT8=",
	}

	var sharedSecret [32]byte

	err := response.DecodeSharedSecret(&sharedSecret)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

/*
	var sharedSecret [32]byte = [32]byte{
	}
*/

	fmt.Print("BYTES OF SHARED-SECRET:")
	for i:=0; i<len(sharedSecret); i++ {
		var b byte = sharedSecret[i]
		fmt.Printf(" 0x%02X", b)
	}
	fmt.Println()

	// Output:
	// BYTES OF SHARED-SECRET: 0x48 0x65 0x6C 0x6C 0x6F 0x20 0x77 0x6F 0x72 0x6C 0x64 0x21 0x0D 0x0A 0x48 0x6F 0x77 0x20 0x61 0x72 0x65 0x20 0x79 0x6F 0x75 0x20 0x74 0x6F 0x64 0x61 0x79 0x3F
}
