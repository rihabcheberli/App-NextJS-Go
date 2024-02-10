package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rihabcheberli/App-NextJS-Go/backend/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewDB() (*sql.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=admin dbname=Users sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	return db, nil
}

func executeQuery(query string, args ...interface{}) error {
	db, err := NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query, args...)
	return err
}

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	query := "SELECT id, email FROM users WHERE id = $1"
	err := executeQuery(query, &user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, password, name, last_name) VALUES ($1, $2, $3, $4)"
	return executeQuery(query, user.Email, user.Password, user.Name, user.LastName)
}

func UpdateUserByID(userID string, updatedUser models.User) error {
	query := "UPDATE users SET email = $1, name = $2, last_name = $3 WHERE id = $4"
	return executeQuery(query, updatedUser.Email, updatedUser.Name, updatedUser.LastName, userID)
}

func DeleteUserByID(userID string) error {
	query := "DELETE FROM users WHERE id = $1"
	return executeQuery(query, userID)
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err = db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]*models.User, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []*models.User
	query := "SELECT id, email, name, last_name FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.LastName)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %v", err)
	}

	return users, nil
}
