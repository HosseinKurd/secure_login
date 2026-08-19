# Just a simple Secure Login Portal with GoLang

A simple and secure login portal built with Go language using **in-memory storage** instead of a database.

> **Note:** Users and sessions are lost when the application restarts.

## Features

* User registration
* User login
* User logout
* Protected `/dashboard`
* In-memory user storage
* In-memory session storage
* bcrypt password hashing
* Secure random session IDs
* HTTP-only session cookies
* Authentication middleware
* JSON API

## Project Structure

```text
secure-login/
├── go.mod
├── main.go
│
├── handlers/
│   ├── auth.go
│   └── dashboard.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   └── user.go
│
├── storage/
│   ├── users.go
│   └── sessions.go
│
├── security/
│   ├── password.go
│   └── session.go
│
└── router/
    └── router.go
```

## Requirements

* Go 1.24+
* No database required

## Dependencies

```go
require (
    github.com/google/uuid v1.6.0
    golang.org/x/crypto v0.41.0
)
```

Install dependencies:

```bash
go mod tidy
```

## Run the Application

```bash
go run .
```

The server starts on:

```text
http://localhost:8181
```

---

# API

## 1. Register

Create a new user.

```bash
curl -X POST http://localhost:8181/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}'
```

Expected response:

```json
{
  "message": "user registered"
}
```

---

## 2. Login

Authenticate a user and save the session cookie.

```bash
curl -i -c cookies.txt -X POST http://localhost:8181/login \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}'
```

The `-c cookies.txt` option stores the session cookie locally.

Expected response:

```json
{
  "message": "login successful"
}
```

---

## 3. Access Dashboard

Access the protected dashboard using the saved session.

```bash
curl -b cookies.txt http://localhost:8181/dashboard
```

Expected response:

```json
{
  "message": "welcome to dashboard",
  "user_id": "...",
  "username": "john"
}
```

Without authentication:

```bash
curl http://localhost:8181/dashboard
```

Expected response:

```text
unauthorized
```

---

## 4. Logout

Destroy the current session.

```bash
curl -b cookies.txt -X POST http://localhost:8181/logout
```

Expected response:

```json
{
  "message": "logout successful"
}
```

After logout, the dashboard should no longer be accessible:

```bash
curl -b cookies.txt http://localhost:8181/dashboard
```

Expected response:

```text
unauthorized
```

---

# Complete Test Flow

You can test the complete authentication flow in this order:

### Register

```bash
curl -X POST http://localhost:8181/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}'
```

### Login

```bash
curl -i -c cookies.txt -X POST http://localhost:8181/login \
  -H "Content-Type: application/json" \
  -d '{"username":"john","password":"secret123"}'
```

### Access Dashboard

```bash
curl -b cookies.txt http://localhost:8181/dashboard
```

### Logout

```bash
curl -b cookies.txt -X POST http://localhost:8181/logout
```

### Verify Logout

```bash
curl -b cookies.txt http://localhost:8181/dashboard
```

---

# Security Notes

Passwords are never stored as plaintext. They are hashed using bcrypt.

Session IDs are generated using `crypto/rand`.

Authentication is handled through an HTTP-only session cookie.

The current implementation uses in-memory storage:

```text
Application
    │
    ├── Users
    │     └── []User
    │
    └── Sessions
          └── map[sessionID]userID
```

For production, the in-memory storage should be replaced with persistent storage such as PostgreSQL or another appropriate database.

## Development Warning

The current cookie configuration uses:

```go
Secure: true
```

This requires HTTPS for browsers to send the cookie. For local development over plain HTTP, use HTTPS locally or configure the cookie appropriately for development.

Do **not** simply disable security settings in production.

---

# License

This is an open-source project for educational purposes.
