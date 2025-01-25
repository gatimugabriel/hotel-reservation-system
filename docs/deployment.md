# Deployment Guide

## Overview
This guide covers deploying the Hotel Reservation System in both manual and containerized environments.

## Manual Deployment

### Prerequisites
- Go 1.19+
- PostgreSQL 14+
- Domain name (for production)
- SSL certificate (for production)

### Environment Configuration
1. Create configuration file
```bash
cp .env.example .env
```

2. Configure the following variables in `.env`:
```
# Server
PORT=8080
ENV=production
API_SECRET=your-jwt-secret-key

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=hotel_reservation
DB_USER=postgres
DB_PASSWORD=your-db-password

# Redis (for caching)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# SMTP (for emails)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASSWORD=your-app-specific-password

# Payment Gateway
STRIPE_SECRET_KEY=your-stripe-secret-key
STRIPE_WEBHOOK_SECRET=your-stripe-webhook-secret
```

### Database Setup
1. Install PostgreSQL
```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib

# CentOS/RHEL
sudo yum install postgresql postgresql-server
```

2. Create database and user
```bash
sudo -u postgres psql
CREATE DATABASE hotel_reservation;
CREATE USER hotelapp WITH PASSWORD 'your-password';
GRANT ALL PRIVILEGES ON DATABASE hotel_reservation TO hotelapp;
```

3. Run migrations
```bash
make migrate-up
```

### Application Deployment

1. Build the application
```bash
go build -o hotel-api ./cmd/api
```

2. Set up systemd service
```bash
sudo nano /etc/systemd/system/hotel-api.service
```

Add the following content:
```ini
[Unit]
Description=Hotel Reservation API
After=network.target postgresql.service

[Service]
User=hotelapp
WorkingDirectory=/path/to/app
Environment=ENV=production
ExecStart=/path/to/app/hotel-api
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

3. Start the service
```bash
sudo systemctl enable hotel-api
sudo systemctl start hotel-api
```

4. Set up Nginx as reverse proxy
```bash
sudo nano /etc/nginx/sites-available/hotel-api
```

Add the following configuration:
```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

5. Enable the site and restart Nginx
```bash
sudo ln -s /etc/nginx/sites-available/hotel-api /etc/nginx/sites-enabled/
sudo systemctl restart nginx
```

## Docker Deployment

### Prerequisites
- Docker
- Docker Compose
- Domain name (for production)
- SSL certificate (for production)

### Configuration
1. Create `.env` file as described in Manual Deployment section

2. Update Docker Compose configuration if needed
```yaml
# docker-compose.yml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    env_file: .env
    depends_on:
      - db
      - redis
    restart: unless-stopped

  db:
    image: postgres:14
    volumes:
      - postgres_data:/var/lib/postgresql/data
    environment:
      POSTGRES_DB: hotel_reservation
      POSTGRES_USER: hotelapp
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    restart: unless-stopped

  redis:
    image: redis:alpine
    volumes:
      - redis_data:/data
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
```

### Deployment Steps

1. Build and start services
```bash
docker-compose up -d --build
```

2. Run database migrations
```bash
docker-compose exec api make migrate-up
```

3. Monitor logs
```bash
docker-compose logs -f api
```

## Production Considerations

### Security
1. Enable SSL/TLS
2. Configure firewall rules
3. Regular security updates
4. Implement rate limiting
5. Set up monitoring and alerting

### Scaling
1. Use container orchestration (Kubernetes/ECS)
2. Implement caching strategy
3. Set up load balancer
4. Database replication
5. CDN for static assets

### Monitoring
1. Set up application monitoring (e.g., Prometheus + Grafana)
2. Configure error tracking (e.g., Sentry)
3. Implement health checks
4. Set up log aggregation (e.g., ELK Stack)
5. Monitor system metrics

### Backup Strategy
1. Regular database backups
2. Test backup restoration
3. Implement backup rotation policy
4. Monitor backup success/failure
5. Document recovery procedures

## CI/CD Pipeline

Example GitHub Actions workflow:
```yaml
name: CI/CD

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.19
    - name: Run tests
      run: make test
    - name: Run linter
      run: make lint

  deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
    - name: Deploy to production
      run: |
        echo "Add deployment steps here"
```

## Rollback Procedures

### Manual Rollback
1. Keep previous version tagged
2. Document version changes
3. Test rollback procedures
4. Monitor after rollback
5. Update DNS if needed

### Docker Rollback
```bash
# Tag the current version
docker tag hotel-api:latest hotel-api:backup

# Roll back to previous version
docker-compose down
docker-compose up -d --build

# If needed, revert database
make migrate-down
```

## Troubleshooting

### Common Issues
1. Database connection issues
2. Redis connection issues
3. SSL certificate problems
4. Permission issues
5. Memory/CPU constraints

### Debugging Steps
1. Check application logs
2. Verify environment variables
3. Check system resources
4. Test network connectivity
5. Verify service dependencies

### Health Checks
```bash
# API health
curl -f http://localhost:8080/health

# Database health
docker-compose exec db pg_isready

# Redis health
docker-compose exec redis redis-cli ping
```