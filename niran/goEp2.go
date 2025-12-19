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
	fmt.Printf("Total : %.02f\n", Total)

}

func ControlStructureTest() {
	// if else
	//switch case
	// loop for

	var grade string
	var score int = 80
	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("Grade: %s\n", grade)

	var Month int = 4
	switch Month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid Month")
	}

}

func LoopTest() {
	// nomal for loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// do while loop

	j := 1
	for {
		fmt.Println(j)
		j++
		if j >= 10 {
			break
		}
	}


	// while loop
	k := 1
	for k < 5{
		fmt.Println(k)
		k++
	}


}


func DataStructureTest() {
	// array  var a [5]int
	// slice var a[]int
	// map 
	// struct

	//Array and Slice

	// var MyArray [5]int = [5]int{10, 20, 30, 40, 50}
	// fmt.Println(MyArray)

	// for i := 0; i < len(MyArray); i++ {
	// 	fmt.Printf("Number %d: %d\n", i, MyArray[i])
	// 	MyArray[0] = 100
	// }
	// fmt.Println(MyArray[0])

	// var mySlice []string = []string{"niran", "kritthapas", "sirawit"}

	// for num := 0; num < len(mySlice); num++{
	// 	fmt.Printf("Name: %s\n", mySlice[num])
	// }
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// fmt.Print("Success !\n")
	// subSlice := mySlice[0:2]
	// fmt.Println(subSlice)
	// fmt.Println(len(subSlice))
	// fmt.Println(cap(subSlice))

	// mySlice = append(mySlice, "Natthakan")
	// fmt.Println(mySlice)

	// //Convert Array to Slice 
	// myArray := [4]string{"niran", "kritthapas", "sirawit", "Natthakan"}

	// arrToSlice := myArray[:]
	
	// arrToSlice = append(arrToSlice, "Teekhayu")
	// fmt.Println(arrToSlice)



	// Map
	// 1 key have 1 value
	// <name> := make(map[type key]type value)
	// myMap := make(map[string]int)
	// myMap["Niran"] = 22
	// myMap["Natthakan"] = 23
	// myMap["Krittapas"] = 23

	// fmt.Println(myMap)
	// fmt.Println("Niran age", myMap["Niran"])
	// fmt.Println("Natthakan age", myMap["Natthakan"])


	// // delete
	// delete(myMap, "Krittapas")
	// for key, value :=  range myMap {
	// 	fmt.Printf("Name: %s Age: %d\n", key, value)
	// }

	// // check if key exit
	// val, ok := myMap["Krittapas"]
	// if ok {
	// 	fmt.Println("Pear value:", val)
	// } else {
	// 	fmt.Println("Not Found")
	// }

	


	



	// type Dog interface {
	// 	Bark() string
	// }

	// Pointer
	// var a *int
}

// Struct
// type Person struct {
// 	Name string
// 	Age int 
// 	Weight float32
// 	Height float32
// 	Address Address
// }

// func TestStruct(){

// 	var Person1 Person
// 	Person1.Name = "Niran"
// 	Person1.Age = 22
// 	Person1.Weight = 70.5
// 	Person1.Height = 175.5
// 	fmt.Printf("Data is Person1: %+v\n", Person1)

// 	var person  []Person
// 	person = append(person, Person{
// 		Name: "Niran", 
// 		Age: 22, 
// 		Weight: 70.5,
// 		Height: 175.5,
// 	})
// 	person = append(person, Person{
// 		Name: "Natthakan", 
// 		Age: 23, 
// 		Weight: 71.5, 
// 		Height: 115.5,
// 	})


// 	fmt.Printf("Data is Person: %+v\n", person)
// 	fmt.Printf("Name: %s\n", person[0].Name)


// 	for i := 0; i < len(person); i++ {
// 		fmt.Printf("Person: %+v\n", person[i])
// 	}

// 	for j := 0; j < len(person); j++{
// 		fmt.Println("Name: ", person[j].Name)
// 		fmt.Println("Age: ", person[j].Age)
// 		fmt.Println("Weight: ", person[j].Weight)
// 		fmt.Println("Height: ", person[j].Height)
// 	}


// }


// Struct in Struct
// type Address struct {
// 	street string
// 	city string
// 	country string
// }
// func TestStructinStruct(){
// 	var person  []Person
// 	person = append(person, Person{
// 		Name: "Niran", 
// 		Age: 22, 
// 		Weight: 70.5,
// 		Height: 175.5,
// 		Address: Address{
// 			street: "333 Mian St",
// 			city: "Bankok",
// 			country: "Thailand",
// 		},
// 	})

// 	for j := 0; j < len(person); j++{
// 		fmt.Println("Name: ", person[j].Name)
// 		fmt.Println("Age: ", person[j].Age)
// 		fmt.Println("Weight: ", person[j].Weight)
// 		fmt.Println("Height: ", person[j].Height)
// 		fmt.Println("Address: ", person[j].Address.street, 
// 						person[j].Address.city, 
// 						person[j].Address.country)
// 	}
// }


func MyMessage(name string, num int) {
	fmt.Println("I just executed!")
	for i := 0; i <= num; i++{
		fmt.Printf("My name is %s\n", name)
	}

}

func Add(a int, b int) int {
	return a + b
}