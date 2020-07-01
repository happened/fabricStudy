package main

import "github.com/hyperledger/fabric/core/chaincode/shim"
import pb "github.com/hyperledger/fabric/protos/peer"

type SimpleChaincode struct {

}

func(t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface)pb.Response{
	return pb.Response{}
}

func (t *SimpleChaincode)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	return pb.Response{}
}


