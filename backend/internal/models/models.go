package models

import "time"

// BookingStatus represents the status of a booking
type BookingStatus string

const (
	StatusPending   BookingStatus = "pending"
	StatusConfirmed BookingStatus = "confirmed"
	StatusCompleted BookingStatus = "completed"
	StatusCancelled BookingStatus = "cancelled"
)

// Booking represents a customer booking
type Booking struct {
	ID            string        `json:"id"`
	CustomerName  string        `json:"customer_name"`
	CustomerPhone string        `json:"customer_phone"`
	CarBrand      string        `json:"car_brand"`
	CarModel      string        `json:"car_model"`
	ServiceID     string        `json:"service_id"`
	ServiceName   string        `json:"service_name"`
	Price         float64       `json:"price"`
	Status        BookingStatus `json:"status"`
	Comment       string        `json:"comment,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

// Service represents a detailing service
type Service struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Duration    int       `json:"duration_minutes"`
	Available   bool      `json:"available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// User represents an admin user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// DashboardStats represents admin dashboard statistics
type DashboardStats struct {
	TotalBookings   int64                  `json:"total_bookings"`
	PendingBookings int64                  `json:"pending_bookings"`
	ConfirmedBookings int64                `json:"confirmed_bookings"`
	CompletedBookings int64                `json:"completed_bookings"`
	TotalRevenue    float64                `json:"total_revenue"`
	RecentBookings  []Booking              `json:"recent_bookings"`
	StatusBreakdown map[BookingStatus]int64 `json:"status_breakdown"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents JWT token response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateBookingRequest represents a booking creation request
type CreateBookingRequest struct {
	CustomerName  string  `json:"customer_name" binding:"required"`
	CustomerPhone string  `json:"customer_phone" binding:"required"`
	CarBrand      string  `json:"car_brand" binding:"required"`
	CarModel      string  `json:"car_model"`
	ServiceID     string  `json:"service_id" binding:"required"`
	Comment       string  `json:"comment"`
}

// UpdateBookingStatusRequest represents a status update request
type UpdateBookingStatusRequest struct {
	Status BookingStatus `json:"status" binding:"required"`
}

// Review represents a customer review
type Review struct {
	ID         string    `json:"id"`
	Author     string    `json:"author"`
	Rating     int       `json:"rating"`
	Text       string    `json:"text"`
	Source     string    `json:"source"` // google, yandex, 2gis, vk, etc.
	SourceURL  string    `json:"source_url"`
	ParsedAt   time.Time `json:"parsed_at"`
	IsApproved bool      `json:"is_approved"`
	IsFavorite bool      `json:"is_favorite"`
}

// ReviewSource represents a source for parsing reviews
type ReviewSource struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"` // Google Maps, Yandex Maps, 2GIS, VK
	URL       string    `json:"url"`
	IsActive  bool      `json:"is_active"`
	LastParse time.Time `json:"last_parse,omitempty"`
}

// AddReviewSourceRequest represents a request to add a review source
type AddReviewSourceRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

// ParseReviewsRequest represents a request to parse reviews
type ParseReviewsRequest struct {
	SourceIDs []string `json:"source_ids"`
	Limit     int      `json:"limit"`
}

// ReviewStats represents review statistics
type ReviewStats struct {
	TotalReviews   int64              `json:"total_reviews"`
	AverageRating  float64            `json:"average_rating"`
	SourceBreakdown map[string]int64  `json:"source_breakdown"`
	RatingBreakdown map[int]int64     `json:"rating_breakdown"`
	RecentReviews   []Review          `json:"recent_reviews"`
}
