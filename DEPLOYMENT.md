# 🚀 Production Deployment Guide

This guide covers deploying Deteleng to production.

---

## 📋 Prerequisites

- Docker & Docker Compose
- Domain name (optional but recommended)
- SSL certificate (recommended via reverse proxy)
- PostgreSQL database (or use included container)

---

## 🔧 Quick Start (5 minutes)

### 1. Clone and Configure

```bash
cd /path/to/deteleng

# Copy environment templates
cp .env.example .env
cp frontend/.env.example frontend/.env
```

### 2. Configure environment variables

Edit `.env` file:

```bash
# Required: Strong random secret (min 32 characters)
JWT_SECRET=your-super-secret-random-string-here

# Required: Strong admin password
ADMIN_PASSWORD=your-secure-password-here

# Required: Your production domain(s)
CORS_ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com

# Required: Database connection
DATABASE_URL=postgres://user:password@db-host:5432/dbname?sslmode=require

# Frontend API URL (in frontend/.env)
VITE_API_BASE_URL=https://api.yourdomain.com/api
```

### 3. Deploy with Docker Compose

```bash
# Production deployment
docker-compose -f docker-compose.prod.yml up -d

# Check logs
docker-compose -f docker-compose.prod.yml logs -f

# Stop deployment
docker-compose -f docker-compose.prod.yml down
```

### 4. Verify deployment

```bash
# Check health
curl http://localhost/api/ping
# Expected: {"message":"pong"}

# Check frontend
curl http://localhost/
# Expected: HTML content
```

---

## 🔐 Security Checklist

Before deploying to production:

- [ ] Change default admin password (`ADMIN_PASSWORD`)
- [ ] Generate secure `JWT_SECRET` (min 32 random characters)
- [ ] Set strong database password
- [ ] Configure CORS for your domain only
- [ ] Enable HTTPS/SSL
- [ ] Remove development ports from exposure
- [ ] Set up firewall rules
- [ ] Enable database SSL mode (`sslmode=require`)

### Generate secure secrets:

```bash
# JWT Secret (32 random characters)
openssl rand -base64 32

# Admin password
openssl rand -base64 24
```

---

## 🌐 Deployment Options

### Option 1: Single Server (Docker Compose)

Best for small to medium deployments.

```bash
docker-compose -f docker-compose.prod.yml up -d
```

**Services:**
- Frontend: Port 80
- Backend: Port 8080
- Database: Internal only (no exposed port)

### Option 2: Separate Backend and Frontend

Deploy frontend on CDN/static hosting, backend on separate server.

**Backend server (.env):**
```bash
DATABASE_URL=postgres://...
CORS_ALLOWED_ORIGINS=https://yourdomain.com
```

**Frontend build:**
```bash
cd frontend
VITE_API_BASE_URL=https://api.yourdomain.com/api
npm run build

# Deploy dist/ folder to your static hosting
```

### Option 3: Kubernetes

For large-scale deployments (configuration not included).

---

## 🗄️ Database Setup

### Using included PostgreSQL container:

```bash
docker-compose -f docker-compose.prod.yml up -d db
```

### Using external PostgreSQL:

1. Create database:
```sql
CREATE DATABASE deteleng;
CREATE USER deteleng WITH PASSWORD 'secure-password';
GRANT ALL PRIVILEGES ON DATABASE deteleng TO deteleng;
```

2. Update `.env`:
```bash
DATABASE_URL=postgres://deteleng:secure-password@db-host:5432/deteleng?sslmode=require
```

---

## 🔍 Monitoring & Maintenance

### View logs:

```bash
# All services
docker-compose -f docker-compose.prod.yml logs -f

# Specific service
docker-compose -f docker-compose.prod.yml logs -f backend
```

### Health checks:

```bash
# Backend health
curl http://localhost:8080/api/ping

# Frontend health
curl http://localhost/

# Database health
docker exec deteleng-db-prod pg_isready
```

### Backup database:

```bash
# Backup
docker exec deteleng-db-prod pg_dump -U deteleng deteleng > backup.sql

# Restore
docker exec -i deteleng-db-prod psql -U deteleng deteleng < backup.sql
```

### Update deployment:

```bash
# Pull latest changes
git pull

# Rebuild and restart
docker-compose -f docker-compose.prod.yml up -d --build
```

---

## 🐛 Troubleshooting

### Frontend cannot connect to backend:

1. Check `VITE_API_BASE_URL` in `frontend/.env`
2. Verify backend is running: `curl http://localhost:8080/api/ping`
3. Check CORS settings in backend `.env`

### Database connection errors:

1. Verify `DATABASE_URL` format
2. Check database is accessible
3. Verify credentials
4. Check SSL mode setting

### CORS errors in browser:

1. Set `CORS_ALLOWED_ORIGINS` to your frontend domain
2. Include protocol: `https://yourdomain.com` (not just `yourdomain.com`)
3. Restart backend after changing CORS settings

### Container won't start:

```bash
# Check logs
docker-compose -f docker-compose.prod.yml logs SERVICE_NAME

# Check container status
docker-compose -f docker-compose.prod.yml ps

# Restart specific service
docker-compose -f docker-compose.prod.yml restart SERVICE_NAME
```

---

## 📊 Environment Variables Reference

### Backend (.env)

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `SERVER_PORT` | No | `8080` | Backend server port |
| `DATABASE_URL` | **Yes** | - | PostgreSQL connection string |
| `JWT_SECRET` | **Yes** | - | Secret key for JWT tokens |
| `ADMIN_EMAIL` | No | `admin@deteleng.com` | Admin login email |
| `ADMIN_PASSWORD` | **Yes** | - | Admin login password |
| `CORS_ALLOWED_ORIGINS` | **Yes** | - | Comma-separated frontend URLs |

### Frontend (frontend/.env)

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `VITE_API_BASE_URL` | **Yes** | - | Backend API URL |

---

## 🎯 Post-Deployment Checklist

After deploying:

- [ ] Test login with admin credentials
- [ ] Create a test booking from frontend
- [ ] Verify booking appears in admin panel
- [ ] Test service management (create/edit/delete)
- [ ] Test review parsing
- [ ] Check all API endpoints work
- [ ] Verify database migrations ran
- [ ] Set up monitoring/alerting
- [ ] Configure log aggregation
- [ ] Set up automated backups

---

## 📞 Support

If you encounter issues:

1. Check logs: `docker-compose -f docker-compose.prod.yml logs -f`
2. Verify environment variables are set correctly
3. Check container health: `docker-compose -f docker-compose.prod.yml ps`
4. Review this guide for common issues

---

## 📝 Example .env for Production

```bash
# Server
SERVER_PORT=8080

# Database (external PostgreSQL example)
DATABASE_URL=postgres://deteleng:SuperSecurePass123@db.example.com:5432/deteleng?sslmode=require

# Security
JWT_SECRET=a8f5f167f44f4964e6c998dee827110c38b5e1f8a0e7e1e8f5f167f44f4964e6

# Admin
ADMIN_EMAIL=admin@yourdomain.com
ADMIN_PASSWORD=VerySecurePassword123!

# CORS (your production domains)
CORS_ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com
```
