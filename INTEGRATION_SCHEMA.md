# 🔗 Схема интеграции Frontend и Backend

## Архитектура приложения

```
┌─────────────────────────────────────────────────────────────────┐
│                         БРАУЗЕР                                 │
│                     http://localhost:5173                       │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                    REACT FRONTEND                               │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐           │
│  │  Public     │  │   Admin      │  │   Auth       │           │
│  │  Pages      │  │   Panel      │  │   Context    │           │
│  │             │  │              │  │              │           │
│  │ - Home      │  │ - Login      │  │ - JWT Token  │           │
│  │ - Services  │  │ - Dashboard  │  │ - User State │           │
│  │ - Contacts  │  │ - Bookings   │  │ - Login/     │           │
│  └─────────────┘  │ - Services   │  │   Logout     │           │
│                   └──────────────┘  └──────────────┘           │
│                          │                    │                 │
│                          └─────────┬──────────┘                 │
│                                    │                            │
│                          ┌─────────▼──────────┐                 │
│                          │   API Client       │                 │
│                          │   (axios)          │                 │
│                          │   - Interceptors   │                 │
│                          │   - JWT Auth       │                 │
│                          └─────────┬──────────┘                 │
└────────────────────────────────────┼────────────────────────────┘
                                     │
                                     │ HTTP Requests
                                     │ Authorization: Bearer <token>
                                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                      GO BACKEND                                 │
│                   http://localhost:8080                         │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │              Gin Router + CORS Middleware               │   │
│  └─────────────────────────────────────────────────────────┘   │
│           │                    │                    │           │
│    ┌──────▼──────┐      ┌─────▼─────┐      ┌──────▼──────┐    │
│    │   Public    │      │   Auth    │      │   Admin     │    │
│    │   Routes    │      │   Routes  │      │   Routes    │    │
│    │             │      │           │      │             │    │
│    │ /api/ping   │      │ /login    │      │ /dashboard  │    │
│    │ /bookings   │      │ /register │      │ /bookings   │    │
│    │ /services   │      │           │      │ /services   │    │
│    └──────┬──────┘      └─────┬─────┘      └──────┬──────┘    │
│           │                    │                    │           │
│           └────────────────────┼────────────────────┘           │
│                                │                                │
│                     ┌──────────▼──────────┐                     │
│                     │     Handlers        │                     │
│                     │  (HTTP Layer)       │                     │
│                     └──────────┬──────────┘                     │
│                                │                                │
│                     ┌──────────▼──────────┐                     │
│                     │     Services        │                     │
│                     │ (Business Logic)    │                     │
│                     └──────────┬──────────┘                     │
│                                │                                │
│                     ┌──────────▼──────────┐                     │
│                     │    Repository       │                     │
│                     │   (Data Access)     │                     │
│                     │  - In-Memory Maps   │                     │
│                     │  - Thread-Safe      │                     │
│                     └──────────┬──────────┘                     │
│                                │                                │
│                     ┌──────────▼──────────┐                     │
│                     │     Data Models     │                     │
│                     │  - Booking          │                     │
│                     │  - User             │                     │
│                     │  - Service          │                     │
│                     └─────────────────────┘                     │
└─────────────────────────────────────────────────────────────────┘
```

---

## 📊 Поток данных

### 1. Создание бронирования (Public)

```
┌──────────┐      ┌──────────┐      ┌──────────┐      ┌──────────┐
│  User    │      │ Frontend │      │ Backend  │      │  Memory  │
│          │      │ Booking  │      │   API    │      │  Storage │
│          │      │  Form    │      │          │      │          │
└────┬─────┘      └────┬─────┘      └────┬─────┘      └────┬─────┘
     │                 │                 │                 │
     │ 1. Заполняет    │                 │                 │
     │    форму        │                 │                 │
     ├────────────────>│                 │                 │
     │                 │                 │                 │
     │                 │ 2. POST         │                 │
     │                 │    /bookings    │                 │
     │                 ├────────────────>│                 │
     │                 │                 │                 │
     │                 │                 │ 3. Validate     │
     │                 │                 │    & Create     │
     │                 │                 ├────────────────>│
     │                 │                 │                 │
     │                 │                 │ 4. Booking ID   │
     │                 │                 │<────────────────┤
     │                 │                 │                 │
     │                 │ 5. Success      │                 │
     │                 │    Response     │                 │
     │                 │<────────────────┤                 │
     │                 │                 │                 │
     │ 6. Показывает   │                 │                 │
     │    уведомление  │                 │                 │
     │<────────────────┤                 │                 │
     │                 │                 │                 │
```

### 2. Аутентификация администратора

```
┌──────────┐      ┌──────────┐      ┌──────────┐      ┌──────────┐
│  Admin   │      │ Frontend │      │ Backend  │      │   JWT    │
│          │      │  Login   │      │   Auth   │      │  Token   │
│          │      │  Page    │      │  Service │      │          │
└────┬─────┘      └────┬─────┘      └────┬─────┘      └────┬─────┘
     │                 │                 │                 │
     │ 1. Вводит       │                 │                 │
     │    credentials  │                 │                 │
     ├────────────────>│                 │                 │
     │                 │                 │                 │
     │                 │ 2. POST         │                 │
     │                 │    /auth/login  │                 │
     │                 ├────────────────>│                 │
     │                 │                 │                 │
     │                 │                 │ 3. Verify       │
     │                 │                 │    password     │
     │                 │                 ├────────────────>│
     │                 │                 │                 │
     │                 │                 │ 4. Generate     │
     │                 │                 │    JWT Token    │
     │                 │                 │<────────────────┤
     │                 │                 │                 │
     │                 │ 5. Token +      │                 │
     │                 │    User Data    │                 │
     │                 │<────────────────┤                 │
     │                 │                 │                 │
     │                 │ 6. Сохраняет    │                 │
     │                 │    в localStorage               │
     │                 │                 │                 │
     │ 7. Redirect     │                 │                 │
     │    на Dashboard │                 │                 │
     │<────────────────┤                 │                 │
     │                 │                 │                 │
```

### 3. Загрузка данных Dashboard (Protected)

```
┌──────────┐      ┌──────────┐      ┌──────────┐      ┌──────────┐
│  Admin   │      │ Frontend │      │ Backend  │      │  Memory  │
│          │      │Dashboard │      │   API    │      │  Storage │
│          │      │  Page    │      │          │      │          │
└────┬─────┘      └────┬─────┘      └────┬─────┘      └────┬─────┘
     │                 │                 │                 │
     │ 1. Открывает    │                 │                 │
     │    страницу     │                 │                 │
     ├────────────────>│                 │                 │
     │                 │                 │                 │
     │                 │ 2. GET          │                 │
     │                 │    /admin/      │                 │
     │                 │    dashboard    │                 │
     │                 │    + JWT Token  │                 │
     │                 ├────────────────>│                 │
     │                 │                 │                 │
     │                 │                 │ 3. Validate     │
     │                 │                 │    JWT Token    │
     │                 │                 │    (Admin?)     │
     │                 │                 ├────────────────>│
     │                 │                 │                 │
     │                 │                 │ 4. Calculate    │
     │                 │                 │    Stats        │
     │                 │                 │<────────────────┤
     │                 │                 │                 │
     │                 │ 5. Stats Data   │                 │
     │                 │<────────────────┤                 │
     │                 │                 │                 │
     │ 6. Отображает   │                 │                 │
     │    статистику   │                 │                 │
     │<────────────────┤                 │                 │
     │                 │                 │                 │
```

---

## 🔐 JWT Authentication Flow

```
┌──────────────────────────────────────────────────────────────┐
│                    Authentication Flow                       │
└──────────────────────────────────────────────────────────────┘

1. LOGIN
   ┌─────────┐         ┌─────────┐
   │ Client  │         │ Server  │
   │         │  POST   │         │
   │         │ ───────>│         │
   │         │  /login │         │
   │         │  {email,│         │
   │         │  password}        │
   │         │         │ Verify credentials
   │         │         │ Generate JWT
   │         │  200 OK │         │
   │         │ <────── │         │
   │         │  {token,│         │
   │         │   user} │         │
   │  Store in localStorage     │
   └─────────┘         └─────────┘

2. PROTECTED REQUEST
   ┌─────────┐         ┌─────────┐
   │ Client  │         │ Server  │
   │  Get token from   │         │
   │  localStorage     │         │
   │         │  GET    │         │
   │         │ ───────>│         │
   │         │  /admin/│         │
   │         │  dashboard        │
   │         │  Authorization:   │
   │         │  Bearer <token>   │
   │         │         │ Validate JWT
   │         │         │ Check role
   │         │  200 OK │         │
   │         │ <────── │         │
   │         │  {data} │         │
   └─────────┘         └─────────┘

3. TOKEN EXPIRED
   ┌─────────┐         ┌─────────┐
   │ Client  │         │ Server  │
   │         │  GET    │         │
   │         │ ───────>│         │
   │         │  /admin/│         │
   │         │  dashboard        │
   │         │  Authorization:   │
   │         │  Bearer <expired> │
   │         │         │ JWT Invalid
   │         │  401    │         │
   │         │ <────── │         │
   │         │  {error}│         │
   │  Redirect to      │         │
   │  /admin/login     │         │
   └─────────┘         └─────────┘
```

---

## 📁 Структура API запросов

### Public Endpoints

```javascript
// POST /api/bookings
{
  "customer_name": "Иван Иванов",
  "customer_phone": "+7 (999) 123-45-67",
  "car_brand": "BMW",
  "car_model": "X5",
  "service_id": "svc_3",
  "comment": "Нужно помыть салон"
}

// Response 201 Created
{
  "message": "Booking created successfully",
  "booking": {
    "id": "uuid-string",
    "customer_name": "Иван Иванов",
    "customer_phone": "+7 (999) 123-45-67",
    "car_brand": "BMW",
    "car_model": "X5",
    "service_id": "svc_3",
    "service_name": "Химчистка салона",
    "price": 8000,
    "status": "pending",
    "comment": "Нужно помыть салон",
    "created_at": "2026-03-26T10:00:00Z",
    "updated_at": "2026-03-26T10:00:00Z"
  }
}
```

### Auth Endpoint

```javascript
// POST /api/auth/login
{
  "email": "admin@deteleng.com",
  "password": "admin123"
}

// Response 200 OK
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "admin_1",
    "email": "admin@deteleng.com",
    "role": "admin",
    "created_at": "2026-03-26T10:00:00Z"
  }
}
```

### Protected Admin Endpoint

```javascript
// GET /api/admin/dashboard
// Headers: Authorization: Bearer <jwt-token>

// Response 200 OK
{
  "stats": {
    "total_bookings": 15,
    "pending_bookings": 3,
    "confirmed_bookings": 5,
    "completed_bookings": 6,
    "total_revenue": 75000,
    "recent_bookings": [...],
    "status_breakdown": {
      "pending": 3,
      "confirmed": 5,
      "completed": 6,
      "cancelled": 1
    }
  }
}
```

---

## 🛡️ CORS Configuration

```go
// Backend (main.go)
cors.Config{
  AllowOrigins:     []string{
    "http://localhost:5173",  // Vite dev server
    "http://localhost:3000"   // React production
  },
  AllowMethods:     []string{
    "GET", "POST", "PUT", "DELETE", "OPTIONS"
  },
  AllowHeaders:     []string{
    "Origin", "Content-Type", "Authorization"
  },
  ExposeHeaders:    []string{"Content-Length"},
  AllowCredentials: true,
  MaxAge:           12 * 3600,
}
```

---

## 📦 Data Models

### Booking Model
```go
type Booking struct {
  ID            string        `json:"id"`
  CustomerName  string        `json:"customer_name"`
  CustomerPhone string        `json:"customer_phone"`
  CarBrand      string        `json:"car_brand"`
  CarModel      string        `json:"car_model"`
  ServiceID     string        `json:"service_id"`
  ServiceName   string        `json:"service_name"`
  Price         float64       `json:"price"`
  Status        BookingStatus `json:"status"`  // pending|confirmed|completed|cancelled
  Comment       string        `json:"comment,omitempty"`
  CreatedAt     time.Time     `json:"created_at"`
  UpdatedAt     time.Time     `json:"updated_at"`
}
```

### User Model
```go
type User struct {
  ID        string    `json:"id"`
  Email     string    `json:"email"`
  Password  string    `json:"-"`  // Never sent to client
  Role      string    `json:"role"`  // admin|user
  CreatedAt time.Time `json:"created_at"`
}
```

### Service Model
```go
type Service struct {
  ID          string  `json:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
  Price       float64 `json:"price"`
  Duration    int     `json:"duration_minutes"`
  Available   bool    `json:"available"`
}
```

---

## 🔄 State Management

### Frontend (React Context)
```javascript
// AuthContext.jsx
{
  user: {
    id: "admin_1",
    email: "admin@deteleng.com",
    role: "admin"
  },
  loading: false,
  isAuthenticated: true,
  isAdmin: true,
  login: (userData, token) => {...},
  logout: () => {...}
}
```

### Backend (In-Memory Repository)
```go
// Thread-safe maps
type InMemoryRepository struct {
  mu        sync.RWMutex
  bookings  map[string]*Booking
  users     map[string]*User
  services  map[string]*Service
}
```

---

## ✅ Checklist интеграции

- [x] Backend API создан и работает
- [x] Frontend API client настроен
- [x] CORS настроен правильно
- [x] JWT аутентификация работает
- [x] Booking форма отправляет данные в API
- [x] Admin login интегрирован с backend
- [x] Dashboard загружает статистику из API
- [x] Bookings management работает через API
- [x] Services management работает через API
- [x] Vite proxy настроен
- [x] Docker конфигурация готова

---

## 🎯 Точки расширения

1. **Database Integration** - Заменить in-memory repository на PostgreSQL
2. **Email Notifications** - Добавить отправку email при создании бронирования
3. **File Upload** - Загрузка изображений для услуг
4. **Booking Calendar** - Календарь для выбора даты/времени
5. **User Management** - Создание/удаление администраторов
6. **Analytics** - Расширенная аналитика и отчеты
