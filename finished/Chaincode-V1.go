package main

import (  
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
const ( 
	UserPrefix	= "USER_" 
 	 
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
	key := UserPrefix + args[0]
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
//Invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	switch function {
	case "Init": if len(args) != 2 {
			return nil, errors.New("Incorrect number of arguments in addUser: expect 2")
		}
		
		t.Init(stub, "init", args)
		  
		return nil, nil

       case "write":
		//patient  := &Patient{}
		var patient Patient
                patient.Username  = args[0]
                patient.Name   = args[1]
                patient.DescriptionOfCurrentAilment = args[2]
                patient.Allergies =args[3]  
                if ((patient.Username =="") && (patient.Name=="") &&(patient.DescriptionOfCurrentAilment =="" )){
		     return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	        }
	         t.write(stub, patient)
		
               return nil, nil
        default:
		errMsg := "No such method in Invoke method: " + function
		fmt.Errorf(errMsg)
		return nil, errors.New(errMsg)
        }
	return nil,nil
}
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
//write
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface,patient Patient) ([]byte, error) {
	var key string
	var err error
	fmt.Println("running write()")

	key = patient.Username //rename for funsies
	//value =( patient.name).append(patient. DescriptionOfCurrentAilment).append(Allergies)
/*var value []byte
	value
	value[0] =(byte)patient.Name
value[1] =(byte)patient.DescriptionOfCurrentAilment
value[2] =(byte)patient.Allergies
	
	err = stub.PutState(key, value) *///write the variable into the chaincode state
	value, err := json.Marshal(&patient)
	
	if err != nil {
		return nil, err
	}
	err = stub.PutState(key, []byte(value))	
	if err != nil {
                        return nil, err
                }
	return nil, nil
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}




