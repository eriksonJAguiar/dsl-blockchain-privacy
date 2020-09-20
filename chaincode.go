/*
 * SPDX-License-Identifier: Apache-2.0
 */

 /*
	alteração no arquivo vendor/gonum.org/v1/gonum/floats/scalar/scalar.go
	//const minNormalFloat64 = 0x1.P-1022
	const minNormalFloat64 = 2.2250738585072014e-308
	
	alteração no arquivo vendor/gonum.org/v1/gonum/lapack/gonum/lapack.go
	//dlamchE = 0x1p-53
	dlamchE = 1.1102230246251565e-16
	//dlamchS = 0x1p-1022
	dlamchS = 2.2250738585072014e-308
	
*/

 package main

 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "math"
	 "strconv"
 
	 "github.com/antlr/antlr4/runtime/Go/antlr"
	 "github.com/eriksonJAguiar/godiffpriv"
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
 
 type DataArray struct {
	 numeric     []float64
	 symbolic    []string
	 maxValue    float64
	 minValue    float64
	 onlyNumbers bool
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
		 buffer.WriteString("\n{\"Key\":")
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
		 // select command
		 return cc.queryCommand(stub, listener)
		 /*resultsIterator, err := stub.GetQueryResult(listener.queryString)
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
		 */
	 }
 }
 
 // INSERT INTO doctype2 (coluna1, coluna2, coluna3) VALUES (\"123\",\"432\",\"1111\")
 
 // INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"123\",\"321\",\"321\")
 // INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"123\",\"654\",\"654\")
 // INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"123\",\"987\",\"987\")
 // INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"3333\",\"3333\",\"3333\")
 // INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"2222\",\"2222\",\"2222\")
 
 // SELECT AVG(coluna2) FROM doctype2 WHERE coluna1=\"123\"
 // SELECT AVG(coluna2), AVG(coluna1), HISTOGRAM(coluna3) FROM doctype3 WHERE coluna1=\"123\"
 // SELECT AVG(coluna2) FROM doctype2 WHERE coluna1="123"
 func (cc *Chaincode) queryCommand(stub shim.ChaincodeStubInterface, listener translatorListener) sc.Response {
	 resultsIterator, err := stub.GetQueryResult(listener.queryString)
	 if err != nil {
		 return shim.Error("Error doing GetQueryResult1:" + err.Error())
	 }
	 var buffer bytes.Buffer
 
	 // create an array with the data for each attribute
	 var data []DataArray = cc.prepareData(listener, resultsIterator)
 
	 // get all the data from the database
	 var alldata []DataArray = cc.getAllData(stub, listener)
 
	 // iterate the attributes applying the diff priv
	 for i, attribute := range listener.attributes {
		 var auxData interface{}
		 if listener.functions[i] == "AVG" {
			 // concatenate the queried data with all the data
			 alldata[i].numeric = append(alldata[i].numeric, data[i].numeric...)
			 auxData = alldata[i].numeric
 
		 } else { // if function == HISTOGRAM
			 // if sturges method must be applied
			 if data[i].onlyNumbers == true {
				 alldata[i].numeric = append(alldata[i].numeric, data[i].numeric...)
				 alldata[i].applySturgesRule()
			 } else {
				 alldata[i].symbolic = append(alldata[i].symbolic, data[i].symbolic...)
			 }
			 auxData = alldata[i].symbolic
		 }
 
		 val := godiffpriv.PrivateDataFactory(auxData)
		 buffer.WriteString("\n" + attribute)
 
		 epsilon := 1.0
		 res, _ := val.ApplyPrivacy(epsilon)
 
		 var response map[string]float64
 
		 err = json.Unmarshal(res, &response)
		 if err != nil {
			 fmt.Println(err.Error())
		 } else {
			 // s := strconv.FormatFloat(response["data"], 'f', 6, 64)
			 buffer.WriteString("\nresponse ")
			 s := fmt.Sprintf("%v", response)
			 buffer.WriteString(s)
 
		 }
	 }
 
	 return shim.Success(buffer.Bytes())
 }
 
 // apply sturges rule on DataArray.numeric insertin the generalized data on DataArray.symbolic
 func (d *DataArray) applySturgesRule() {
	 d.symbolic = make([]string, 0, len(d.numeric))
	 if d.maxValue == d.minValue {
		 panic(1)
	 }
	 N := len(d.numeric)
	 k := 1 + math.Round(math.Log2(float64(N)))
	 delta := d.maxValue - d.minValue
	 a := delta / k
	 for _, value := range d.numeric {
		 classIndex := (value - d.minValue) / a
		 classIndex = math.Floor(classIndex)
		 if classIndex == k { // maximum value doesnt fit to the formula
			 classIndex = k - 1
		 }
		 d.symbolic = append(d.symbolic, strconv.FormatFloat(classIndex, 'f', 0, 64))
	 }
 }
 
 func (cc *Chaincode) getAllData(stub shim.ChaincodeStubInterface, listener translatorListener) []DataArray {
	 s := "{\"selector\":{\"docType\":\"" + listener.docType + "\"}}"
	 resultsIterator, err := stub.GetQueryResult(s)
	 if err != nil {
		 panic(1)
	 }
	 return cc.prepareData(listener, resultsIterator)
 }
 
 // transform the query result into an map
 func (cc *Chaincode) prepareData(listener translatorListener, iterator shim.StateQueryIteratorInterface) []DataArray {
	 defer iterator.Close()
 
	 var data = make([]DataArray, len(listener.attributes))
	 // initialize structs
	 for i, _ := range listener.attributes {
		 data[i].onlyNumbers = true
	 }
	 maxMinFlag := true // flag to define first value as max and min
	 for iterator.HasNext() {
		 queryResponse, err := iterator.Next()
		 if err != nil {
			 panic(1)
		 }
		 var object map[string]string
		 err = json.Unmarshal(queryResponse.Value, &object)
		 if err != nil {
			 panic(1)
		 }
		 // iterate through each attribute queried
		 for i, attribute := range listener.attributes {
			 // check if there is the attribute in the current row
			 if _, found := object[attribute]; !found {
				 panic(1)
			 }
 
			 // if the function is 'average', only numbers are allowed
			 if listener.functions[i] == "AVG" {
				 f, err := strconv.ParseFloat(object[attribute], 64)
				 if err != nil {
					 panic(1)
				 }
				 data[i].numeric = append(data[i].numeric, f)
 
				 // if the function is 'histogram', it can be symbolic data or only numeric data
			 } else if listener.functions[i] == "HISTOGRAM" {
				 data[i].symbolic = append(data[i].symbolic, object[attribute])
 
				 // verify if the data still only has numbers then append the numeric array
				 if data[i].onlyNumbers == true {
					 f, err := strconv.ParseFloat(object[attribute], 64)
					 if err != nil {
						 data[i].onlyNumbers = false // stop appending the numeric array
					 } else {
						 data[i].numeric = append(data[i].numeric, f)
						 // set max and min values
						 if maxMinFlag { // if we are in first row
							 data[i].maxValue = f
							 data[i].minValue = f
							 maxMinFlag = false
						 } else {
							 if data[i].maxValue < f {
								 data[i].maxValue = f
							 } else if data[i].minValue > f {
								 data[i].minValue = f
							 }
						 }
					 }
				 }
			 }
		 }
	 }
	 return data
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