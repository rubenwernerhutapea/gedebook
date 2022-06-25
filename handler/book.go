package handler

import (
	"gedebook/book"
	"gedebook/helper"
	"gedebook/user"
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

func (h *bookHandler) GetBook(c *gin.Context) {
	// api/v1/books/id
	// handler : mapping id yang ada di url ke struct input => service, call formatter
	// service : inputnya struct input => menangkap id di url, manggil repo
	// repository : get book by id

	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookDetail, err := h.service.GetBookByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Book detail", http.StatusOK, "success", book.FormatBookDetail(bookDetail))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var input book.CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create book", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newBook, err := h.service.CreateBook(input)
	if err != nil {
		response := helper.APIResponse("Failed to create book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create book", http.StatusOK, "success", book.FormatBook(newBook))
	c.JSON(http.StatusOK, response)
}

// tangkap parameter dari user ke input struct
// ambil current user dari jwt/handler
// panggil service, parameternya input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru