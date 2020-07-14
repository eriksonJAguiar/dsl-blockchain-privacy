/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

type Record struct {
	LastName           string `json:"lastname"`
	FirstName          string `json:"first_name"`
	NationalProviderId string `json:"national_provider_id"`
	LicenseNumber      string `json:"license_number"`
	MedicalProviderId  string `json:"medical_provider_id"`
	ManagedCarePlansId string `json:"managed_care_plans_id"`
	PrimaryDesignation string `json:"primary_designation"`
	ProviderType       string `json:"provider_type"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Init()", fcn, params)
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createRecord" {
		return cc.createRecord(stub, args)
	} else if function == "queryAllRecords" {
		return cc.queryAllRecords(stub)
	} else if function == "queryTeste" {
		return cc.queryTeste(stub, args)
	} else if function == "createRecordTeste" {
		return cc.createRecordTeste(stub, args)
	} else if function == "executeCommand" {
		return cc.executeCommand(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

//
//{\"selector\":{\"first_name\":\"134\"}}

func (cc *Chaincode) queryTeste(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	resultsIterator, err := stub.GetQueryResult(args[0])
	//	s := "{\"selector\":{\"lastname\":\"134\"}}"
	//	resultsIterator, err := stub.GetQueryResult(s)
	if err != nil {
		return shim.Error("Error doing GetQueryResult1:" + err.Error())
	}

	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error("Error doing GetQueryResult2")
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

func (cc *Chaincode) createRecord(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	var rec = Record{
		LastName:           args[1],
		FirstName:          args[2],
		NationalProviderId: args[3],
		LicenseNumber:      args[4],
		MedicalProviderId:  args[5],
		ManagedCarePlansId: args[6],
		PrimaryDesignation: args[7],
		ProviderType:       args[8],
	}

	recAsBytes, _ := json.Marshal(rec)
	stub.PutState(args[0], recAsBytes)

	return shim.Success(nil)
}

func (cc *Chaincode) createRecordTeste(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	s := "{\"lastname\":\"134\"}"
	//recAsBytes, _ := json.Marshal(args[1])
	///	recAsBytes, _ := json.Marshal(s)
	err := stub.PutState(args[0], []byte(s))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)

}

func (cc *Chaincode) queryAllRecords(stub shim.ChaincodeStubInterface) sc.Response {

	startKey := "key0"
	endKey := "key999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",\n")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllRecords:\n%s\n", buffer.String())
	antlr.NewInputStream("d")
	return shim.Success(buffer.Bytes())
}

func (cc *Chaincode) executeCommand(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	is := antlr.NewInputStream(args[0])

	// Create the Lexer
	lexer := NewTranslatorLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := NewTranslatorParser(stream)

	var listener translatorListener
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// insert command
	if len(listener.insertString) != 0 {
		err := stub.PutState(args[1], []byte(listener.insertString))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	} else {
		resultsIterator, err := stub.GetQueryResult(listener.queryString)

		//	s := "{\"selector\":{\"lastname\":\"134\"}}"
		//	resultsIterator, err := stub.GetQueryResult(s)
		if err != nil {
			return shim.Error("Error doing GetQueryResult1:" + err.Error())
		}
		defer resultsIterator.Close()
		buffer, err := constructQueryResponseFromIterator(resultsIterator)
		if err != nil {
			return shim.Error("Error doing GetQueryResult2")
		}

		fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
		return shim.Success(buffer.Bytes())
	}
}

/*

		"key16",
        "134",
		"FirstName",
		"NationalProviderId",
		"LicenseNumber",
		"MedicalProviderId",
		"ManagedCarePlansId",
		"PrimaryDesignation",
		"ProviderType"


`{lastname:"134",first_name:134,national_provider_id:'134',license_number:'134',medical_provider_id:'134',managed_care_plans_id:'134',primary_designation:'134',provider_type:'134'}`

{adf:'3dd'}

{"selector":{"docType":"marble"}

Error upgrading smart contract: error starting container: error starting container: Failed to generate platform-specific docker build: Error returned from build: 1 "chaincode/input/src/newproj2/chaincode.go:13:2: cannot find package "github.com/eriksonJAguiar/DP-Tool-blockchain/parser" in any of: /chaincode/input/src/newproj2/vendor/github.com/eriksonJAguiar/DP-Tool-blockchain/parser (vendor tree) /opt/go/src/github.com/eriksonJAguiar/DP-Tool-blockchain/parser (from $GOROOT) /chaincode/input/src/github.com/eriksonJAguiar/DP-Tool-blockchain/parser (from $GOPATH) /opt/gopath/src/github.com/eriksonJAguiar/DP-Tool-blockchain/parser "
*/
