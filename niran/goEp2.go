package niran

// 1. Variable และ Basic type (int, float, string, bool)
// 2.Control structure (if, else, switch, loop)
// 3.Data sturce (Arra, map, sturct)
// 4.Function (paramiter, method, return)
// 5. Pointer
// 6. Error Hamding (Throw error)

import (
	"fmt"
)

func VariableTest() {
	// camelCase => firstName
	// snake_case => first_name
	// How to Variable
	//1. var <ชื่อตัวแปร> <ชนิดตัวแปร>
	//2. var <ชื่อตัวแปร> <ชนิดตัวแปร> = <ค่าเริ่มต้น>
	//3. var <ชื่อตัวแปร> = <ค่าเริ่มต้น>
	//4. <ชื่อตัวแปร> := <ต่าเริ่มต้น>

	firstName := "Niran"
	fmt.Printf("First Name: %s\n", firstName)

}

func PointerTest() {
	var a int = 20
	var b int = 30
	const PI float32 = 3.14

	Total := float32(a*b) * PI
	fmt.Printf("Total : %.02f", Total)

}
