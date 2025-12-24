package niran

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

// func AddBook(b Books) []Books {
// 	// fmt.Println(books)
// 	books = append(books, Books{ID: b.ID, Title: b.Title, Author: b.Author})
// 	return books
// }

func GetAllBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {

	book := new(Books)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString(" Not Found!!! ")
}

func UpdateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	bookUpdate := new(Books)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Update Book Not Found!!")
}

func DeleteBook(c *fiber.Ctx) error {

	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			// ... คือการกระจาย slice
			// [1,2,3,4,5]
			// [1, 2] + [4, 5]
			books = append(books[:i], books[i+1:]...)
			return c.Status(fiber.StatusNoContent).SendString("Delete Success!")
		}

	}
	return c.Status(fiber.StatusNotFound).SendString("Delete Book Not Found!!")
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./upload/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File Upload Success!")

}

func TestHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello World!",
		"Name":  "Niran",
	})
}

func GetEnv(c *fiber.Ctx) error {

	// if value, exits := os.LookupEnv("SECRET"); exits {
	// 	return c.JSON(fiber.Map{
	// 		"SECRET": value,
	// 	})
	// }

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = Users{
	Email:    "user@xample.com",
	Password: "password1234",
}

func LogIn(c *fiber.Ctx) error {
	user := new(Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"message": "Login Success",
	})
}

func CheckMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}
