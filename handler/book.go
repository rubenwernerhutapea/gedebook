package handler

import (
	"gedebook/book"
	"gedebook/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap paramater di handler
// handler ke service
// service yang menentukan repository mana yang dipanggil
// repository : GetAll, GetByUserID
// db

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

// api/v1/books
func (h *bookHandler) GetBooks(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	books, err := h.service.GetBooks(userID)
	if err != nil {
		response := helper.APIResponse("Error to get books", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("List of books", http.StatusOK, "success", book.FormatBooks(books))
	c.JSON(http.StatusOK, response)
}