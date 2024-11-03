# go-jwt-mk1-showcase

A Go-based showcase project demonstrating JWT authentication using [Fiber](https://gofiber.io) and [gofiber/jwt](https://github.com/golang-jwt/jwt) for handling secure routes. This project includes user authentication, token-based access control, and protected routes.

## Project Structure

```
go-jwt-mk1-showcase/
├── gojwt/
│   └── gojwt.go
├── server/
│   ├── handlers/
│   │   └── handlers.go
│   ├── middleware/
│   │   └── middleware.go
│   └── templates/
│       └── index.html
├── go.mod
├── go.sum
└── main.go
```

## Project Overview

This application includes three main parts:

- **Login Page:** A form where users enter their username and password.
- **JWT Authentication:** Users receive a JWT upon successful login, stored as a secure HTTP-only cookie.
- **Protected Route:** A secure page accessible only to authenticated users with a valid JWT.

**Key Components**

1. JWT Management **(gojwt.go)**

The gojwt package handles token creation and verification.

- **ConfigJWT:** Generates a JWT token with an expiration time and a subject identifier based on the username.
- **VerifyJWT:** Middleware function that validates incoming JWTs against the secret key.

2. Middleware **(middleware.go)**

The Auth middleware in middleware/middleware.go checks for a valid JWT in the user’s cookies. If the token is invalid, the user is redirected to the login page.

3. Handlers **(handlers.go)**

Defines route handlers for user interactions:

- **Root:** Renders the login page (index.html).
- **Submit:** Handles login form submissions, authenticates users against a mock database (Users map), and issues a JWT if credentials are correct.
- **SecretPage:** A protected route that displays the JWT token if the user is authenticated.

**Routes**

```
GET /: Serves the login page.
POST /submit: Authenticates users; issues a JWT on success.
GET /secret-page: A protected route that only accessible with a valid JWT.
```

**Error Handling**

The custom Error handler in handlers/handlers.go manages 404 and 500 errors:

- Displays a 404 message if a page is not found.
- Shows a generic message for any server-side issues.

## Requirements

- **Go** (v1.16 or later)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-jwt-mk1-showcase
```

```bash
cd go-jwt-mk1-showcase
```

### 2. Run the Application

Start the server with:

```bash
go run main.go
```

The application will start on **http://0.0.0.0:4032** and will be accessible on any device in your WI-FI network.

## License

[MIT](https://choosealicense.com/licenses/mit)
