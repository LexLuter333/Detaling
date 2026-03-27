import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { authAPI } from '../../api/api';
import { useAuth } from '../../context/AuthContext';
import './Login.css';

const Login = () => {
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await authAPI.login('admin@deteleng.com', password);
      const { token, user } = response.data;
      
      login(user, token);
      navigate('/admin/dashboard');
    } catch (err) {
      setError('Неверный пароль');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      <div className="login-container">
        <div className="login-card">
          <div className="login-logo">
            <h1>OSG Detailing</h1>
            <p>Admin Panel</p>
          </div>

          <form onSubmit={handleSubmit} className="login-form">
            <h2>Вход</h2>
            
            {error && <div className="error-message">{error}</div>}

            <div className="form-group">
              <label htmlFor="password">Пароль</label>
              <input
                type="password"
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
                required
                autoFocus
              />
            </div>

            <button type="submit" className="login-btn" disabled={loading}>
              {loading ? 'Вход...' : 'Войти'}
            </button>

            <div className="login-footer">
              <Link to="/">← На главную</Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
