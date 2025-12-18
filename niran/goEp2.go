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

func ControlStructureTest() {
	// if else
	//switch case
	// loop for

	var grade string
	var score int = 80
	if score >= 80{
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60{
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("Grade: %s\n", grade)

	switch score {
	case 80: grade = "A"
	case 70: grade = "B"
	case 60: grade = "C"
	case 50: grade = "D"
	default: grade = "F"	
	}
	fmt.Printf("Grade: %s\n", grade)

}
