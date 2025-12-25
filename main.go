package main

// fmt สามารถใช้สำหรับปริ้นข้อความออกมาทางหน้าจอ

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	_ "github.com/niran/go-example/docs"
	"github.com/niran/go-example/niran"

	"github.com/gofiber/jwt/v2"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Books
// @Router /books [get]
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Load Env Error")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		})

	app.Get("/swagger/*", swagger.HandlerDefault) //default
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	t := app.Post("/login", niran.LogIn)
	fmt.Print(t)
	
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(niran.CheckMiddleware)

	app.Get("/testHtml", niran.TestHtml)
	app.Post("/books", niran.AddBook)
	app.Get("/books", niran.GetAllBooks)
	app.Get("/books/:id", niran.GetBook)
	app.Put("/books/:id", niran.UpdateBook)
	app.Delete("/books/:id", niran.DeleteBook)
	app.Post("/upload", niran.UploadFile)
	app.Get("/config", niran.GetEnv)

	app.Listen(":8080")

}
