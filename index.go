package main

// fmt สามารถใช้สำหรับปริ้นข้อความออกมาทางหน้าจอ

import (
	"fmt"
	"github.com/google/uuid"
	)


func main(){
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s", id)
}