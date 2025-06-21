package main

import (
	"time"

	"github.com/dassajib/prohor-api/config"
	"github.com/dassajib/prohor-api/internal/handler"
	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/repository"
	"github.com/dassajib/prohor-api/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	db := config.DB

	// auto migrate User model
	db.AutoMigrate(&model.User{})

	// Initialize the application components in a layered structure using dependency injection.
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// initialize the router with default middleware
	r := gin.Default()

	// cors config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// listen and serve on localhost:8080
	r.Run(":8080")
}
