package main

import (
	"gedebook/handler"
	"gedebook/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// input dari user
	// handler, mapping input dari user -> struct input
	// service, melakukan mapping dari struct input ke struct User
	// repository
	// db
}
