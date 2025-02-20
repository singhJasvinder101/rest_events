# Event Management API 🎉

![Go](https://img.shields.io/badge/Go-1.20-blue?style=flat-square)
![Gin](https://img.shields.io/badge/Gin-Framework-orange?style=flat-square)
![SQLite](https://img.shields.io/badge/SQLite-Database-lightgrey?style=flat-square)
![JWT](https://img.shields.io/badge/JWT-Authentication-green?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)

## 🌟 Overview
Event Management API is a **secure** and **scalable** backend for creating, updating, and managing events. Built with **Go (Gin framework)**, this API includes **JWT authentication**, role-based access, and database persistence using **SQLite**.

## 🚀 Features
✅ User authentication (Signup & Login) using **JWT** 🔑
✅ Create, Read, Update & Delete (CRUD) **events** 📅
✅ **Middleware for authentication & authorization** 🔐
✅ **Efficient database handling** with SQLite 🗄️
✅ **RESTful API design** with structured responses 🌐



## 📂 Project Structure
```
📦 event-management-api
 ┣ 📂 models         # Database models
 ┣ 📂 routes         # API route handlers
 ┣ 📂 middleware     # JWT Authentication middleware
 ┣ 📂 utils          # Utility functions
 ┣ 📂 db             # Database initialization
 ┣ 📜 main.go        # Entry point
 ┗ 📜 README.md      # Project documentation
```



## 🛠️ Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/singhJasvinder101/event-management-api.git
cd event-management-api
```

### 2️⃣ Install Dependencies
Ensure you have **Go** installed on your machine.
```sh
go mod tidy
```

### 3️⃣ Setup & Run the Server
```sh
go run main.go
```

The server runs on **http://localhost:8080** 🚀



## 🔑 Authentication (JWT)
This API uses **JWT** for secure authentication. Include your token in the `Authorization` header for protected routes.
```sh
Authorization: Bearer <your_token>
```



## 🔄 API Endpoints
### 🧑‍💻 User Routes
| Method | Endpoint      | Description         |
|--|-||
| POST   | `/signup`    | Register a new user |
| POST   | `/login`     | Authenticate & get token |

### 🎟️ Event Routes (Protected)
| Method | Endpoint         | Description              |
|--|--|--|
| GET    | `/events`        | Fetch all events        |
| POST   | `/events`        | Create a new event      |
| GET    | `/events/:id`    | Get event by ID         |
| PUT    | `/events/:id`    | **Update event (owner only)** |
| DELETE | `/events/:id`    | **Delete event (owner only)** |



## 🔧 Middleware
Middleware ensures security & role-based access control:
- **AuthMiddleware**: Protects routes, verifies JWT tokens.
- **OwnerMiddleware**: Ensures only event creators can modify or delete their events.




## 🤝 Contribution Guidelines
We welcome contributions! Follow these steps:
1. **Fork** the repository 🍴
2. **Create a new branch** (`feature-new`)
3. **Commit changes** (`git commit -m "Added new feature"`)
4. **Push** to your forked repository
5. **Submit a Pull Request** 🚀




### 🎉 Happy Coding! 🚀
