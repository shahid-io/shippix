package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/shahid-io/shippix/internal/auth/handlers"
	"github.com/shahid-io/shippix/internal/auth/repository"
	"github.com/shahid-io/shippix/internal/auth/service"
)

func main() {
	// dsn := "host=localhost user=shippix password=password dbname=shippix_db port=5432 sslmode=disable"
	dsn := "host=postgres user=shippix password=password dbname=shippix_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo)
	handler := handlers.NewAuthHandler(svc)

	r := gin.Default()
	r.GET("/health", handler.HealthCheck)
	r.POST("/auth/signup", handler.Signup)
	r.POST("/auth/login", handler.Login)

	fmt.Println("Auth service running on :8081")
	r.Run(":8081")
}
