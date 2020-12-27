package util

import (
	"regexp"

	"github.com/EunsangJeon/golang-react-jwt/backend/db"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

// ValidateUser returns a slice of string of validation errors
func ValidateUser(user db.Register) []string {
	err := []string{}

	if checkUserExists(user) {
		err = append(err, "Email already exists.")
		return err
	}

	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)

	if emailCheck != true {
		err = append(err, "Invalid email.")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters.")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name.")
	}

	return err
}

// Checks if user exists
func checkUserExists(user db.Register) bool {
	rows, err := db.DB.Query(db.CheckUserExists, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}
