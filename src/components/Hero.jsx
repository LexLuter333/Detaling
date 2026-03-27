import React from 'react';
import '../styles/hero.css';

function Hero() {
  return (
    <section className="hero">
      <video
        className="hero-video"
        autoPlay
        muted
        loop
        playsInline
      >
        <source src="/public/IMG_2145.webm" type="video/webm" />
      </video>

      <div className="hero-overlay" />

      <div className="hero-content">
        <h1>Детейлинг в Екатеринбурге</h1>
        <p>Мы преобразим Ваш автомобиль до неузнаваемости</p>
        <button className="hero-btn">Записаться на консультацию</button>
      </div>
    </section>
  );
}

export default Hero;
