package book

type GetBookDetailInput struct {
	ID int `uri:"id" binding:"required"`
}