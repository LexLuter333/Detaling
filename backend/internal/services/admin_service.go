package services

import (
	"deteleng-backend/internal/models"
	"deteleng-backend/internal/repository"
)

type AdminService struct {
	repo repository.Repository
}

func NewAdminService(repo repository.Repository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) GetDashboardStats() (*models.DashboardStats, error) {
	bookings, err := s.repo.GetAllBookings()
	if err != nil {
		return nil, err
	}

	stats := &models.DashboardStats{
		RecentBookings:  []models.Booking{},
		StatusBreakdown: make(map[models.BookingStatus]int64),
	}

	for _, booking := range bookings {
		stats.TotalBookings++
		stats.StatusBreakdown[booking.Status]++

		if booking.Status == models.StatusCompleted {
			stats.TotalRevenue += booking.Price
		}

		switch booking.Status {
		case models.StatusPending:
			stats.PendingBookings++
		case models.StatusConfirmed:
			stats.ConfirmedBookings++
		case models.StatusCompleted:
			stats.CompletedBookings++
		}
	}

	// Get recent bookings (last 10)
	recentCount := 0
	for i := len(bookings) - 1; i >= 0 && recentCount < 10; i-- {
		stats.RecentBookings = append(stats.RecentBookings, bookings[i])
		recentCount++
	}

	return stats, nil
}

func (s *AdminService) GetAllBookings() ([]models.Booking, error) {
	return s.repo.GetAllBookings()
}

func (s *AdminService) UpdateBookingStatus(id string, status models.BookingStatus) (*models.Booking, error) {
	booking, err := s.repo.GetBooking(id)
	if err != nil {
		return nil, err
	}

	booking.Status = status
	if err := s.repo.UpdateBooking(booking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *AdminService) DeleteBooking(id string) error {
	return s.repo.DeleteBooking(id)
}

func (s *AdminService) GetBooking(id string) (*models.Booking, error) {
	return s.repo.GetBooking(id)
}
