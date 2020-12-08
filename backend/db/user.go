package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Register struc
type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ResetPassword struct
type ResetPassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Login struct
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// CreateReset struct
type CreateReset struct {
	Email string `json:"email"`
}

//User struct
type User struct {
	//ID string
	Password  string `json:"password"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// HashPassword hashes password
func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

// CreateHashedPassword generates hashed password with the one given
func CreateHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

//CheckPasswordHash compares hash with password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
