# 🔗 URL Shortener

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/MongoDB-4.4+-47A248?style=for-the-badge&logo=mongodb&logoColor=white" alt="MongoDB" />
  <img src="https://img.shields.io/badge/Redis-6.0+-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis" />
  <img src="https://img.shields.io/badge/React-18+-61DAFB?style=for-the-badge&logo=react&logoColor=black" alt="React" />
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge" alt="License: MIT" />
</div>

<p align="center">
  A modern, high-performance URL shortening service built with Go, MongoDB, Redis, and React.
</p>
<img width="1440" alt="image" src="https://github.com/user-attachments/assets/dd361fed-dcc6-4fde-a89c-d87e5fe7917a" />

## 🚀 Features

- **Fast URL Shortening**: Generate short, unique codes for long URLs in milliseconds
- **Redis Caching**: High-performance caching for frequently accessed URLs
- **Easy-to-Use API**: RESTful API for all URL operations
- **Modern Dashboard**: React-based frontend for user-friendly URL management
- **URL Management**: Create, read, update, and delete short URLs
- **Analytics**: Track how many times each short URL has been accessed
- **Secure**: Validation of URLs to prevent abuse
- **Customizable**: Configure base URL, port, and database settings
- **Responsive**: Works on desktop and mobile devices

## 📋 Table of Contents

- [Features](#-features)
- [Architecture](#-architecture)
- [Installation](#-installation)
- [Usage](#-usage)
- [API Reference](#-api-reference)
- [Frontend](#-frontend)
- [Configuration](#-configuration)
- [Development](#-development)
- [Contributing](#-contributing)
- [License](#-license)

## 🏗️ Architecture

This project follows a clean architecture pattern:

```
url-shortener-golang
├── config/        # Configuration handling
├── controllers/   # HTTP request handlers
├── models/        # Data models
├── repositories/  # Database access layer
├── services/      # Business logic layer
├── utils/         # Utility functions
└── frontend/      # React frontend
```

The system works in layers:
1. **Controller Layer**: Handles HTTP requests and responses
2. **Service Layer**: Implements business logic
3. **Repository Layer**: Manages database interactions
4. **Model Layer**: Defines data structures

### Caching Architecture

The application uses Redis for high-performance caching:

- URL lookups are first attempted from the Redis cache
- Cache misses fall back to MongoDB database queries
- Successful database queries populate the cache for future requests
- Cache entries have a configurable TTL (Time To Live)
- URL updates/deletes automatically invalidate cache entries

## 📦 Installation

### Prerequisites

- Go 1.16+
- MongoDB 4.4+
- Redis 6.0+ (optional, but recommended for production)
- Node.js 14+
- npm or yarn

### Setup

1. **Clone the repository**

```bash
git clone https://github.com/askarbtw/url-shortener-golang.git
cd url-shortener-golang
```

2. **Install backend dependencies**

```bash
go mod download
```

3. **Install frontend dependencies**

```bash
cd frontend
npm install
cd ..
```

4. **Configure the application**

Edit the `.env` file in the root directory:

```
MONGO_URI=mongodb://localhost:27017
DB_NAME=URL_shortener
PORT=8080
BASE_URL=http://localhost:8080/
REDIS_URI=localhost:6379
REDIS_PASSWORD=
CACHE_TTL=3600
```

5. **Start MongoDB and Redis**

Ensure your MongoDB and Redis instances are running.

## 🔧 Usage

### Running the application

#### Option 1: Run Backend and Frontend Separately

1. Start the backend:

```bash
go run main.go
```

2. Start the frontend (in a separate terminal):

```bash
cd frontend
npm run dev
```

#### Option 2: Run Both with a Single Command

```bash
./run-app.sh
```

The backend API will be available at `http://localhost:8080` and the frontend at `http://localhost:5174`.

### Creating a short URL

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/some/long/path"}'
```

### Using a short URL

Simply visit `http://localhost:8080/r/YOUR_SHORT_CODE` in your browser, or:

```bash
curl -i http://localhost:8080/r/YOUR_SHORT_CODE
```

## 📚 API Reference

### Create Short URL

```
POST /shorten
```

**Request Body:**
```json
{
  "url": "https://www.example.com/some/long/url"
}
```

**Response:**
```json
{
  "id": "5f50c31a4f3c2a1d1c9c0c1d",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2023-03-20T12:00:00Z",
  "updatedAt": "2023-03-20T12:00:00Z"
}
```

### Retrieve Original URL

```
GET /shorten/{shortCode}
```

**Response:**
```json
{
  "id": "5f50c31a4f3c2a1d1c9c0c1d",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2023-03-20T12:00:00Z",
  "updatedAt": "2023-03-20T12:00:00Z"
}
```

### Update Short URL

```
PUT /shorten/{shortCode}
```

**Request Body:**
```json
{
  "url": "https://www.example.com/some/updated/url"
}
```

**Response:**
```json
{
  "id": "5f50c31a4f3c2a1d1c9c0c1d",
  "url": "https://www.example.com/some/updated/url",
  "shortCode": "abc123",
  "createdAt": "2023-03-20T12:00:00Z",
  "updatedAt": "2023-03-20T12:30:00Z"
}
```

### Delete Short URL

```
DELETE /shorten/{shortCode}
```

**Response:** 204 No Content

### Get URL Statistics

```
GET /shorten/{shortCode}/stats
```

**Response:**
```json
{
  "id": "5f50c31a4f3c2a1d1c9c0c1d",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2023-03-20T12:00:00Z",
  "updatedAt": "2023-03-20T12:00:00Z",
  "accessCount": 10
}
```

### Redirect

```
GET /r/{shortCode}
```

Redirects to the original URL.

## 🖥️ Frontend

The project includes a modern React frontend with:

- Dashboard to view all shortened URLs
- Form to create new short URLs
- Details view for URL statistics
- Update and delete functionality
- Responsive design

## ⚙️ Configuration

The application can be configured using environment variables in the `.env` file:

| Variable        | Description                   | Default                  |
|-----------------|-------------------------------|--------------------------|
| MONGO_URI       | MongoDB connection URI        | mongodb://localhost:27017|
| DB_NAME         | MongoDB database name         | URL_shortener            |
| PORT            | Port for the backend API      | 8080                     |
| BASE_URL        | Base URL for short links      | http://localhost:8080/   |
| REDIS_URI       | Redis connection URI          | localhost:6379           |
| REDIS_PASSWORD  | Redis password (if required)  | (empty)                  |
| CACHE_TTL       | Cache time to live in seconds | 3600 (1 hour)            |

## 🛠️ Development

### Project Structure

```
├── config/
│   ├── config.go          # Configuration handling
│   ├── db.go              # MongoDB connection
│   └── redis.go           # Redis connection
├── controllers/
│   └── url_controller.go  # HTTP handlers for URL operations
├── models/
│   ├── errors.go          # Custom error definitions
│   └── url.go             # URL data model
├── repositories/
│   └── url_repository.go  # MongoDB data access layer
├── services/
│   ├── cache_service.go   # Redis caching service
│   └── url_service.go     # Business logic for URL operations
├── utils/
│   └── shortcode.go       # Short code generation utilities
├── frontend/
│   ├── src/               # React frontend code
│   ├── public/            # Static assets
│   └── package.json       # Frontend dependencies
├── main.go                # Application entry point
├── go.mod                 # Go dependencies
├── go.sum                 # Go dependencies checksums
├── .env                   # Environment variables
└── README.md              # Project documentation
```

### Adding New Features

To add new features:

1. Create or modify models in `models/`
2. Update the repository functions in `repositories/`
3. Implement business logic in `services/`
4. Create API endpoints in `controllers/`
5. Register routes in `main.go`

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some amazing feature'`)
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

Please ensure your code follows the project's style guidelines.

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/askarbtw">askarbtw</a>
</p> 
