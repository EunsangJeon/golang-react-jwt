package db

// Queries to DB
const (
	CheckUserExists = `SELECT id from users WHERE email = $1`
	LoginQuery      = `SELECT * from users WHERE email = $1`
	CreateUserQuery = `INSERT INTO users(id,name,password,email) VALUES (DEFAULT, $1 , $2, $3);`
)
