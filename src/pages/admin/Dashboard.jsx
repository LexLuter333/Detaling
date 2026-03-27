import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { adminAPI } from '../../api/api';
import { useAuth } from '../../context/AuthContext';
import AdminLayout from '../../components/admin/AdminLayout';
import './Dashboard.css';

const Dashboard = () => {
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const { logout } = useAuth();

  useEffect(() => {
    fetchDashboard();
  }, []);

  const fetchDashboard = async () => {
    try {
      const response = await adminAPI.getDashboard();
      setStats(response.data.stats);
    } catch (err) {
      setError('Failed to load dashboard data');
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    logout();
    window.location.href = '/admin/login';
  };

  if (loading) {
    return (
      <AdminLayout onLogout={handleLogout}>
        <div className="loading-container">
          <div className="spinner"></div>
          <p>Loading dashboard...</p>
        </div>
      </AdminLayout>
    );
  }

  if (error) {
    return (
      <AdminLayout onLogout={handleLogout}>
        <div className="error-container">
          <p className="error-message">{error}</p>
          <button onClick={fetchDashboard} className="retry-btn">Retry</button>
        </div>
      </AdminLayout>
    );
  }

  return (
    <AdminLayout onLogout={handleLogout}>
      <div className="dashboard">
        <div className="dashboard-header">
          <h1>Dashboard</h1>
          <p>Обзор статистики и управления бронированиями</p>
        </div>

        <div className="stats-grid">
          <div className="stat-card total">
            <div className="stat-icon">📋</div>
            <div className="stat-info">
              <span className="stat-value">{stats?.total_bookings || 0}</span>
              <span className="stat-label">Всего бронирований</span>
            </div>
          </div>

          <div className="stat-card pending">
            <div className="stat-icon">⏳</div>
            <div className="stat-info">
              <span className="stat-value">{stats?.pending_bookings || 0}</span>
              <span className="stat-label">Ожидают</span>
            </div>
          </div>

          <div className="stat-card confirmed">
            <div className="stat-icon">✅</div>
            <div className="stat-info">
              <span className="stat-value">{stats?.confirmed_bookings || 0}</span>
              <span className="stat-label">Подтверждено</span>
            </div>
          </div>

          <div className="stat-card completed">
            <div className="stat-icon">🎉</div>
            <div className="stat-info">
              <span className="stat-value">{stats?.completed_bookings || 0}</span>
              <span className="stat-label">Завершено</span>
            </div>
          </div>

          <div className="stat-card revenue">
            <div className="stat-icon">💰</div>
            <div className="stat-info">
              <span className="stat-value">{stats?.total_revenue?.toLocaleString('ru-RU') || 0} ₽</span>
              <span className="stat-label">Выручка</span>
            </div>
          </div>
        </div>

        <div className="dashboard-actions">
          <Link to="/admin/bookings" className="action-btn primary">
            📋 Все бронирования
          </Link>
          <Link to="/admin/services" className="action-btn secondary">
            🔧 Услуги
          </Link>
        </div>

        {stats?.recent_bookings?.length > 0 && (
          <div className="recent-bookings">
            <h2>Последние бронирования</h2>
            <div className="bookings-table-container">
              <table className="bookings-table">
                <thead>
                  <tr>
                    <th>Клиент</th>
                    <th>Автомобиль</th>
                    <th>Услуга</th>
                    <th>Цена</th>
                    <th>Статус</th>
                  </tr>
                </thead>
                <tbody>
                  {stats.recent_bookings.slice(0, 5).map((booking) => (
                    <tr key={booking.id}>
                      <td>{booking.customer_name}</td>
                      <td>{booking.car_brand} {booking.car_model}</td>
                      <td>{booking.service_name}</td>
                      <td>{booking.price.toLocaleString('ru-RU')} ₽</td>
                      <td>
                        <span className={`status-badge status-${booking.status}`}>
                          {getStatusText(booking.status)}
                        </span>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )}
      </div>
    </AdminLayout>
  );
};

function getStatusText(status) {
  const statusMap = {
    pending: 'Ожидает',
    confirmed: 'Подтверждено',
    completed: 'Завершено',
    cancelled: 'Отменено',
  };
  return statusMap[status] || status;
}

export default Dashboard;
