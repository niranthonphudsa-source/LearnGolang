package niran

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

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

func GetAllBooks(c *fiber.Ctx) error {
	books = append(books, Books{ID: 1, Title: "2003", Author: "Niran"})
	books = append(books, Books{ID: 2, Title: "2005", Author: "Niran TH"})
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

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set clamis
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encode token and send it as respons
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   t,
	})

}

func CheckMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println(claims)
	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}
	// start := time.Now()
	// fmt.Printf("URL = %s Method = %s Time = %s\n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}
