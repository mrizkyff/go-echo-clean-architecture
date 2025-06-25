This is for golang learning.

# Go Echo Clean Architecture

Sebuah aplikasi web backend yang dibangun menggunakan Go dengan arsitektur clean architecture, mengimplementasikan URL shortener service dengan sistem autentikasi yang lengkap.

## 🚀 Fitur Utama

- **URL Shortener**: Layanan untuk memperpendek URL dengan tracking user
- **Sistem Autentikasi**: Login, register, dan logout dengan JWT token
- **User Management**: CRUD operations untuk manajemen user
- **Access Logging**: Pencatatan aktivitas user dengan RabbitMQ
- **API Documentation**: Swagger/OpenAPI documentation
- **Middleware**: CORS, Error Handling, Recovery, Authentication
- **Caching**: Redis untuk session management dan caching
- **Message Queue**: RabbitMQ untuk asynchronous processing

## 🛠️ Tech Stack

- **Framework**: Echo v4 (Go Web Framework)
- **Database**: PostgreSQL dengan GORM ORM
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Configuration**: Viper dengan Consul support
- **Authentication**: JWT (JSON Web Token)
- **Documentation**: Swagger/OpenAPI
- **Monitoring**: Elastic APM
- **Testing**: Testify
- **Deployment**: Kubernetes ready

## 📁 Struktur Project

```
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── dto/            # Data Transfer Objects
│   ├── handlers/       # HTTP handlers
│   ├── services/       # Business logic
│   ├── repositories/   # Data access layer
│   ├── models/         # Domain models
│   ├── middlewares/    # HTTP middlewares
│   ├── routes/         # Route definitions
│   ├── utils/          # Utility functions
│   └── validation/     # Input validation
├── pkg/                # Public packages
│   ├── config/         # Configuration management
│   ├── database/       # Database connections
│   └── rabbitmq/       # Message queue setup
└── static/             # Static files
```

## 🔧 Konfigurasi

Project ini menggunakan Viper untuk configuration management dengan support untuk:
- Environment variables
- Configuration files
- Remote configuration (Consul)

### Environment Variables

```env
# Database
POSTGRES_HOST=localhost
POSTGRES_USERNAME=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DBNAME=your_database
POSTGRES_PORT=5432

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DBINDEX=0

# RabbitMQ
RABBITMQ_USERNAME=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_VHOST=/

# JWT
JWT_SECRET_KEY=your_secret_key
JWT_EXPIRATION_HOURS=24
JWT_REFRESH_TOKEN_DAYS=7
JWT_ISSUER=go-echo-clean-architecture
```

## 🚀 Cara Menjalankan

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd go-echo-clean-architecture
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup database dan services**
   - PostgreSQL
   - Redis
   - RabbitMQ

4. **Set environment variables**
   ```bash
   cp .env.example .env
   # Edit .env file dengan konfigurasi yang sesuai
   ```

5. **Run application**
   ```bash
   go run cmd/server/main.go
   ```

6. **Akses aplikasi**
   - API: `http://localhost:1323`
   - Swagger Documentation: `http://localhost:1323/swagger/`
   - Static files: `http://localhost:1323/`

## 📚 API Endpoints

### Authentication
- `POST /auth/login` - User login
- `POST /auth/register` - User registration
- `POST /auth/logout` - User logout

### User Management
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

### Link Management
- `POST /links` - Create short link
- `GET /links` - Get user's links
- `GET /links/:id` - Get link by ID
- `PUT /links/:id` - Update link
- `DELETE /links/:id` - Delete link

### Access Logs
- `GET /access-logs` - Get access logs

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/tests/models/
```

## 🏗️ Arsitektur

Project ini mengimplementasikan **Clean Architecture** dengan layer:

1. **Handlers** - HTTP request/response handling
2. **Services** - Business logic
3. **Repositories** - Data access abstraction
4. **Models** - Domain entities

### Design Patterns
- **Dependency Injection** - Untuk loose coupling
- **Repository Pattern** - Untuk data access abstraction
- **Middleware Pattern** - Untuk cross-cutting concerns
- **Publisher-Subscriber** - Untuk asynchronous processing

## 🔒 Security Features

- JWT-based authentication
- Password hashing dengan bcrypt
- CORS configuration
- Input validation
- Error handling middleware
- SQL injection protection (GORM)

## 📊 Monitoring & Logging

- **Structured Logging** dengan Logrus
- **APM Integration** dengan Elastic APM
- **Access Logging** dengan RabbitMQ
- **Error Tracking** dengan middleware

## 🚢 Deployment

Project ini siap untuk deployment di:
- **Kubernetes** cluster
- **Docker** containers
- **Cloud platforms** (AWS, GCP, Azure)

## 🤝 Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## 📄 License

MIT License - lihat file LICENSE untuk detail lengkap.
