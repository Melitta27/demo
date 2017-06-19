package main

import ( 
             "encoding/json"
            "strings"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
//Patient Struct
  type Patient struct {
                Username   string  `json:"Username"`
                Name     string  `json:"Name"`
                DescriptionOfCurrentAilment     string  `json:"DescriptionOfCurrentAilment"`
                Allergies      string  `json:"Allergies"`
            }

type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
//Init
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var key string
	var err error
key = UserPrefix + args[0]
	if len(args) != 2{
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err = stub.PutState(key, []byte(args[1]))
fmt.Println("store user:%s sucessfully", args[0])
fmt.Printf("your username is:%s",key)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
//write
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface,Patient patient) ([]byte, error) {
	var key string
	var err error
	fmt.Println("running write()")

	key = patient.username //rename for funsies
	//value =( patient.name).append(patient. DescriptionOfCurrentAilment).append(Allergies)
Value []byte
value[0] =patient.name
value[1] =patient. DescriptionOfCurrentAilment
value[2] =patient. Allergies
	err = stub.PutState(key, value) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//Invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var username string
	var value string
	var err error
	switch function {
	case "Init": if len(args) != 2 {
			return nil, errors.New("Incorrect number of arguments in addUser: expect 2")
		}
		username = args[0]
		value = args[1]
		err = t.Init (stub, "init", args)
		if err != nil {
			fmt.Println("addUser error: ", err)
		}
		return nil, err

case "write":
		patient  := &Patient{}
patient.Username  = args[0]
 patient.Name   = args[1]
   patient.DescriptionOfCurrentAilment = args[2]
    patient.Allergies =args[3]  
if ((patient.Username="")&& (patient.name="")&&(patient. DescriptionOfCurrentAilment ="" )){
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
err = t.write (stub, patient)
		if err != nil {
			fmt.Println("error: ", err)
		}
return nil, err
}




