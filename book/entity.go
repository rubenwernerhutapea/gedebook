package book

import "gedebook/user"

type Book struct {
	ID               int
	UserID           int
	Name             string
	FileImage        string
	ShortDescription string
	Description      string
	Quantity         int
	Slug             string
	CreatedAt        int
	UpdatedAt        int
	User             user.User
}