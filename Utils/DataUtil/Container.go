package DataUtil

import (
	"apiGateway/Core/Domain"
	"bytes"
	"encoding/binary"
	"reflect"
	"unsafe"
)

// collection是否包含source，适用map、slice、array
func Contain(source interface{}, collection interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == source {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(source)).IsValid() {
			return true
		}
	}
	return false
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

// 与[]Byte对应的数据结果
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func MessageToBytes(t Domain.Message) []byte {
	// 构造一个与[]byte底层数据结构完全一致的slicemock进行转化
	len := unsafe.Sizeof(t)
	sliceMockTest := SliceMock{
		addr: uintptr(unsafe.Pointer(&t)),
		len:  int(len),
		cap:  int(len),
	}
	return *(*[]byte)(unsafe.Pointer(&sliceMockTest))
}

func BytesToMessage(t []byte) Domain.Message {
	// []byte转换成数据结构，只需取出addr地址即可，然后转换成对应的数据结构类型即可
	return *(*Domain.Message)(unsafe.Pointer(&t[0]))
}

func MapToBytes(t map[string]interface{}) []byte {
	sizeOfMyStruct := int(unsafe.Sizeof(t))

	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(&t))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMap(t []byte) map[string]interface{} {
	// []byte转换成数据结构，只需取出addr地址即可，然后转换成对应的数据结构类型即可
	return *(*map[string]interface{})(unsafe.Pointer(&t[0]))
}
