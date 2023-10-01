package lib

import (
	"encoding/json"
	"fmt"
	"github.com/abolfazlalz/goeventer/internal/interpreter/miscs"
)

type JsonDecode struct {
}

func (JsonDecode) Name() string {
	return "json_decode"
}

func (j JsonDecode) Call(args ...interface{}) interface{} {
	if len(args) != 1 {
		panic(fmt.Sprintf("invalid arguments for %s method", j.Name()))
	}

	jsonStr := args[0].(*miscs.Variable).String()
	var unmarshalValue map[string]interface{}
	data := make(miscs.StructType)
	if err := json.Unmarshal([]byte(jsonStr), &unmarshalValue); err != nil {
		panic(fmt.Sprintf("invalid json data: %v", err))
	}

	for key, value := range unmarshalValue {
		data[key] = miscs.NewVariable(value)
	}

	return miscs.NewReturnStatement(miscs.NewVariable(data))
}

type JsonEncode struct {
}

func (JsonEncode) Name() string {
	return "json_encode"
}

func (j JsonEncode) Call(args ...interface{}) interface{} {
	if len(args) != 1 {
		panic(fmt.Sprintf("invalid arguments for %s method", j.Call()))
	}

	data := args[0].(*miscs.Variable).StructValue()
	dataDict := make(map[string]interface{})
	for key, value := range data {
		dataDict[key] = value.Value
	}

	result, err := json.Marshal(dataDict)
	if err != nil {
		panic(err)
	}

	return miscs.NewReturnStatement(miscs.NewVariable(string(result)))
}
