package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
import pb "github.com/hyperledger/fabric/protos/peer"

//声明一个结构体
type SimpleChaincode struct {

}

func(t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface)pb.Response{
	//该方法中实现链码初始化或升级时的处理逻辑
	fmt.Println("call init func")
	return pb.Response{}
}

func (t *SimpleChaincode)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	//该方法中实现链码运行中被调用或查询时的处理逻辑
	fmt.Println("call invoke func")
	return pb.Response{}
}


