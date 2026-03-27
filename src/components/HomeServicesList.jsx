import React from 'react';
import '../styles/services.css';

const homeServicesData = [
  {
    id: 1,
    name: 'Трехфазная мойка кузова',
    price: 'Цена от 1500',
    description: 'Способ бесконтактной очистки автомобиля, при котором химия наносится в три этапа.'
  },
  {
    id: 2,
    name: 'Химчистка салона',
    price: 'Цена от 6500',
    description: 'Комплексная глубокая очистка внутренних поверхностей автомобиля.'
  },
  {
    id: 3,
    name: 'Полировка кузова',
    price: 'Цена от 10000',
    description: 'Процесс восстановления лакокрасочного покрытия и придания ему зеркального блеска.'
  },
  {
    id: 4,
    name: 'Озонация',
    price: 'Цена от 1500',
    description: 'Обработка салона озоном для удаления запахов.'
  },
  {
    id: 5,
    name: 'Антидождь',
    price: 'Цена от 1000',
    description: 'Нанесение гидрофобного покрытия на стекла автомобиля.'
  },
  {
    id: 6,
    name: 'Оклейка зон риска',
    price: 'Цена от 45000',
    description: 'Нанесение защитной пленки на наиболее уязвимые участки кузова.'
  }
];

function HomeServicesList() {
  return (
    <section className="services-list home-specific">
      <div className="container">
        <h2>Наши услуги</h2>
        <div className="services-grid">
          {homeServicesData.map(service => (
            <div key={service.id} className="service-card">
              <div className="service-header">
                <h3>{service.name}</h3>
                {service.price && <span className="service-price">{service.price}</span>}
              </div>
              {service.description && <p>{service.description}</p>}
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

export default HomeServicesList;
