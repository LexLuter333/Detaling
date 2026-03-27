import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { bookingsAPI } from '../api/api';
import '../styles/booking.css';

const services = [
  { id: 'svc_1', name: 'Мойка кузова' },
  { id: 'svc_2', name: 'Мойка двигателя' },
  { id: 'svc_3', name: 'Химчистка салона' },
  { id: 'svc_4', name: 'Полировка кузова' },
  { id: 'svc_5', name: 'Керамическое покрытие' },
  { id: 'svc_6', name: 'Оклейка пленкой' },
  { id: 'svc_7', name: 'Тонировка стекол' },
  { id: 'svc_8', name: 'Химчистка крыши' },
  { id: 'svc_9', name: 'Полировка фар' },
  { id: 'svc_10', name: 'Чернение резины' },
];

function BookingForm() {
  const location = useLocation();
  const [formData, setFormData] = useState({
    name: '',
    carBrand: '',
    carModel: '',
    service: '',
    phone: '',
    comment: ''
  });
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    // Читаем параметр service из URL
    const params = new URLSearchParams(location.search);
    const serviceFromUrl = params.get('service');
    
    if (serviceFromUrl) {
      setFormData(prev => ({
        ...prev,
        service: serviceFromUrl
      }));
    }
  }, [location.search]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setSuccess(false);

    try {
      const selectedService = services.find(s => s.name === formData.service);

      await bookingsAPI.createBooking({
        customer_name: formData.name,
        customer_phone: formData.phone,
        car_brand: formData.carBrand,
        car_model: formData.carModel,
        service_id: selectedService?.id || services[0].id,
        comment: formData.comment
      });

      setSuccess(true);
      setFormData({ name: '', carBrand: '', carModel: '', service: '', phone: '', comment: '' });

      setTimeout(() => setSuccess(false), 5000);
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка при отправке заявки');
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="booking" id="booking-section">
      <div className="container">
        <h2>Запишитесь на бесплатную консультацию</h2>

        {success && (
          <div className="success-message">
            ✓ Заявка успешно отправлена! Мы свяжемся с вами в ближайшее время.
          </div>
        )}

        {error && (
          <div className="error-message">
            ✗ {error}
          </div>
        )}

        <form className="booking-form" onSubmit={handleSubmit}>
          <input
            type="text"
            name="name"
            placeholder="Ваше имя"
            value={formData.name}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="carBrand"
            placeholder="Марка автомобиля"
            value={formData.carBrand}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="carModel"
            placeholder="Модель автомобиля"
            value={formData.carModel}
            onChange={handleChange}
          />
          <select
            name="service"
            value={formData.service}
            onChange={handleChange}
            required
          >
            <option value="">Интересующая услуга</option>
            {services.map(s => (
              <option key={s.id} value={s.name}>{s.name}</option>
            ))}
          </select>
          <input
            type="tel"
            name="phone"
            placeholder="+7 (__) ___-__-__"
            value={formData.phone}
            onChange={handleChange}
            required
          />
          <input
            type="text"
            name="comment"
            placeholder="Комментарий (необязательно)"
            value={formData.comment}
            onChange={handleChange}
          />
          <button type="submit" className="submit-btn" disabled={loading}>
            {loading ? 'Отправка...' : 'Записаться'}
          </button>
        </form>
      </div>
    </section>
  );
}

export default BookingForm;
