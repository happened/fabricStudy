package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)
import pb "github.com/hyperledger/fabric/protos/peer"

/**
链码有4个功能

初始化的时候设定AB两个账户的余额

Invoke提供
- invoke 两个账户之间的转账 A B money
- query 查询某个账户的余额  A
- delete 删除某个账户的信息 A
*/
//声明一个结构体
type SimpleChaincode struct {
}

func(t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface)pb.Response{
	//该方法中实现链码初始化或升级时的处理逻辑
	fmt.Println("call init func")

	_,args:=stub.GetFunctionAndParameters()

	var A,B string
	var Aval,Bval int

	var err error

	if len(args)!=4 {
		return shim.Error("Incorrect number of arguments.Expecting 4")
	}

	A=args[0]
	Aval,err=strconv.Atoi(args[1])
	if err!=nil{
		return shim.Error("error when get A balance: "+err.Error())
	}

	B=args[2]
	Bval,err=strconv.Atoi(args[3])
	if err!=nil{
		return shim.Error("error when get B balance :"+err.Error())
	}

	err=stub.PutState(A,[]byte(strconv.Itoa(Aval)))
	if err!=nil{
		return shim.Error("error when put A state: "+err.Error())
	}
	err=stub.PutState(B,[]byte(strconv.Itoa(Bval)))
	if err!=nil{
		return shim.Error("error when put B state: "+err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	//该方法中实现链码运行中被调用或查询时的处理逻辑
	fmt.Println("call invoke func")

	function,args:=stub.GetFunctionAndParameters()
	if function=="invoke" {
		return t.invoke(stub,args)
	}else if function=="query" {
		return t.query(stub,args)
	}else if function=="delete" {
		return t.delete(stub,args)
	}
	return pb.Response{}
}

func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var src,dst string
	var value int
	var err error

	if len(args)!=3 {
		return shim.Error("The usage is wrong. invoke account1 account2 moneyNum")
	}

	src=args[0]
	dst=args[1]
	value,err=strconv.Atoi(args[2])

	AvalByte,err:=stub.GetState(src)
	if err!=nil {
		return shim.Error("failed to get src balance : "+err.Error())
	}
	BvalByte,err:=stub.GetState(dst)
	if err!=nil {
		return shim.Error("failed to get src balance : "+err.Error())
	}

	Aval,_:=strconv.Atoi(string(AvalByte))
	Bval,_:=strconv.Atoi(string(BvalByte))

	if Aval<value{
		return shim.Error("src account does not have enough money")
	}

	Aval=Aval-value
	Bval=Bval+value

	err=stub.PutState(src,[]byte(strconv.Itoa(Aval)))
	if err!=nil{
		return shim.Error("failed to put src balance : "+err.Error())
	}

	err=stub.PutState(dst,[]byte(strconv.Itoa(Bval)))
	if err!=nil{
		return shim.Error("failed to put dst balance: "+err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args)!=1 {
		return shim.Error("The usage is wrong. invoke accountName")
	}

	account:=args[0]

	valueByte,err:=stub.GetState(account)
	if err!=nil{
		return shim.Error("faile to get account balance: "+err.Error())
	}

	fmt.Printf("Account: %s Balacne: %s \n",account,string(valueByte))

	return shim.Success(valueByte)

}

func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args)!=1 {
		return shim.Error("The usage is wrong. invoke accountName")
	}

	account:=args[0]
	err:=stub.DelState(account)
	if err!=nil{
		return shim.Error("failed to delete account : "+err.Error())
	}

	return shim.Success(nil)
}


