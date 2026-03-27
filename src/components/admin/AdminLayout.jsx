import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import './AdminLayout.css';

const AdminLayout = ({ children, onLogout }) => {
  const [sidebarOpen, setSidebarOpen] = useState(true);
  const location = useLocation();

  const menuItems = [
    { path: '/admin/dashboard', label: 'Dashboard', icon: '📊' },
    { path: '/admin/bookings', label: 'Бронирования', icon: '📋' },
    { path: '/admin/services', label: 'Услуги', icon: '🔧' },
    { path: '/admin/reviews', label: 'Отзывы', icon: '💬' },
  ];

  const isActive = (path) => location.pathname === path;

  return (
    <div className="admin-layout">
      <aside className={`sidebar ${sidebarOpen ? 'open' : 'closed'}`}>
        <div className="sidebar-header">
          <Link to="/admin/dashboard" className="sidebar-logo">
            <span className="logo-icon">🏁</span>
            {sidebarOpen && <span className="logo-text">OSG Detailing</span>}
          </Link>
        </div>

        <nav className="sidebar-nav">
          {menuItems.map((item) => (
            <Link
              key={item.path}
              to={item.path}
              className={`nav-item ${isActive(item.path) ? 'active' : ''}`}
            >
              <span className="nav-icon">{item.icon}</span>
              {sidebarOpen && <span className="nav-label">{item.label}</span>}
            </Link>
          ))}
        </nav>

        <div className="sidebar-footer">
          <button onClick={onLogout} className="logout-btn">
            <span className="nav-icon">🚪</span>
            {sidebarOpen && <span>Выйти</span>}
          </button>
        </div>
      </aside>

      <div className="main-content">
        <header className="top-header">
          <button
            className="menu-toggle"
            onClick={() => setSidebarOpen(!sidebarOpen)}
          >
            ☰
          </button>
          <div className="header-right">
            <span className="admin-badge">Admin</span>
          </div>
        </header>

        <main className="content">
          {children}
        </main>
      </div>
    </div>
  );
};

export default AdminLayout;
