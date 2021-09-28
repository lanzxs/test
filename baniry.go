package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {
	var i uint16 = 100
	//var t uint32 = 222;
	by := bytes.NewBuffer([]byte{})
	err := binary.Write(by, binary.LittleEndian, i)
	if err != nil {
		panic("error")
	}
	err = binary.Write(by, binary.LittleEndian, []byte("lalalalallalala"))
	if err != nil {
		panic(err)
	}
	var j uint16
	var s []byte = make([]byte, 15)
	fmt.Println(len(by.Bytes()))

	err = binary.Read(by, binary.LittleEndian, &j)
	if err != nil {
		panic("error")
	}
	fmt.Println(len(by.Bytes()))
	fmt.Println(j)
	err = binary.Read(by, binary.LittleEndian, s)
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
	}
	fmt.Println(j)
	fmt.Println(s)
	fmt.Println(string(s))
}
