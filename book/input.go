package book

import "gedebook/user"

type GetBookDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateBookInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"Description" binding:"required"`
	FileImage        string `json:"file_image" binding:"required"`
	Quantity         int    `json:"quantity" binding:"required"`
	User             user.User
}