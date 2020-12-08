package controller

import (
	"fmt"
	"net/http"

	"github.com/EunsangJeon/golang-react-jwt/backend/config"
	"github.com/EunsangJeon/golang-react-jwt/backend/db"
	"github.com/EunsangJeon/golang-react-jwt/backend/errors"
	"github.com/EunsangJeon/golang-react-jwt/backend/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")

// Claims jwt claims struct
type Claims struct {
	db.User
	jwt.StandardClaims
}

// Pong tests that API is working
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

//Returns -1 as ID if the user doesnt exist in the table
func checkAndRetrieveUserIDViaEmail(createReset db.CreateReset) (int, bool) {
	rows, err := db.DB.Query(db.CheckUserExists, createReset.Email)
	if err != nil {
		return -1, false
	}
	if !rows.Next() {
		return -1, false
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return -1, false
	}
	return id, true
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

//InitiatePasswordReset initiates password reset email with reset url
func InitiatePasswordReset(c *gin.Context) {
	var createReset db.CreateReset
	c.Bind(&createReset)
	if id, ok := checkAndRetrieveUserIDViaEmail(createReset); ok {
		link := fmt.Sprintf("%s/reset/%d", config.ClientURL, id)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Successfully sent reset mail to " + createReset.Email, "link": link})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "errors": "No user found for email: " + createReset.Email})
	}
}

// ResetPassword resets password
func ResetPassword(c *gin.Context) {
	var resetPassword db.ResetPassword
	c.Bind(&resetPassword)
	if ok, errStr := util.ValidatePasswordReset(resetPassword); ok {
		password := db.CreateHashedPassword(resetPassword.Password)
		_, err := db.DB.Query(db.UpdateUserPasswordQuery, resetPassword.ID, password)
		errors.HandleErr(c, err)
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User password reset successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "errors": errStr})
	}
}
