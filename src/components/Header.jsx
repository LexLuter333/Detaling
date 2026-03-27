import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/header.css';

function Header() {
  return (
    <header className="header">
      <div className="container">
        <Link to="/" className="logo">
          <img src="/Icon_logo.webp" alt="OSG Detailing" className="logo-icon" />
        </Link>
        <nav className="nav">
          <Link to="/" className="nav-link">Главная</Link>
          <Link to="/services" className="nav-link">Услуги</Link>
          <Link to="/reviews" className="nav-link">Отзывы</Link>
          <Link to="/contacts" className="nav-link">Контакты</Link>
        </nav>
      </div>
    </header>
  );
}

export default Header;
