# Deteleng - Detailing Studio Platform

A full-stack web application for a car detailing studio with a public website and admin panel.

## 🚀 Features

### Public Website
- Modern, responsive design with dark theme
- Services catalog with pricing
- Online booking form
- Contact information with map integration
- Reviews and testimonials

### Admin Panel
- Secure authentication with JWT
- Dashboard with statistics
- Booking management (view, update status, delete)
- Services management
- Real-time data updates

## 📁 Project Structure

```
Deteleng/
├── backend/                 # Go Gin backend
│   ├── cmd/
│   │   └── api/
│   │       └── main.go     # Application entry point
│   ├── internal/
│   │   ├── config/         # Configuration management
│   │   ├── handlers/       # HTTP handlers
│   │   ├── middleware/     # Auth & CORS middleware
│   │   ├── models/         # Data models
│   │   ├── repository/     # Data access layer
│   │   └── services/       # Business logic
│   ├── go.mod
│   └── Dockerfile
├── frontend/                # React + Vite frontend
│   ├── src/
│   │   ├── api/            # API client
│   │   ├── components/     # Reusable components
│   │   ├── context/        # React context (Auth)
│   │   ├── pages/          # Page components
│   │   │   ├── admin/      # Admin panel pages
│   │   │   ├── Home.jsx
│   │   │   ├── Services.jsx
│   │   │   └── Contacts.jsx
│   │   └── styles/         # Global styles
│   ├── package.json
│   └── Dockerfile
└── docker-compose.yml
```

## 🛠️ Tech Stack

### Backend
- **Go 1.22** with Gin Framework
- **JWT** for authentication
- **In-memory repository** (easily replaceable with database)
- **CORS** enabled for frontend communication

### Frontend
- **React 18** with hooks
- **React Router** for navigation
- **Axios** for API calls
- **Vite** for fast development
- **CSS3** with custom variables

## 🚦 Getting Started

### Prerequisites
- Go 1.22+
- Node.js 18+
- Docker & Docker Compose (optional)

### Running with Docker (Recommended)

```bash
# Start both services
docker-compose up -d

# Backend: http://localhost:8080
# Frontend: http://localhost:5173
# Admin Panel: http://localhost:5173/admin/login
```

### Running Locally

#### Backend
```bash
cd backend

# Install dependencies
go mod download

# Run server
go run ./cmd/api

# Server runs on http://localhost:8080
```

#### Frontend
```bash
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev

# App runs on http://localhost:5173
```

## 🔐 Admin Panel Access

**Default Credentials:**
- Email: `admin@deteleng.com`
- Password: `admin123`

**Admin Routes:**
- `/admin/login` - Login page
- `/admin/dashboard` - Dashboard with statistics
- `/admin/bookings` - Manage bookings
- `/admin/services` - Manage services

## 📡 API Endpoints

### Public Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/ping` | Health check |
| POST | `/api/bookings` | Create booking |
| GET | `/api/bookings` | Get all bookings |
| GET | `/api/bookings/:id` | Get booking by ID |
| POST | `/api/auth/login` | Login |
| POST | `/api/auth/register` | Register new user |

### Protected Endpoints (Admin)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/admin/dashboard` | Get dashboard stats |
| GET | `/api/admin/bookings` | Get all bookings |
| PUT | `/api/admin/bookings/:id/status` | Update booking status |
| DELETE | `/api/admin/bookings/:id` | Delete booking |
| GET | `/api/admin/stats` | Get statistics |

## 📦 Available Services

The system comes pre-configured with 10 detailing services:
1. Мойка кузова - 800 ₽
2. Мойка двигателя - 1,500 ₽
3. Химчистка салона - 8,000 ₽
4. Полировка кузова - 15,000 ₽
5. Керамическое покрытие - 25,000 ₽
6. Оклейка пленкой - 12,000 ₽
7. Тонировка стекол - 5,000 ₽
8. Химчистка крыши - 3,000 ₽
9. Полировка фар - 2,500 ₽
10. Чернение резины - 500 ₽

## 🔒 Security

- JWT token-based authentication
- Password hashing with bcrypt
- CORS protection
- Protected admin routes
- Role-based access control

## 🎨 Design Features

- Dark theme with orange accents (#f4593b)
- Responsive design (mobile-first)
- Smooth animations and transitions
- Glassmorphism effects
- Video background in hero section

## 📝 Environment Variables

### Backend
```bash
SERVER_PORT=8080
JWT_SECRET=deteleng-secret-key-2026
ADMIN_EMAIL=admin@deteleng.com
ADMIN_PASSWORD=admin123
```

## 🔄 Future Enhancements

- PostgreSQL database integration
- Email notifications
- Image upload for services
- Booking calendar view
- Analytics dashboard
- Multi-language support

## 📄 License

This project is created for educational purposes.
