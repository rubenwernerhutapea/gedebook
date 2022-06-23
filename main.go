package main

import (
	"gedebook/auth"
	"gedebook/book"
	"gedebook/handler"
	"gedebook/helper"
	"gedebook/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	bookRepository := book.NewRepository(db)

	userService := user.NewService(userRepository)
	bookService := book.NewService(bookRepository)
	authService := auth.NewService()
	
	userHandler := handler.NewUserHandler(userService, authService)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars",  authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/books", bookHandler.GetBooks)

	router.Run()

	// input dari user
	// handler, mapping input dari user -> struct input
	// service, melakukan mapping dari struct input ke struct User
	// repository
	// db
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}


// ambil nilai header authorization: bearer tokentokentoken
// dari header authorization, kita ambil nilai tokennya saja
// validasi token
// kita ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user