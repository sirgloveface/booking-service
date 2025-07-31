package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirgloveface/booking-service/internal/model"
	services "github.com/sirgloveface/booking-service/internal/service"
	"gorm.io/gorm"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(service *services.BookingService) *BookingHandler {
	return &BookingHandler{service}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var input model.Booking
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.StartTime.After(input.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_time must be before end_time"})
		return
	}

	conflict, err := h.service.HasConflict(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if conflict {
		c.JSON(http.StatusConflict, gin.H{"error": "there is a reservation for this time"})
		return
	}

	if err := h.service.CreateBooking(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": input.ID, "status": input.Status})
}

func (h *BookingHandler) GetBooking(c *gin.Context) {
	id := c.Param("id")

	booking, err := h.service.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (h *BookingHandler) ListBookings(c *gin.Context) {
	result, err := h.service.ListBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la reserva"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
