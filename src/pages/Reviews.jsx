import { useState, useEffect } from 'react';
import './ReviewsPublic.css';

const ReviewsPublic = () => {
  const [reviews, setReviews] = useState([]);
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState(null);

  useEffect(() => {
    fetchReviews();
  }, []);

  const fetchReviews = async () => {
    try {
      const [reviewsRes, statsRes] = await Promise.all([
        reviewsAPI.getPublic(50),
        reviewsAPI.getStats(),
      ]);
      // Filter only approved and favorite reviews for display
      const allReviews = reviewsRes.data.reviews || [];
      const displayReviews = allReviews.filter(r => r.is_approved);
      
      // Sort: favorite first, then by rating
      displayReviews.sort((a, b) => {
        if (b.is_favorite && !a.is_favorite) return 1;
        if (b.rating !== a.rating) return b.rating - a.rating;
        return 0;
      });

      setReviews(displayReviews);
      setStats(statsRes.data.stats || null);
    } catch (err) {
      console.error('Failed to fetch reviews:', err);
    } finally {
      setLoading(false);
    }
  };

  const getStarRating = (rating) => {
    return '⭐'.repeat(rating);
  };

  const getSourceIcon = (sourceName) => {
    const icons = {
      'Google Maps': '🗺️',
      'Яндекс.Карты': '🗺️',
      '2GIS': '📍',
      'VK': '👥',
    };
    return icons[sourceName] || '📄';
  };

  if (loading) {
    return (
      <>
        <Header />
        <main className="reviews-public-page">
          <div className="loading-container">
            <div className="spinner"></div>
            <p>Загрузка отзывов...</p>
          </div>
        </main>
        <Footer />
      </>
    );
  }

  return (
    <>
      <Header />
      <main className="reviews-public-page">
        <div className="reviews-hero">
          <h1>Отзывы наших клиентов</h1>
          <p>Более {stats?.total_reviews || 0} отзывов со средним рейтингом {stats?.average_rating?.toFixed(1) || '0'} ⭐</p>
        </div>

        {/* Stats */}
        {stats && (
          <div className="reviews-stats-public">
            <div className="stat-item">
              <span className="stat-number">{stats.total_reviews}</span>
              <span className="stat-text">отзывов</span>
            </div>
            <div className="stat-item">
              <span className="stat-number">{stats.average_rating?.toFixed(1) || '0'}</span>
              <span className="stat-text">средний рейтинг</span>
            </div>
            <div className="stat-item">
              <span className="stat-number">{Object.keys(stats.rating_breakdown || {}).filter(r => r >= 4).reduce((acc, r) => acc + (stats.rating_breakdown[r] || 0), 0)}</span>
              <span className="stat-text">положительных</span>
            </div>
          </div>
        )}

        {/* Reviews Grid */}
        <div className="reviews-grid-public">
          {reviews.length === 0 ? (
            <div className="empty-reviews">
              <p>Отзывов пока нет. Будьте первым!</p>
            </div>
          ) : (
            reviews.map((review, index) => (
              <div
                key={review.id}
                className={`review-card-public ${review.is_favorite ? 'favorite' : ''}`}
                style={{ animationDelay: `${index * 0.1}s` }}
              >
                {review.is_favorite && <span className="favorite-badge">⭐ Избранный</span>}
                
                <div className="review-header-public">
                  <div className="review-author-public">
                    <span className="author-name">{review.author}</span>
                  </div>
                  <div className="review-rating-public">
                    {getStarRating(review.rating)}
                  </div>
                </div>

                <p className="review-text-public">{review.text}</p>

                <div className="review-footer-public">
                  <span className="review-source-public">
                    {getSourceIcon(review.source)} {review.source}
                  </span>
                  <span className="review-date-public">
                    {new Date(review.parsed_at).toLocaleDateString('ru-RU', {
                      day: 'numeric',
                      month: 'long',
                      year: 'numeric'
                    })}
                  </span>
                </div>
              </div>
            ))
          )}
        </div>

        {/* CTA Section */}
        <div className="reviews-cta">
          <h2>Оставьте свой отзыв</h2>
          <p>Поделитесь своим опытом работы с нами</p>
          <div className="cta-buttons">
            <a
              href="https://yandex.ru/maps"
              target="_blank"
              rel="noopener noreferrer"
              className="cta-btn"
            >
              🗺️ Оставить на Яндекс.Картах
            </a>
            <a
              href="https://go.2gis.com"
              target="_blank"
              rel="noopener noreferrer"
              className="cta-btn"
            >
              📍 Оставить на 2GIS
            </a>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default ReviewsPublic;
