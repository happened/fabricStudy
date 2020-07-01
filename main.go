package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main(){
	err:=shim.Start(new(SimpleChaincode))
	if err!=nil {
		fmt.Println("Error start Simple chaincode: ",err.Error())
	}
}
