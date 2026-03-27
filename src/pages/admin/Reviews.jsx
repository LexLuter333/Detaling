import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import AdminLayout from '../../components/admin/AdminLayout';
import { useAuth } from '../../context/AuthContext';
import { reviewsAPI, reviewSourcesAPI } from '../../api/api';
import './Reviews.css';

const Reviews = () => {
  const [reviews, setReviews] = useState([]);
  const [sources, setSources] = useState([]);
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);
  const [parseLoading, setParseLoading] = useState(false);
  const [parseLimit, setParseLimit] = useState(50);
  const [selectedSources, setSelectedSources] = useState([]);
  const [showAddSource, setShowAddSource] = useState(false);
  const [newSource, setNewSource] = useState({ name: '', url: '' });
  const { logout } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const [reviewsRes, sourcesRes, statsRes] = await Promise.all([
        reviewsAPI.getAll(),
        reviewSourcesAPI.getAll(),
        reviewsAPI.getStats(),
      ]);
      setReviews(reviewsRes.data.reviews || []);
      setSources(sourcesRes.data.sources || []);
      setStats(statsRes.data.stats || null);
    } catch (err) {
      console.error('Failed to fetch data:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/admin/login');
  };

  const handleParseReviews = async () => {
    setParseLoading(true);
    try {
      await reviewsAPI.parse(selectedSources, parseLimit);
      await fetchData();
      alert(`Парсинг завершен! Найдено отзывов: ${parseLimit}`);
    } catch (err) {
      alert('Ошибка при парсинге: ' + (err.response?.data?.error || err.message));
    } finally {
      setParseLoading(false);
    }
  };

  const handleToggleSource = (sourceId) => {
    setSelectedSources(prev =>
      prev.includes(sourceId)
        ? prev.filter(id => id !== sourceId)
        : [...prev, sourceId]
    );
  };

  const handleAddSource = async () => {
    if (!newSource.name || !newSource.url) {
      alert('Заполните все поля');
      return;
    }
    try {
      await reviewSourcesAPI.create(newSource);
      setNewSource({ name: '', url: '' });
      setShowAddSource(false);
      await fetchData();
    } catch (err) {
      alert('Ошибка при добавлении источника');
    }
  };

  const handleToggleSourceActive = async (source) => {
    try {
      await reviewSourcesAPI.update(source.id, {
        ...source,
        is_active: !source.is_active,
      });
      await fetchData();
    } catch (err) {
      alert('Ошибка при обновлении источника');
    }
  };

  const handleDeleteSource = async (id) => {
    if (!confirm('Удалить источник отзывов?')) return;
    try {
      await reviewSourcesAPI.delete(id);
      await fetchData();
    } catch (err) {
      alert('Ошибка при удалении источника');
    }
  };

  const handleDeleteReview = async (id) => {
    if (!confirm('Удалить этот отзыв?')) return;
    try {
      await reviewsAPI.delete(id);
      await fetchData();
    } catch (err) {
      alert('Ошибка при удалении отзыва');
    }
  };

  const handleToggleFavorite = async (review) => {
    try {
      await reviewsAPI.update(review.id, {
        ...review,
        is_favorite: !review.is_favorite,
      });
      await fetchData();
    } catch (err) {
      alert('Ошибка при обновлении отзыва');
    }
  };

  const handleToggleApproved = async (review) => {
    try {
      await reviewsAPI.update(review.id, {
        ...review,
        is_approved: !review.is_approved,
      });
      await fetchData();
    } catch (err) {
      alert('Ошибка при обновлении отзыва');
    }
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

  const getStarRating = (rating) => {
    return '⭐'.repeat(rating);
  };

  if (loading) {
    return (
      <AdminLayout onLogout={handleLogout}>
        <div className="loading-container">
          <div className="spinner"></div>
          <p>Загрузка отзывов...</p>
        </div>
      </AdminLayout>
    );
  }

  return (
    <AdminLayout onLogout={handleLogout}>
      <div className="reviews-page">
        <div className="page-header">
          <h1>💬 Управление отзывами</h1>
        </div>

        {/* Stats */}
        {stats && (
          <div className="reviews-stats">
            <div className="stat-card">
              <span className="stat-value">{stats.total_reviews}</span>
              <span className="stat-label">Всего отзывов</span>
            </div>
            <div className="stat-card">
              <span className="stat-value">{stats.average_rating?.toFixed(1) || '0'}</span>
              <span className="stat-label">Средний рейтинг</span>
            </div>
            <div className="stat-card">
              <span className="stat-value">{Object.keys(stats.source_breakdown || {}).length}</span>
              <span className="stat-label">Источников</span>
            </div>
          </div>
        )}

        {/* Parse Section */}
        <div className="parse-section">
          <h2>🔄 Парсинг отзывов</h2>
          <p className="parse-description">
            Выберите источники и нажмите "Начать парсинг" для загрузки отзывов (макс. 50)
          </p>

          <div className="sources-list">
            {sources.map(source => (
              <div key={source.id} className="source-item">
                <label className="source-checkbox">
                  <input
                    type="checkbox"
                    checked={selectedSources.includes(source.id)}
                    onChange={() => handleToggleSource(source.id)}
                    disabled={!source.is_active || !source.url}
                  />
                  <span className="source-name">
                    {getSourceIcon(source.name)} {source.name}
                  </span>
                  {source.url && (
                    <a href={source.url} target="_blank" rel="noopener noreferrer" className="source-link">
                      🔗
                    </a>
                  )}
                  <span className={`source-status ${source.is_active && source.url ? 'active' : 'inactive'}`}>
                    {source.is_active && source.url ? 'Активен' : 'Не активен'}
                  </span>
                </label>
              </div>
            ))}
          </div>

          <div className="parse-controls">
            <div className="limit-control">
              <label>Лимит отзывов:</label>
              <input
                type="number"
                min="1"
                max="50"
                value={parseLimit}
                onChange={(e) => setParseLimit(Math.min(50, Math.max(1, parseInt(e.target.value) || 1)))}
              />
            </div>
            <button
              className="parse-btn"
              onClick={handleParseReviews}
              disabled={parseLoading || selectedSources.length === 0}
            >
              {parseLoading ? '⏳ Парсинг...' : `🚀 Начать парсинг (${selectedSources.length} выбрано)`}
            </button>
          </div>
        </div>

        {/* Add Source Section */}
        <div className="add-source-section">
          <button className="add-source-btn" onClick={() => setShowAddSource(!showAddSource)}>
            {showAddSource ? '✕ Отмена' : '+ Добавить источник'}
          </button>

          {showAddSource && (
            <div className="add-source-form">
              <input
                type="text"
                placeholder="Название (например: Google Maps)"
                value={newSource.name}
                onChange={(e) => setNewSource({ ...newSource, name: e.target.value })}
              />
              <input
                type="url"
                placeholder="URL страницы с отзывами"
                value={newSource.url}
                onChange={(e) => setNewSource({ ...newSource, url: e.target.value })}
              />
              <button className="save-source-btn" onClick={handleAddSource}>
                Сохранить
              </button>
            </div>
          )}
        </div>

        {/* Sources Management */}
        <div className="sources-management">
          <h2>📋 Источники отзывов</h2>
          <div className="sources-table">
            {sources.map(source => (
              <div key={source.id} className="source-row">
                <span>{getSourceIcon(source.name)} {source.name}</span>
                <span className="source-url-cell">{source.url || '—'}</span>
                <button
                  className={`toggle-btn ${source.is_active ? 'active' : ''}`}
                  onClick={() => handleToggleSourceActive(source)}
                >
                  {source.is_active ? '✓ Вкл' : '✗ Выкл'}
                </button>
                <button className="delete-btn" onClick={() => handleDeleteSource(source.id)}>
                  🗑️
                </button>
              </div>
            ))}
          </div>
        </div>

        {/* Reviews List */}
        <div className="reviews-section">
          <h2>📝 Все отзывы ({reviews.length})</h2>
          <div className="reviews-list">
            {reviews.length === 0 ? (
              <div className="empty-state">
                <p>Отзывов пока нет. Нажмите "Начать парсинг" для загрузки.</p>
              </div>
            ) : (
              reviews.map(review => (
                <div key={review.id} className={`review-card ${review.is_favorite ? 'favorite' : ''}`}>
                  <div className="review-header">
                    <div className="review-author">
                      <span className="author-name">{review.author}</span>
                      <span className="review-rating">{getStarRating(review.rating)}</span>
                    </div>
                    <div className="review-meta">
                      <span className="review-source">{getSourceIcon(review.source)} {review.source}</span>
                      <span className="review-date">
                        {new Date(review.parsed_at).toLocaleDateString('ru-RU')}
                      </span>
                    </div>
                  </div>
                  <p className="review-text">{review.text}</p>
                  <div className="review-actions">
                    <label className="action-checkbox">
                      <input
                        type="checkbox"
                        checked={review.is_approved}
                        onChange={() => handleToggleApproved(review)}
                      />
                      Одобрено
                    </label>
                    <label className="action-checkbox">
                      <input
                        type="checkbox"
                        checked={review.is_favorite}
                        onChange={() => handleToggleFavorite(review)}
                      />
                      Избранное ⭐
                    </label>
                    <button className="delete-review-btn" onClick={() => handleDeleteReview(review.id)}>
                      🗑️ Удалить
                    </button>
                  </div>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </AdminLayout>
  );
};

export default Reviews;
