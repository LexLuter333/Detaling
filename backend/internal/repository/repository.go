package repository

import (
	"deteleng-backend/internal/models"
)

type Repository interface {
	// Service operations
	GetAllServices() ([]models.Service, error)
	GetService(id string) (*models.Service, error)
	CreateService(service *models.Service) error
	UpdateService(service *models.Service) error
	DeleteService(id string) error

	// Booking operations
	CreateBooking(booking *models.Booking) error
	GetBooking(id string) (*models.Booking, error)
	GetAllBookings() ([]models.Booking, error)
	UpdateBooking(booking *models.Booking) error
	DeleteBooking(id string) error
	DeleteOldCompletedBookings() error

	// User operations
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error

	// Review operations
	GetAllReviews() ([]models.Review, error)
	GetApprovedReviews(limit int) ([]models.Review, error)
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(id string) error

	// Review Source operations
	GetReviewSources() ([]models.ReviewSource, error)
	CreateReviewSource(source *models.ReviewSource) error
	UpdateReviewSource(source *models.ReviewSource) error
	DeleteReviewSource(id string) error
}
