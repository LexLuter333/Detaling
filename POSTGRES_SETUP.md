# 🚀 Инструкция по запуску Deteleng на PostgreSQL

## 📋 Требования

### Для локальной разработки:
- **Go 1.22+**
- **Node.js 18+**
- **PostgreSQL 15+**

### Для Docker:
- **Docker**
- **Docker Compose**

---

## 🔧 Вариант 1: Запуск через Docker (Рекомендуется)

### Шаг 1: Запуск всех сервисов

```bash
cd D:\Deteleng

# Запустить базу данных, backend и frontend
docker-compose up -d

# Просмотр логов
docker-compose logs -f
```

### Шаг 2: Проверка работы

- **Frontend:** http://localhost:5173
- **Backend API:** http://localhost:8080/api/ping
- **Admin Panel:** http://localhost:5173/admin/login

### Шаг 3: Остановка

```bash
docker-compose down

# С удалением данных базы данных
docker-compose down -v
```

---

## 🔧 Вариант 2: Локальная разработка

### Шаг 1: Установка PostgreSQL

#### Windows:
1. Скачайте с https://www.postgresql.org/download/windows/
2. Установите PostgreSQL 15
3. Запомните пароль пользователя `postgres`

#### Или через Docker:
```bash
docker run -d \
  --name deteleng-db \
  -e POSTGRES_USER=deteleng \
  -e POSTGRES_PASSWORD=deteleng123 \
  -e POSTGRES_DB=deteleng \
  -p 5432:5432 \
  postgres:15-alpine
```

### Шаг 2: Создание базы данных

```bash
# Подключение к PostgreSQL
psql -U postgres

# Создание БД и пользователя
CREATE DATABASE deteleng;
CREATE USER deteleng WITH PASSWORD 'deteleng123';
GRANT ALL PRIVILEGES ON DATABASE deteleng TO deteleng;
\c deteleng
GRANT ALL ON SCHEMA public TO deteleng;
\q
```

### Шаг 3: Запуск миграций

```bash
cd D:\Deteleng\backend

# Установить переменную окружения
set DATABASE_URL=postgres://deteleng:deteleng123@localhost:5432/deteleng?sslmode=disable

# Запустить миграции вручную (опционально)
psql -U deteleng -d deteleng -f database/migrations/001_create_tables.sql
psql -U deteleng -d deteleng -f database/migrations/002_seed_data.sql
```

### Шаг 4: Запуск Backend

```bash
cd D:\Deteleng\backend

# Установить зависимости
go mod download

# Запустить сервер
go run ./cmd/api
```

**Ожидайте:** `✅ Server starting on port 8080`

### Шаг 5: Запуск Frontend (новый терминал)

```bash
cd D:\Deteleng\frontend

# Установить зависимости
npm install

# Запустить dev-сервер
npm run dev
```

**Ожидайте:** `Local: http://localhost:5173/`

---

## 🗄️ Файлы базы данных

### Миграции:

| Файл | Описание |
|------|----------|
| `backend/database/migrations/001_create_tables.sql` | Создание всех таблиц |
| `backend/database/migrations/002_seed_data.sql` | Начальные данные (админ, услуги) |
| `backend/database/services_data.sql` | Все 21 услуга для импорта |

### Таблицы:

1. **users** - Пользователи (администраторы)
2. **services** - Услуги (21 услуга)
3. **bookings** - Бронирования
4. **reviews** - Отзывы
5. **review_sources** - Источники отзывов

---

## 🔐 Доступы

### Админ-панель:
- **URL:** http://localhost:5173/admin/login
- **Email:** `admin@deteleng.com`
- **Пароль:** `admin123`

### База данных (PostgreSQL):
- **Host:** localhost:5432
- **Database:** deteleng
- **User:** deteleng
- **Password:** deteleng123

---

## 🧪 Тестирование

### 1. Проверка API

```bash
# Health check
curl http://localhost:8080/api/ping
# Ответ: {"message":"pong"}

# Получить все услуги
curl http://localhost:8080/api/services

# Создать бронирование
curl -X POST http://localhost:8080/api/bookings \
  -H "Content-Type: application/json" \
  -d "{\"customer_name\":\"Иван\",\"customer_phone\":\"+79990000000\",\"car_brand\":\"BMW\",\"car_model\":\"X5\",\"service_id\":\"svc_1\"}"

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"admin@deteleng.com\",\"password\":\"admin123\"}"
```

### 2. Проверка базы данных

```bash
# Подключение к БД
psql -U deteleng -d deteleng

# Проверить количество услуг
SELECT COUNT(*) FROM services;

# Проверить услуги
SELECT id, name, price FROM services ORDER BY price;

# Проверить бронирования
SELECT * FROM bookings ORDER BY created_at DESC LIMIT 10;

# Выйти
\q
```

### 3. Проверка frontend

1. Откройте http://localhost:5173
2. Заполните форму записи
3. Проверьте, что заявка появилась в базе:
   ```sql
   SELECT * FROM bookings ORDER BY created_at DESC LIMIT 1;
   ```

4. Войдите в админ-панель: http://localhost:5173/admin/login
5. Проверьте разделы:
   - **Dashboard** - статистика
   - **Бронирования** - управление заявками
   - **Услуги** - CRUD операции (создание, редактирование, удаление)
   - **Отзывы** - парсинг и модерация

---

## 📝 Управление услугами через админ-панель

### Создание услуги:
1. Перейдите в `/admin/services`
2. Нажмите "+ Добавить услугу"
3. Заполните форму:
   - Название
   - Описание
   - Цена (₽)
   - Длительность (мин)
   - Доступно для записи (чекбокс)
4. Нажмите "Создать"

### Редактирование:
1. Нажмите ✏️ на нужной услуге
2. Измените данные
3. Нажмите "Сохранить"

### Удаление:
1. Нажмите 🗑️ на нужной услуге
2. Подтвердите удаление

Или через SQL:
```sql
-- Отключить услугу
UPDATE services SET available = FALSE WHERE id = 'svc_1';

-- Удалить услугу
DELETE FROM services WHERE id = 'svc_1';

-- Добавить новую услугу
INSERT INTO services (id, name, description, price, duration, available)
VALUES ('svc_22', 'Новая услуга', 'Описание', 5000, 120, TRUE);
```

---

## 🔄 Автоматическое удаление старых бронирований

Backend автоматически удаляет выполненные бронирования старше 7 дней.

Запускается раз в 24 часа.

Вручную:
```sql
DELETE FROM bookings 
WHERE status = 'completed' 
AND created_at < NOW() - INTERVAL '7 days';
```

---

## 🛠️ Переменные окружения

### Backend (.env или в docker-compose):

```bash
SERVER_PORT=8080
JWT_SECRET=deteleng-secret-key-2026
ADMIN_EMAIL=admin@deteleng.com
ADMIN_PASSWORD=admin123
DATABASE_URL=postgres://deteleng:deteleng123@localhost:5432/deteleng?sslmode=disable
```

### Frontend:
- Прокси настроен в `vite.config.js`
- API URL: `http://localhost:8080/api`

---

## 📊 Структура проекта

```
D:\Deteleng\
├── backend/
│   ├── cmd/api/main.go          # Точка входа
│   ├── internal/
│   │   ├── database/            # PostgreSQL подключение
│   │   ├── handlers/            # HTTP обработчики
│   │   ├── middleware/          # Middleware (auth, cors)
│   │   ├── models/              # Модели данных
│   │   ├── repository/          # Repository layer
│   │   └── services/            # Business logic
│   ├── database/
│   │   ├── migrations/          # SQL миграции
│   │   │   ├── 001_create_tables.sql
│   │   │   └── 002_seed_data.sql
│   │   └── services_data.sql    # Данные услуг
│   ├── go.mod
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── api/api.js           # API клиент
│   │   ├── pages/admin/         # Админ-панель
│   │   └── ...
│   ├── package.json
│   └── Dockerfile
└── docker-compose.yml
```

---

## 🐛 Решение проблем

### Backend не подключается к базе:
```bash
# Проверьте, что PostgreSQL запущен
docker ps | grep deteleng-db

# Или для локального
pg_isready -h localhost -p 5432
```

### Ошибка миграций:
```bash
# Подключитесь к БД и проверьте таблицы
psql -U deteleng -d deteleng
\dt
```

### Frontend не видит backend:
- Убедитесь, что backend запущен на порту 8080
- Проверьте CORS в логах браузера (F12)

### Порт занят:
```bash
# Windows
netstat -ano | findstr :5432
netstat -ano | findstr :8080
netstat -ano | findstr :5173

# Убить процесс
taskkill /PID <PID> /F
```

---

## 📞 Поддержка

При возникновении проблем:
1. Проверьте логи backend в консоли
2. Проверьте консоль браузера (F12)
3. Убедитесь, что все сервисы запущены
4. Проверьте подключение к базе данных
