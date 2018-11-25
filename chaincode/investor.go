package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Investor struct{

}

type IDetail struct{
	name string
	score string
	asset IAsset
}

type IAsset struct{
	typ string
	name string
	value int
}

func (i *Investor) Init(stub shim.ChaincodeStubInterface) peer.Response{
	return shim.Success(nil);
}

func (i *Investor) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fn, args := stub.GetFunctionAndParameters()
	var err error
	if fn == "new"{
		_,err = nInvestor(stub,args);
	}else if fn == "createasset" {
		_,err = createAsset(stub,args);
	}
	var er error
	if err != nil {
		return shim.Error(er.Error())
	}
	return shim.Success(nil)
}

func nInvestor(stub shim.ChaincodeStubInterface, args []string) (string, error){
	stub.PutState(args[0],[]byte(args[1]))
	stub.PutState(args[2],[]byte(args[3]))
	return "done",nil
}

func createAsset(stub shim.ChaincodeStubInterface, args []string) (string, error){
	stub.PutState(args[0],[]byte(args[1]))
	stub.PutState(args[2],[]byte(args[3]))
	stub.PutState(args[4],[]byte(args[5]))
	return "done",nil
}

func queryAssets(stub shim.ChaincodeStubInterface, args []string) (string, error){
	
}



func main(){
	shim.Start(new(Investor))
}


