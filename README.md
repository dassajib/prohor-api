# Prohor API

High-performance RESTful API backend for the Prohor platform, built with Go, Gin, and PostgreSQL.

---

## Table of Contents

- [About](#about)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Running Locally](#running-locally)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

---

## About

`prohor-api` is the backend service powering the Prohor client applications.  
It offers secure, scalable REST APIs for authentication, user management, and business logic.

Built with:

- [Gin](https://github.com/gin-gonic/gin) for lightweight, fast HTTP routing  
- [GORM](https://gorm.io) for ORM and database management  
- PostgreSQL as the primary database  
- JWT authentication for secure access control  
- CORS support for cross-origin requests  

---

## Tech Stack

| Component          | Technology / Library                  |
| ------------------ | ----------------------------------- |
| Language           | Go 1.23                            |
| HTTP Framework     | Gin (github.com/gin-gonic/gin)      |
| Database ORM       | GORM (gorm.io/gorm)                 |
| Database Driver    | PostgreSQL (gorm.io/driver/postgres)|
| Authentication     | JWT (github.com/golang-jwt/jwt/v5) |
| Validation         | go-playground/validator             |
| CORS Middleware    | gin-contrib/cors                    |
| Environment Config | joho/godotenv                      |

---

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL (v13+ recommended)
- Git

### Installation

1. Clone the repository

```bash
git clone https://github.com/dassajib/prohor-api.git
cd prohor-api

### Install dependencies
go mod tidy

### Setup environment variables
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_name
DB_PASSWORD=your_password
DB_NAME=dbname
ACCESS_SECRET=access
REFRESH_SECRET=refresh

### Prepare your database
createdb db_name

### Run the server in development mode
go run cmd/main.go

## Credits

This project was developed by **Das Sajib**.
Name : Sajib Das
github : github.com/dassajib

