package main

// fmt สามารถใช้สำหรับปริ้นข้อความออกมาทางหน้าจอ

import (
	"fmt"

	"github.com/google/uuid"

	// import package niran
	"github.com/niran/go-example/niran"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s\n", id)

	// How to use package niran
	// niran.SayHelloNiran()
	// niran.VariableTest()
	// niran.PointerTest()
	// niran.ControlStructureTest()
	// niran.LoopTest()

	niran.DataStructureTest()
	

}

func DataStructureTest() {
	// array
	// slice
	// map 
	// struct

	

}