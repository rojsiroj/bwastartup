package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "siroj:1234qwer@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	input := user.LoginInput{
		Email:    "test@test.com",
		Password: "test",
	}
	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.GET("/users/fetch", userHandler.FetchUser)

	router.Run()
}
