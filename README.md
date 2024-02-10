## Fullstack web application with NextJs and Go

### Description

This project is a web application built using Next.js for the frontend and Go for the backend. 

### Features

- User registration with email, password, name, and last name.
- User login with email and password.
- User management: create, read, update, and delete users.
- Secure password storage using bcrypt hashing.
- RESTful API endpoints for user management.

### Installation

#### Prerequisites

- Node.js
- Go
- PostgreSQL

#### Frontend Setup

1. Navigate to the `frontend` directory.
2. Run `npm install` to install dependencies.
3. Run `npm run dev` to start the development server.
4. Access the application at `http://localhost:3000` in your browser.

#### Backend Setup

1. Set up a PostgreSQL database and ensure it's running.
2. Navigate to the `backend` directory.
3. Update the database connection details if necessary.
4. Run `go run main.go` to start the Go server.
5. The server will start on port 8080 by default.

### Usage

#### Register

To register a new user, navigate to the registration page and fill out the required fields: email, password, name, and last name.

#### Login

Once registered, you can log in using your email and password.

#### User Management

After logging in, you can manage users:
- Create: Add a new user.
- Read: View user details.
- Update: Modify user information.
- Delete: Remove a user from the system.
