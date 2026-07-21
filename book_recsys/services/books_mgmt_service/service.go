package booksmgmtservice

import (
	"book_recsys/models"
	"book_recsys/repositories/reldb"
)

type BookService struct {
	repoRelDB *reldb.DBClient
}

func (s *BookService) RegisterBook(title, author, description, genre, theme, link string) error {
	book := models.NewBook(title, author, description, genre, theme, link)
	return s.repoRelDB.InsertBook(book)
}

func (s *BookService) UpdateBook(id, title, author, description, genre, theme, link string, price, priceWithDiscount float64) error {

	if id == "" {
		return &models.ErrInvalidBookID{ID: id}
	}

	if title == "" {
		return &models.ErrInvalidBookTitle{Title: title}
	}

	if author == "" {
		return &models.ErrInvalidBookAuthor{Author: author}
	}

	if description == "" {
		return &models.ErrInvalidBookDescription{Description: description}
	}

	if genre == "" {
		return &models.ErrInvalidBookGenre{Genre: genre}
	}

	if theme == "" {
		return &models.ErrInvalidBookTheme{Theme: theme}
	}

	if link == "" {
		return &models.ErrInvalidBookLink{Link: link}
	}

	book, err := s.repoRelDB.GetBookByID(id)
	if err != nil {
		return err
	}

	book.Title = title
	book.Author = author
	book.Description = description
	book.Genre = genre
	book.Theme = theme
	book.Link = link
	book.Price = price
	book.PriceWithDiscount = priceWithDiscount
	return s.repoRelDB.UpdateBook(book)
}

func (s *BookService) DeleteBook(id string) error {
	return s.repoRelDB.DeleteBook(id)
}

func (s *BookService) GetBookByID(id string) (*models.Book, error) {
	return s.repoRelDB.GetBookByID(id)
}

func (s *BookService) GetAllBooks() ([]*models.Book, error) {
	return s.repoRelDB.GetAllBooks()
}

func (s *BookService) GetBooksByGenre(genre string) ([]*models.Book, error) {
	return s.repoRelDB.GetBooksByGenre(genre)
}

func (s *BookService) GetBooksByTheme(theme string) ([]*models.Book, error) {
	return s.repoRelDB.GetBooksByTheme(theme)
}

func (s *BookService) GetBooksByAuthor(author string) ([]*models.Book, error) {
	return s.repoRelDB.GetBooksByAuthor(author)
}

func (s *BookService) GetBooksByFields(genre, theme, author, title string) ([]*models.Book, error) {
	return s.repoRelDB.GetBooksByFields(genre, theme, author, title)
}

func NewBookService(repoRelDB *reldb.DBClient) *BookService {
	return &BookService{
		repoRelDB: repoRelDB,
	}
}
