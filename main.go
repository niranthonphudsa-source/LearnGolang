package main

// fmt สามารถใช้สำหรับปริ้นข้อความออกมาทางหน้าจอ

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/niran/go-example/niran"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Load Env Error")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", niran.LogIn)
	
	app.Use(niran.CheckMiddleware)
	app.Get("/testHtml", niran.TestHtml)
	app.Post("books", niran.AddBook)
	app.Get("/books", niran.GetAllBooks)
	app.Get("/books/:id", niran.GetBook)
	app.Put("/books/:id", niran.UpdateBook)
	app.Delete("/books/:id", niran.DeleteBook)
	app.Post("/upload", niran.UploadFile)
	app.Get("/config", niran.GetEnv)
	
	app.Listen(":8080")

}
