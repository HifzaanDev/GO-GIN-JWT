# 🎯 Go JWT Gin - Event Management API

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=gin&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

A robust, secure, and scalable RESTful API built with Go, Gin framework, and JWT authentication for managing events and attendees.

</div>

---

## 🚀 Features

- **🔐 JWT Authentication**: Secure user authentication with JSON Web Tokens
- **👥 User Management**: Complete user registration and login system
- **🎪 Event Management**: Create, read, update, and delete events
- **🎫 Attendee System**: Manage event attendees with full CRUD operations
- **🛡️ Middleware Protection**: Protected routes with authentication middleware
- **📊 SQLite Database**: Lightweight database with proper migrations
- **🔄 Hot Reload**: Development environment with Air for hot reloading
- **📝 API Documentation**: Swagger-ready API documentation comments

---

## 🏗️ Project Structure

```
Go-JWT-Gin/
├── api/                    # Main API handlers
│   ├── auth.go            # Authentication endpoints
│   ├── events.go          # Event management endpoints
│   ├── middleware.go      # JWT middleware
│   ├── routes.go          # Route definitions
│   ├── server.go          # HTTP server setup
│   ├── context.go         # Context utilities
│   └── main.go            # Application entry point
├── internal/              # Internal packages
│   ├── database/          # Database models and operations
│   │   ├── models.go      # Model definitions
│   │   ├── users.go       # User operations
│   │   ├── events.go      # Event operations
│   │   └── attendees.go   # Attendee operations
│   └── env/               # Environment utilities
├── migrate/               # Database migrations
│   ├── main.go            # Migration runner
│   └── migrations/        # SQL migration files
├── .air.toml              # Air configuration for hot reload
├── go.mod                 # Go module definition
└── data.db                # SQLite database file
```

---

## 🛠️ Prerequisites

- **Go 1.22+** - [Download Go](https://golang.org/dl/)
- **Git** - [Download Git](https://git-scm.com/downloads)

---

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/HifzaanDev/go-jwt-gin.git
cd go-jwt-gin
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Set Environment Variables
Create a `.env` file in the root directory:
```env
PORT=8080
JWT_SECRET=your-super-secret-jwt-key-here
```

### 4. Run Database Migrations
```bash
go run migrate/main.go
```

### 5. Start the Server
```bash
# Development with hot reload
air

# Or run directly
go run api/main.go
```

The server will start on `http://localhost:8080` 🎉

---

## 📚 API Endpoints

### 🔓 Public Routes

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/events` | Get all events |
| `GET` | `/api/v1/events/:id` | Get specific event |
| `GET` | `/api/v1/events/:id/attendees` | Get event attendees |
| `GET` | `/api/v1/attendees/:id/events` | Get events by attendee |
| `POST` | `/api/v1/auth/register` | Register new user |
| `POST` | `/api/v1/auth/login` | Login user |

### 🔒 Protected Routes (Requires JWT)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/events` | Create new event |
| `PUT` | `/api/v1/events/:id` | Update event |
| `DELETE` | `/api/v1/events/:id` | Delete event |
| `POST` | `/api/v1/events/:id/attendees/:userId` | Add attendee to event |
| `DELETE` | `/api/v1/events/:id/attendees/:userId` | Remove attendee from event |

---

## 🔍 API Usage Examples

### User Registration
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### User Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Event (Protected)
```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Tech Conference 2024",
    "description": "Annual technology conference",
    "date": "2024-12-25T10:00:00Z",
    "location": "San Francisco, CA"
  }'
```

### Get All Events
```bash
curl -X GET http://localhost:8080/api/v1/events
```

---

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `JWT_SECRET` | JWT signing secret | `some-secret-123456` |

### Database Configuration

The application uses SQLite by default. The database file is located at `data.db` in the project root. You can modify the database configuration in `api/main.go`.

---

## 🏃‍♂️ Development

### Hot Reload with Air

The project is configured with [Air](https://github.com/cosmtrek/air) for hot reloading during development:

```bash
# Install Air globally
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Database Migrations

To run migrations:
```bash
go run migrate/main.go
```

Migration files are located in `migrate/migrations/` directory.

---

## 🧪 Testing

### Manual Testing with curl

You can use the provided curl examples above or import the API into tools like:
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [Thunder Client](https://www.thunderclient.io/) (VS Code extension)

---

## 🛡️ Security Features

- **Password Hashing**: Uses bcrypt for secure password storage
- **JWT Authentication**: Stateless authentication with configurable expiration
- **Input Validation**: Request validation with Gin binding
- **Authorization Checks**: Route-level and resource-level authorization
- **CORS Ready**: Easy to configure for cross-origin requests

---

## 📦 Dependencies

### Core Dependencies
- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[JWT-Go](https://github.com/golang-jwt/jwt)** - JWT implementation
- **[SQLite3](https://github.com/mattn/go-sqlite3)** - Database driver
- **[bcrypt](https://golang.org/x/crypto/bcrypt)** - Password hashing
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable loading
- **[migrate](https://github.com/golang-migrate/migrate)** - Database migration tool

---

## 🚀 Deployment

### Docker (Optional)
```dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main api/main.go

EXPOSE 8080
CMD ["./main"]
```

### Build for Production
```bash
CGO_ENABLED=1 GOOS=linux go build -o server api/main.go
```

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Connect with the Developer

<div align="center">

### **Hifzaan Mohammad**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/hifzaan-mohammad)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/HifzaanDev)
[![X (Twitter)](https://img.shields.io/badge/X-000000?style=for-the-badge&logo=x&logoColor=white)](https://x.com/AslHif76539)

---

### 💡 *"Building robust APIs with Go, one endpoint at a time!"*

</div>

---

<div align="center">

**⭐ Star this repository if you found it helpful!**

Made with ❤️ and Go

</div>
