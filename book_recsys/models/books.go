package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID string `json:"id,omitempty" gorm:"primaryKey"`

	Title             string    `json:"title" required:"true" gorm:"unique"`
	Author            string    `json:"author" required:"true"`
	Description       string    `json:"description" required:"true"`
	Genre             string    `json:"genre" required:"true"`
	Theme             string    `json:"theme" required:"true"`
	Link              string    `json:"link" required:"true" gorm:"unique"`
	Price             float64   `json:"price"`
	PriceWithDiscount float64   `json:"price_with_discount"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewBook(title, author, description, genre, theme, link string) *Book {

	id := uuid.NewString()

	return &Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Description: description,
		Genre:       genre,
		Theme:       theme,
		Link:        link,
	}

}
