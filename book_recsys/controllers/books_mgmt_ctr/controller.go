package booksmgmtctr

import (
	"book_recsys/models"
	booksmgmtservice "book_recsys/services/books_mgmt_service"
	"log"

	"github.com/gofiber/fiber/v3"
)

type BookMgmtController struct {
	service *booksmgmtservice.BookService
}

func (c *BookMgmtController) RegisterBook(ctx fiber.Ctx) error {

	book := &models.Book{}

	if err := ctx.Bind().Body(book); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.service.RegisterBook(book.Title, book.Author, book.Description, book.Genre, book.Theme, book.Link); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao registrar o livro",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Livro registrado com sucesso",
	})
}

func (c *BookMgmtController) UpdateBook(ctx fiber.Ctx) error {

	book := &models.Book{}

	if err := ctx.Bind().Body(book); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.service.UpdateBook(book.ID, book.Title, book.Author, book.Description, book.Genre, book.Theme, book.Link, book.Price, book.PriceWithDiscount); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao atualizar o livro." + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Livro atualizado com sucesso",
	})

}

func (c *BookMgmtController) GetBookByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := c.service.GetBookByID(id)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao buscar o livro.",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(book)
}

func (c *BookMgmtController) GetBookByQuery(ctx fiber.Ctx) error {

	genre := ctx.Query("genre")
	theme := ctx.Query("theme")
	author := ctx.Query("author")
	title := ctx.Query("title")

	books, err := c.service.GetBooksByFields(genre, theme, author, title)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao buscar os livros.",
		})
	}

	resp := map[string]interface{}{
		"books": books,
		"count": len(books),
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)

}

func (c *BookMgmtController) GetAllBooks(ctx fiber.Ctx) error {

	books, err := c.service.GetAllBooks()
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao buscar os livros.",
		})
	}

	resp := map[string]interface{}{
		"books": books,
		"count": len(books),
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func NewBookMgmtController(service *booksmgmtservice.BookService) *BookMgmtController {
	return &BookMgmtController{
		service: service,
	}
}
