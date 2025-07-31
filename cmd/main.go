// cmd/main.go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sirgloveface/booking-service/cmd/docs"

	//"github.com/sirgloveface/booking-service/internal/auth"
	"github.com/sirgloveface/booking-service/internal/config"
	"github.com/sirgloveface/booking-service/internal/handler"
	services "github.com/sirgloveface/booking-service/internal/service"

	// "github.com/sirgloveface/booking-service/internal/handler"
	// "github.com/sirgloveface/booking-service/internal/middleware"
	repositories "github.com/sirgloveface/booking-service/internal/repository"
	// "github.com/sirgloveface/booking-service/internal/service"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task API
// @version         1.0
// @description     API para gestionar tareas usando Go, Gin y MongoDB.
// @host            localhost:8080
// @BasePath        /

func main() {
	r := gin.Default()
	db := config.ConnectPostgres()
	repo := repositories.NewBookingRepository(db)
	service := services.NewBookingService(repo)
	handler := handler.NewBookingHandler(service)

	booking := r.Group("/bookings")
	{
		booking.POST("/", handler.CreateBooking)
		booking.GET("/:id", handler.GetBooking)
		booking.GET("/", handler.ListBookings)
		booking.DELETE("/", handler.DeleteBooking)
	}

	r.Run()
}
