package models

type ErrInvalidBookID struct {
	ID string
}

func (e *ErrInvalidBookID) Error() string {
	return "Invalid book ID: " + e.ID
}

type ErrInvalidBookTitle struct {
	Title string
}

func (e *ErrInvalidBookTitle) Error() string {
	return "Invalid book title: " + e.Title
}

type ErrInvalidBookAuthor struct {
	Author string
}

func (e *ErrInvalidBookAuthor) Error() string {
	return "Invalid book author: " + e.Author
}

type ErrInvalidBookDescription struct {
	Description string
}

func (e *ErrInvalidBookDescription) Error() string {
	return "Invalid book description: " + e.Description
}

type ErrInvalidBookGenre struct {
	Genre string
}

func (e *ErrInvalidBookGenre) Error() string {
	return "Invalid book genre: " + e.Genre
}

type ErrInvalidBookTheme struct {
	Theme string
}

func (e *ErrInvalidBookTheme) Error() string {
	return "Invalid book theme: " + e.Theme
}

type ErrInvalidBookLink struct {
	Link string
}

func (e *ErrInvalidBookLink) Error() string {
	return "Invalid book link: " + e.Link
}
