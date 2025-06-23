package main

import (
	"time"

	"github.com/dassajib/prohor-api/config"
	"github.com/dassajib/prohor-api/internal/handler"
	"github.com/dassajib/prohor-api/internal/middleware"
	"github.com/dassajib/prohor-api/internal/model"
	"github.com/dassajib/prohor-api/internal/repository"
	"github.com/dassajib/prohor-api/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	db := config.DB

	// to auto migrate user model
	db.AutoMigrate(&model.User{})

	// layered structure with dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(&userService)

	// to initialize gin router with default middleware
	r := gin.Default()

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// handler func call after call api route
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// note routes after checking authorized or not
	noteGroup := r.Group("/notes")
	noteGroup.Use(middleware.AuthMiddleware())
	{
		noteGroup.POST("/")
		noteGroup.GET("/")
		noteGroup.PUT("/")
		noteGroup.DELETE("/:id")
		noteGroup.PUT("restore/:id")
	}

	// serve port on this address
	r.Run(":8080")
}
