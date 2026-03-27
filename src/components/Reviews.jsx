import React from 'react';
import '../styles/reviews.css';

const reviewsData = [
  {
    id: 1,
    text: 'Парни сделали очень круто! Машина стала выглядеть как новая. Профессиональный подход к своему делу, сделали все быстро.',
    author: 'Алексей Павлов',
    carModel: 'Opel Astra'
  },
  {
    id: 2,
    text: 'Однозначно вернусь еще. Сделали всё быстро и качественно. Убрали все пятна с сидений, салон как новый. Рекомендую!',
    author: 'Владимир Зайцев',
    carModel: 'Mercedes-Benz GLE'
  },
  {
    id: 3,
    text: 'Восстановили фары, заклеили пленкой, думал уже менять фары, но ребята дали фарам вторую жизнь, выглядят как новые.',
    author: 'Кирилл К',
    carModel: 'Lexus IS'
  },
  {
    id: 4,
    text: 'Быстро и качественно очистили салон, убрали все загрязнения. Буду рекомендовать знакомым.',
    author: 'Дмитрий М',
    carModel: 'Skoda Octavia'
  }
];

function Reviews() {
  return (
    <section className="reviews">
      <div className="container">
        <h2>Отзывы клиентов</h2>
        <div className="reviews-grid">
          {reviewsData.map(review => (
            <div key={review.id} className="review-card">
              <p className="review-text">"{review.text}"</p>
              <p className="review-author">{review.author}</p>
              <p className="review-car">{review.carModel}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

export default Reviews;
