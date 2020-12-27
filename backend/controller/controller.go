package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EunsangJeon/golang-react-jwt/backend/db"
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

// CreateResponse is type for create response
type createResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"msg"`
	Errors  []string `json:"errors"`
}

// Create new user
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user db.Register
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Could not json decode from register body.", http.StatusInternalServerError)
		return
	}

	valErr := util.ValidateUser(user)
	if len(valErr) > 0 {
		res := createResponse{Success: false, Message: "User creation failed", Errors: valErr}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(res)
		return
	}

	db.HashPassword(&user)
	_, err = db.DB.Query(db.CreateUserQuery, user.Name, user.Password, user.Email)
	if err != nil {
		http.Error(w, "Could not write new user to DB.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := createResponse{Success: true, Message: "User created successfully", Errors: []string{}}
	json.NewEncoder(w).Encode(res)
	return
}

// Session returns JSON of user info
func Session(c *gin.Context) {
	user, isAuthenticated := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "user": nil, "msg": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user, "msg": "authorized"})
}

// Login controller
func Login(c *gin.Context) {
	var user db.Login
	c.Bind(&user)

	row := db.DB.QueryRow(db.LoginQuery, user.Email)

	var id int
	var name, email, password, createdAt, updatedAt string

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	match := db.CheckPasswordHash(user.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(10 * time.Second)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{
		User: db.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, err := token.SignedString(jwtKey)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "logged in succesfully", "user": claims.User, "token": tokenString})
}
