package mysql

import "fmt"

func BytesTransfer2String(byte []byte) string{
	return string(byte[:])
}

func StringTransfer2Bytes(string string) []byte{
	return  []byte(string)
}

func BytesTransfer2StringExample() {
	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)
}