package handlers

import (
	"deteleng-backend/internal/models"
	"deteleng-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(service *services.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

// CreateBooking handles POST /api/bookings
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req models.CreateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking, err := h.service.CreateBooking(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"booking": booking,
	})
}

// GetAllBookings handles GET /api/bookings
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

// GetBooking handles GET /api/bookings/:id
func (h *BookingHandler) GetBooking(c *gin.Context) {
	id := c.Param("id")

	booking, err := h.service.GetBooking(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}
