package asnrtapi

import (
	."github.com/ouyangbob/asnrtgo"
	"github.com/tcard/pluginunmarshal"
	"sync"
)

var asnrtApi *AsnrtAPI

var once sync.Once

func GetAsnrtAPI() *AsnrtAPI {
	once.Do(func(){
		asnrtApi = &AsnrtAPI{}
		err := pluginunmarshal.Open("asnrt.so", asnrtApi)
		if err != nil {
			panic(err)
		}
	})
	return asnrtApi
}

