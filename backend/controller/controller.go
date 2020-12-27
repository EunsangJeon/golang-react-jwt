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
)

var jwtKey = []byte("secret")

type userClaims struct {
	db.User
	jwt.StandardClaims
}

type createResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"msg"`
	Errors  []string `json:"errors"`
}

type loginResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"msg"`
	User    db.User `json:"user"`
	Token   string  `json:"token"`
}

type sessionResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"msg"`
	User    db.User `json:"user"`
}

// Create new user
func Create(w http.ResponseWriter, r *http.Request) {
	user := db.Register{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Could not json decode from register body.", http.StatusInternalServerError)
		return
	}

	valErr := util.ValidateUser(user)
	if len(valErr) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = json.NewEncoder(w).Encode(createResponse{Success: false, Message: "User creation failed", Errors: valErr})
		if err != nil {
			http.Error(w, "Could not json encode to response", http.StatusInternalServerError)
			return
		}
		return
	}

	db.HashPassword(&user)
	_, err = db.DB.Query(db.CreateUserQuery, user.Name, user.Password, user.Email)
	if err != nil {
		http.Error(w, "Could not write new user to DB.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(createResponse{Success: true, Message: "User created successfully", Errors: []string{}})
	if err != nil {
		http.Error(w, "Could not json encode to response", http.StatusInternalServerError)
		return
	}
	return
}

// Session returns JSON of user info
func Session(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		c = &http.Cookie{}
	}

	ss := c.Value
	token, err := jwt.ParseWithClaims(
		ss,
		&userClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Cookie had not been signed with SHA256")
			}

			return jwtKey, nil
		},
	)
	isValid := err == nil && token.Valid

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(sessionResponse{Success: false, User: db.User{}, Message: "Unauthenticated"})
		if err != nil {
			http.Error(w, "Error while json encoding", http.StatusInternalServerError)
			return
		}
		return
	}

	claims := token.Claims.(*userClaims)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(sessionResponse{Success: true, User: claims.User, Message: "Authenticated"})
	if err != nil {
		http.Error(w, "Error while json encoding", http.StatusInternalServerError)
		return
	}
	return
}

// Login controller
func Login(w http.ResponseWriter, r *http.Request) {
	user := db.Login{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Could not json decode ogin request body.", http.StatusInternalServerError)
		return
	}

	row := db.DB.QueryRow(db.LoginQuery, user.Email)

	var id int
	var name, email, password, createdAt, updatedAt string

	errorNoRows := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)
	match := db.CheckPasswordHash(user.Password, password)

	if errorNoRows == sql.ErrNoRows || !match {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(loginResponse{Success: false, Message: "incorrect credentials", User: db.User{}, Token: ""})
		if err != nil {
			http.Error(w, "Error while json encoding", http.StatusInternalServerError)
		}
		return
	}

	expirationTime := time.Now().Add(10 * time.Second)
	claims := &userClaims{
		User: db.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not getJWT", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(loginResponse{Success: true, Message: "logged in succesfully", User: claims.User, Token: signedString})
	if err != nil {
		http.Error(w, "Error while json encoding", http.StatusInternalServerError)
		return
	}
	return
}
