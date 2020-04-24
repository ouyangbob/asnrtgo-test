package main

import (
	"asnrtgo-test/asnrtapi"
	"github.com/ouyangbob/asnrtgo"
)

func main() {
	asnrtApi:=asnrtapi.GetAsnrtAPI()
	bitstring := asnrtgo.BitString{}
	asnrtApi.Length(bitstring)
	asnrtApi.Bit(bitstring, 0)
	asnrtApi.SetBit(bitstring, 0, true)
	asnrtApi.AllocateBuffer(0, 0)
		//AsnrtInstance.Encode(bitstring)
		//AsnrtInstance.Length(bitstring)

}
