package reldb

import (
	"book_recsys/models"
	"book_recsys/settings"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	Con *gorm.DB
}

func (c *DBClient) InsertBook(book *models.Book) error {
	return c.Con.Create(book).Error
}

func (c *DBClient) GetBookByID(id string) (*models.Book, error) {
	var book models.Book
	if err := c.Con.First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (c *DBClient) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	if err := c.Con.Order("title ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *DBClient) GetBooksByGenre(genre string) ([]*models.Book, error) {
	var books []*models.Book
	if err := c.Con.Where("genre = ?", genre).Order("title ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *DBClient) GetBooksByTheme(theme string) ([]*models.Book, error) {
	var books []*models.Book
	if err := c.Con.Where("theme = ?", theme).Order("title ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *DBClient) GetBooksByAuthor(author string) ([]*models.Book, error) {
	var books []*models.Book
	if err := c.Con.Where("author = ?", author).Order("title ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *DBClient) GetBooksByFields(genre, theme, author, title string) ([]*models.Book, error) {
	var books []*models.Book
	query := c.Con

	if genre != "" {
		query = query.Where("genre = ?", genre)
	}
	if theme != "" {
		query = query.Where("theme = ?", theme)
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}
	if title != "" {
		query = query.Where("title = ?", title)
	}
	if err := query.Order("title ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (c *DBClient) UpdateBook(book *models.Book) error {
	return c.Con.Save(book).Error
}

func (c *DBClient) DeleteBook(id string) error {
	return c.Con.Delete(&models.Book{}, "id = ?", id).Error
}

func NewDBClient(cfg *settings.Config) (*DBClient, error) {

	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, cfg.DBUsername, cfg.DBPassword, cfg.DBURL, cfg.DBPort, cfg.DBName)

	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client := &DBClient{
		Con: con,
	}

	return client, nil
}
