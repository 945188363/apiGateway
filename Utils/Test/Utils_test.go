package Test

import (
	"apiGateway/Core/Domain"
	"apiGateway/Utils"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	msg := Domain.Message{
		Code: 2,
		Msg:  "3w1231",
		Data: nil,
	}
	bytes3 := Utils.MessageToBytes(msg)
	fmt.Println(bytes3)
	message3 := Utils.BytesToMessage(bytes3)
	fmt.Println(message3)
}

func TestMap(t *testing.T) {
	paramMap := make(map[string]interface{})
	paramMap["Size"] = 22
	bytes2 := Utils.IntToBytes(33)
	fmt.Println(bytes2)
	bytes3 := Utils.MapToBytes(paramMap)
	fmt.Println(bytes3)
	message3 := Utils.BytesToMap(bytes3)
	fmt.Println(message3)
}

func TestMap2(t *testing.T) {
	paramMap := make(map[string]interface{})
	paramMap["Size"] = 22
	bytes3 := Utils.MapToBytes(paramMap)
	fmt.Println(bytes3)
	fmt.Println(Utils.BytesToMap(Utils.MapToBytes(paramMap)))
	message3 := Utils.BytesToMap(bytes3)
	fmt.Println(message3)
}
