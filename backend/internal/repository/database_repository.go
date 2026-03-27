package repository

import (
	"database/sql"
	"deteleng-backend/internal/database"
	"deteleng-backend/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type DatabaseRepository struct {
	db *sql.DB
}

func NewDatabaseRepository() *DatabaseRepository {
	return &DatabaseRepository{
		db: database.DB,
	}
}

// ============ SERVICE OPERATIONS ============

func (r *DatabaseRepository) GetAllServices() ([]models.Service, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price, duration, available, created_at, updated_at
		FROM services ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var s models.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Price, &s.Duration, &s.Available, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func (r *DatabaseRepository) GetService(id string) (*models.Service, error) {
	var s models.Service
	err := r.db.QueryRow(`
		SELECT id, name, description, price, duration, available, created_at, updated_at
		FROM services WHERE id = ?
	`, id).Scan(&s.ID, &s.Name, &s.Description, &s.Price, &s.Duration, &s.Available, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("service not found")
		}
		return nil, err
	}
	return &s, nil
}

func (r *DatabaseRepository) CreateService(service *models.Service) error {
	service.ID = uuid.New().String()
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	_, err := r.db.Exec(`
		INSERT INTO services (id, name, description, price, duration, available, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, service.ID, service.Name, service.Description, service.Price, service.Duration, service.Available, service.CreatedAt, service.UpdatedAt)
	return err
}

func (r *DatabaseRepository) UpdateService(service *models.Service) error {
	service.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE services
		SET name = ?, description = ?, price = ?, duration = ?, available = ?, updated_at = ?
		WHERE id = ?
	`, service.Name, service.Description, service.Price, service.Duration, service.Available, service.UpdatedAt, service.ID)
	return err
}

func (r *DatabaseRepository) DeleteService(id string) error {
	_, err := r.db.Exec("DELETE FROM services WHERE id = ?", id)
	return err
}

// ============ BOOKING OPERATIONS ============

func (r *DatabaseRepository) CreateBooking(booking *models.Booking) error {
	booking.ID = uuid.New().String()
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()
	if booking.Status == "" {
		booking.Status = models.StatusPending
	}

	_, err := r.db.Exec(`
		INSERT INTO bookings (id, customer_name, customer_phone, car_brand, car_model, service_id, service_name, price, status, comment, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, booking.ID, booking.CustomerName, booking.CustomerPhone, booking.CarBrand, booking.CarModel, booking.ServiceID, booking.ServiceName, booking.Price, booking.Status, booking.Comment, booking.CreatedAt, booking.UpdatedAt)
	return err
}

func (r *DatabaseRepository) GetBooking(id string) (*models.Booking, error) {
	var b models.Booking
	err := r.db.QueryRow(`
		SELECT id, customer_name, customer_phone, car_brand, car_model, service_id, service_name, price, status, comment, created_at, updated_at
		FROM bookings WHERE id = ?
	`, id).Scan(&b.ID, &b.CustomerName, &b.CustomerPhone, &b.CarBrand, &b.CarModel, &b.ServiceID, &b.ServiceName, &b.Price, &b.Status, &b.Comment, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &b, nil
}

func (r *DatabaseRepository) GetAllBookings() ([]models.Booking, error) {
	rows, err := r.db.Query(`
		SELECT id, customer_name, customer_phone, car_brand, car_model, service_id, service_name, price, status, comment, created_at, updated_at
		FROM bookings ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		err := rows.Scan(&b.ID, &b.CustomerName, &b.CustomerPhone, &b.CarBrand, &b.CarModel, &b.ServiceID, &b.ServiceName, &b.Price, &b.Status, &b.Comment, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *DatabaseRepository) UpdateBooking(booking *models.Booking) error {
	booking.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE bookings
		SET customer_name = ?, customer_phone = ?, car_brand = ?, car_model = ?, service_id = ?, service_name = ?, price = ?, status = ?, comment = ?, updated_at = ?
		WHERE id = ?
	`, booking.CustomerName, booking.CustomerPhone, booking.CarBrand, booking.CarModel, booking.ServiceID, booking.ServiceName, booking.Price, booking.Status, booking.Comment, booking.UpdatedAt, booking.ID)
	return err
}

func (r *DatabaseRepository) DeleteBooking(id string) error {
	_, err := r.db.Exec("DELETE FROM bookings WHERE id = ?", id)
	return err
}

// DeleteOldCompletedBookings deletes completed bookings older than 7 days
func (r *DatabaseRepository) DeleteOldCompletedBookings() error {
	_, err := r.db.Exec(`
		DELETE FROM bookings 
		WHERE status = 'completed' 
		AND created_at < datetime('now', '-7 days')
	`)
	return err
}

// ============ USER OPERATIONS ============

func (r *DatabaseRepository) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
		SELECT id, email, password, role, created_at
		FROM users WHERE email = ?
	`, email).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *DatabaseRepository) GetUserByID(id string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
		SELECT id, email, password, role, created_at
		FROM users WHERE id = ?
	`, id).Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *DatabaseRepository) CreateUser(user *models.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	_, err := r.db.Exec(`
		INSERT INTO users (id, email, password, role, created_at)
		VALUES (?, ?, ?, ?, ?)
	`, user.ID, user.Email, user.Password, user.Role, user.CreatedAt)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return errors.New("user already exists")
		}
		return err
	}
	return nil
}

// ============ REVIEW OPERATIONS ============

func (r *DatabaseRepository) GetAllReviews() ([]models.Review, error) {
	rows, err := r.db.Query(`
		SELECT id, author, rating, text, source, source_url, parsed_at, is_approved, is_favorite
		FROM reviews ORDER BY parsed_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var rev models.Review
		err := rows.Scan(&rev.ID, &rev.Author, &rev.Rating, &rev.Text, &rev.Source, &rev.SourceURL, &rev.ParsedAt, &rev.IsApproved, &rev.IsFavorite)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, rev)
	}
	return reviews, nil
}

func (r *DatabaseRepository) GetApprovedReviews(limit int) ([]models.Review, error) {
	rows, err := r.db.Query(`
		SELECT id, author, rating, text, source, source_url, parsed_at, is_approved, is_favorite
		FROM reviews 
		WHERE is_approved = 1 
		ORDER BY is_favorite DESC, parsed_at DESC 
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var rev models.Review
		err := rows.Scan(&rev.ID, &rev.Author, &rev.Rating, &rev.Text, &rev.Source, &rev.SourceURL, &rev.ParsedAt, &rev.IsApproved, &rev.IsFavorite)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, rev)
	}
	return reviews, nil
}

func (r *DatabaseRepository) CreateReview(review *models.Review) error {
	review.ID = uuid.New().String()
	if review.ParsedAt.IsZero() {
		review.ParsedAt = time.Now()
	}

	_, err := r.db.Exec(`
		INSERT INTO reviews (id, author, rating, text, source, source_url, parsed_at, is_approved, is_favorite)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, review.ID, review.Author, review.Rating, review.Text, review.Source, review.SourceURL, review.ParsedAt, review.IsApproved, review.IsFavorite)
	return err
}

func (r *DatabaseRepository) UpdateReview(review *models.Review) error {
	_, err := r.db.Exec(`
		UPDATE reviews
		SET author = ?, rating = ?, text = ?, source = ?, source_url = ?, is_approved = ?, is_favorite = ?
		WHERE id = ?
	`, review.Author, review.Rating, review.Text, review.Source, review.SourceURL, review.IsApproved, review.IsFavorite, review.ID)
	return err
}

func (r *DatabaseRepository) DeleteReview(id string) error {
	_, err := r.db.Exec("DELETE FROM reviews WHERE id = ?", id)
	return err
}

// ============ REVIEW SOURCE OPERATIONS ============

func (r *DatabaseRepository) GetReviewSources() ([]models.ReviewSource, error) {
	rows, err := r.db.Query(`
		SELECT id, name, url, is_active, last_parse
		FROM review_sources ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sources []models.ReviewSource
	for rows.Next() {
		var s models.ReviewSource
		err := rows.Scan(&s.ID, &s.Name, &s.URL, &s.IsActive, &s.LastParse)
		if err != nil {
			return nil, err
		}
		sources = append(sources, s)
	}
	return sources, nil
}

func (r *DatabaseRepository) CreateReviewSource(source *models.ReviewSource) error {
	source.ID = uuid.New().String()
	_, err := r.db.Exec(`
		INSERT INTO review_sources (id, name, url, is_active)
		VALUES (?, ?, ?, ?)
	`, source.ID, source.Name, source.URL, source.IsActive)
	return err
}

func (r *DatabaseRepository) UpdateReviewSource(source *models.ReviewSource) error {
	_, err := r.db.Exec(`
		UPDATE review_sources
		SET name = ?, url = ?, is_active = ?, last_parse = ?
		WHERE id = ?
	`, source.Name, source.URL, source.IsActive, source.LastParse, source.ID)
	return err
}

func (r *DatabaseRepository) DeleteReviewSource(id string) error {
	_, err := r.db.Exec("DELETE FROM review_sources WHERE id = ?", id)
	return err
}
