package services

import (
	"deteleng-backend/internal/models"
	"deteleng-backend/internal/repository"
	"math/rand"
	"time"
)

type ReviewService struct {
	repo repository.Repository
}

func NewReviewService(repo repository.Repository) *ReviewService {
	return &ReviewService{repo: repo}
}

// GetAllReviews returns all reviews
func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.repo.GetAllReviews()
}

// GetApprovedReviews returns approved reviews with limit
func (s *ReviewService) GetApprovedReviews(limit int) ([]models.Review, error) {
	return s.repo.GetApprovedReviews(limit)
}

// CreateReview creates a new review
func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
}

// UpdateReview updates an existing review
func (s *ReviewService) UpdateReview(review *models.Review) error {
	return s.repo.UpdateReview(review)
}

// DeleteReview deletes a review
func (s *ReviewService) DeleteReview(id string) error {
	return s.repo.DeleteReview(id)
}

// GetReviewSources returns all review sources
func (s *ReviewService) GetReviewSources() ([]models.ReviewSource, error) {
	return s.repo.GetReviewSources()
}

// CreateReviewSource creates a new review source
func (s *ReviewService) CreateReviewSource(source *models.ReviewSource) error {
	return s.repo.CreateReviewSource(source)
}

// UpdateReviewSource updates an existing review source
func (s *ReviewService) UpdateReviewSource(source *models.ReviewSource) error {
	return s.repo.UpdateReviewSource(source)
}

// DeleteReviewSource deletes a review source
func (s *ReviewService) DeleteReviewSource(id string) error {
	return s.repo.DeleteReviewSource(id)
}

// ParseReviewsFromSources simulates parsing reviews from sources
// In production, this would use web scraping libraries or APIs
func (s *ReviewService) ParseReviewsFromSources(sourceIDs []string, limit int) ([]models.Review, error) {
	sources, err := s.repo.GetReviewSources()
	if err != nil {
		return nil, err
	}

	// Filter active sources
	activeSources := make([]models.ReviewSource, 0)
	for _, src := range sources {
		if src.IsActive && src.URL != "" {
			if len(sourceIDs) == 0 { // If no specific IDs, use all active
				activeSources = append(activeSources, src)
			} else {
				for _, id := range sourceIDs {
					if src.ID == id {
						activeSources = append(activeSources, src)
						break
					}
				}
			}
		}
	}

	if len(activeSources) == 0 {
		return []models.Review{}, nil
	}

	// Simulate parsed reviews (in production, this would be real scraping)
	parsedReviews := s.simulateParsedReviews(activeSources, limit)

	// Save reviews to repository
	for i := range parsedReviews {
		s.repo.CreateReview(&parsedReviews[i])
	}

	// Update last parse time for sources
	for i := range activeSources {
		activeSources[i].LastParse = time.Now()
		s.repo.UpdateReviewSource(&activeSources[i])
	}

	return parsedReviews, nil
}

// simulateParsedReviews generates sample reviews for demonstration
// In production, replace with actual web scraping logic
func (s *ReviewService) simulateParsedReviews(sources []models.ReviewSource, limit int) []models.Review {
	rand.Seed(time.Now().UnixNano())

	// Sample review data for simulation
	sampleReviews := []struct {
		Author string
		Text   string
		Rating int
	}{
		{"Александр М.", "Отличный сервис! Быстро и качественно сделали полировку кузова. Рекомендую!", 5},
		{"Елена К.", "Заказывала химчистку салона. Результат превзошел ожидания. Очень довольна!", 5},
		{"Дмитрий П.", "Делал тонировку стекол. Всё сделали за 2 часа. Качество на высоте.", 5},
		{"Сергей В.", "Хороший детейлинг центр. Цены адекватные, персонал вежливый.", 4},
		{"Анна С.", "Оклеивала зону риска пленкой. Машина как новая! Спасибо мастерам!", 5},
		{"Максим Р.", "Профессиональный подход. Делал керамику - машина блестит!", 5},
		{"Ольга Н.", "Быстрая и качественная мойка. Пользуюсь услугами уже полгода.", 4},
		{"Игорь Т.", "Удаляли вмятины по технологии PDR. От результата в восторге!", 5},
		{"Наталья Б.", "Заказывала комплексную химчистку. Салон как новый!", 5},
		{"Андрей Л.", "Делал антидождь на стекла. В дождь видимость отличная!", 4},
	}

	reviews := make([]models.Review, 0)
	reviewsPerSource := limit / len(sources)
	if reviewsPerSource < 1 {
		reviewsPerSource = 1
	}

	for _, source := range sources {
		for i := 0; i < reviewsPerSource && len(reviews) < limit; i++ {
			idx := rand.Intn(len(sampleReviews))
			sample := sampleReviews[idx]

			review := models.Review{
				Author:     sample.Author,
				Rating:     sample.Rating,
				Text:       sample.Text,
				Source:     source.Name,
				SourceURL:  source.URL,
				ParsedAt:   time.Now(),
				IsApproved: true,
				IsFavorite: rand.Intn(3) == 0, // 33% chance to be favorite
			}
			reviews = append(reviews, review)
		}
	}

	// Ensure we don't exceed limit
	if len(reviews) > limit {
		reviews = reviews[:limit]
	}

	return reviews
}

// GetReviewStats returns review statistics
func (s *ReviewService) GetReviewStats() (*models.ReviewStats, error) {
	reviews, err := s.repo.GetAllReviews()
	if err != nil {
		return nil, err
	}

	stats := &models.ReviewStats{
		TotalReviews:    int64(len(reviews)),
		SourceBreakdown: make(map[string]int64),
		RatingBreakdown: make(map[int]int64),
		RecentReviews:   make([]models.Review, 0),
	}

	totalRating := 0
	for _, review := range reviews {
		totalRating += review.Rating
		stats.SourceBreakdown[review.Source]++
		stats.RatingBreakdown[review.Rating]++
	}

	if stats.TotalReviews > 0 {
		stats.AverageRating = float64(totalRating) / float64(stats.TotalReviews)
	}

	// Get recent reviews (last 10)
	count := 0
	for i := len(reviews) - 1; i >= 0 && count < 10; i-- {
		if reviews[i].IsApproved {
			stats.RecentReviews = append(stats.RecentReviews, reviews[i])
			count++
		}
	}

	return stats, nil
}
