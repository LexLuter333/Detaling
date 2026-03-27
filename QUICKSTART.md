# ⚡ Быстрый старт Deteleng на PostgreSQL

## 🎯 3 команды для запуска

### 1. Запустить PostgreSQL (если нет установленного)
```bash
docker run -d --name deteleng-db -e POSTGRES_USER=deteleng -e POSTGRES_PASSWORD=deteleng123 -e POSTGRES_DB=deteleng -p 5432:5432 postgres:15-alpine
```

### 2. Запустить Backend
```bash
cd D:\Deteleng\backend
go run ./cmd/api
```

### 3. Запустить Frontend (новый терминал)
```bash
cd D:\Deteleng\frontend
npm run dev
```

---

## ✅ Проверка

1. **Откройте:** http://localhost:5173
2. **Админ-панель:** http://localhost:5173/admin/login
3. **Логин:** `admin@deteleng.com`
4. **Пароль:** `admin123`

---

## 📦 Что работает:

- ✅ **21 услуга** в базе данных
- ✅ **Бронирования** сохраняются в PostgreSQL
- ✅ **Админ-панель** для управления услугами (CRUD)
- ✅ **Автоматическое удаление** старых бронирований (через 7 дней)
- ✅ **Отзывы** с парсингом из источников

---

## 🗄️ База данных

**Таблицы:**
- `users` - администраторы
- `services` - услуги (21 запись)
- `bookings` - бронирования
- `reviews` - отзывы
- `review_sources` - источники отзывов

**Проверка:**
```bash
docker exec -it deteleng-db psql -U deteleng -d deteleng -c "SELECT COUNT(*) FROM services;"
```

---

## 📁 SQL файлы

| Файл | Зачем |
|------|-------|
| `backend/database/migrations/001_create_tables.sql` | Создание таблиц |
| `backend/database/migrations/002_seed_data.sql` | Админ + услуги |
| `backend/database/services_data.sql` | Все 21 услуга |

---

## 🔄 Docker Compose (альтернатива)

```bash
cd D:\Deteleng
docker-compose up -d
```

Запускает: PostgreSQL + Backend + Frontend

---

## 🛑 Остановка

```bash
# Docker Compose
docker-compose down

# PostgreSQL Docker
docker stop deteleng-db
docker rm deteleng-db
```
