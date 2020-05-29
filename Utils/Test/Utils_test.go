package Test

import (
	"apiGateway/Core/Domain"
	"apiGateway/Utils/DataUtil"
	"apiGateway/Utils/RedisUtil"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	msg := Domain.Message{
		Code: 2,
		Msg:  "3w1231",
		Data: nil,
	}
	bytes3 := DataUtil.MessageToBytes(msg)
	fmt.Println(bytes3)
	message3 := DataUtil.BytesToMessage(bytes3)
	fmt.Println(message3)
}

func TestMap(t *testing.T) {
	paramMap := make(map[string]interface{})
	paramMap["Size"] = 22
	bytes2 := DataUtil.IntToBytes(33)
	fmt.Println(bytes2)
	bytes3 := DataUtil.MapToBytes(paramMap)
	fmt.Println(bytes3)
	message3 := DataUtil.BytesToMap(bytes3)
	fmt.Println(message3)
}

func TestMap2(t *testing.T) {
	paramMap := make(map[string]interface{})
	paramMap["Size"] = 22
	bytes3 := DataUtil.MapToBytes(paramMap)
	fmt.Println(bytes3)
	fmt.Println(DataUtil.BytesToMap(DataUtil.MapToBytes(paramMap)))
	message3 := DataUtil.BytesToMap(bytes3)
	fmt.Println(message3)
}

func TestRedis(t *testing.T) {
	get, _ := RedisUtil.Get("Name")

	fmt.Println(get)
}
