package main

// fmt สามารถใช้สำหรับปริ้นข้อความออกมาทางหน้าจอ

import (
	// "fmt"

	// "github.com/google/uuid"
	// "fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/niran/go-example/niran"
)

func main() {
	// id := uuid.New()
	// fmt.Println("Hello World")
	// fmt.Printf("UUID: %s\n", id)

	// How to use package niran
	// niran.SayHelloNiran()
	// niran.VariableTest()
	// niran.PointerTest()
	// niran.ControlStructureTest()
	// niran.LoopTest()

	// niran.DataStructureTest()

	// niran.TestStruct()

	// niran.TestStructinStruct()

	// niran.MyMessage("Niran", 4)
	// Total := niran.Add(10, 50)
	// fmt.Println("Total:", Total)

	// This is call Method
	// student := niran.Student {
	// 	FirstName: "Niran",
	// 	LastName: "Thonphudsa",
	// }

	// fullName := student.FulllName()
	// fmt.Println("Full Name Student is", fullName)

	//interface
	// dog := niran.Dog{Name: "Buddy"}
	// perSon := niran.Person{Name: "Natthakan"}

	// niran.MakeSound(dog)
	// niran.MakeSound(perSon)
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// books1 := niran.Books{
	// 	ID: 1,
	// 	Title: "2003",
	// 	Author: "Niran",
	// }
	// niran.AddBook(books1)
	// books2 := niran.Books{
	// 	ID: 2,
	// 	Title: "2003",
	// 	Author: "Niran TH",
	// }
	// niran.AddBook(books2)
	app.Post("books", niran.AddBook)
	app.Get("/books", niran.GetAllBooks)
	app.Get("/books/:id", niran.GetBook)
	app.Put("/books/:id", niran.UpdateBook)
	app.Delete("/books/:id", niran.DeleteBook)
	app.Listen(":8080")

}
