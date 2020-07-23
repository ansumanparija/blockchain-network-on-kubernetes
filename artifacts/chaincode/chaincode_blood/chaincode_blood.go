package main
import (
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
)
// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}
// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
  // params := stub.GetStringArgs()
  // fmt.Println("Init()", params)
  // if len(params) != 3 {
  //  return shim.Error("Incorrect arguments. Expecting a key and a value")
  // }
  // asBytes, err := json.Marshal(params[2])
  // if err != nil {
  //  log.Panic(err)
  // }
  // err = stub.PutState(params[1], asBytes)
  // handleError(err)
  return shim.Success(nil)
}
// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
  fcn, params := stub.GetFunctionAndParameters()
  fmt.Println("Invoke()", fcn, params)
  if fcn == "createDonor" {
    return cc.createDonor(stub, params)
  }
  return shim.Success(nil)
}
func (cc *Chaincode) createDonor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
  if len(args) != 2 {
    return shim.Error("Incorrect arguments. Expecting a key and a value")
  }
  err := stub.PutState(args[0], []byte(args[1]))
  if err != nil {
    return shim.Error(err.Error())
  }
  return shim.Success([]byte(args[1]))
}
func (cc *Chaincode) getDonor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect arguments. Expecting a key only")
  }
  asBytes, err := stub.GetState(args[0])
  if err != nil {
    return shim.Error(err.Error())
  }
  return shim.Success(asBytes)
}
func main() {
  err := shim.Start(new(Chaincode))
  if err != nil {
    panic(err)
  }
}
