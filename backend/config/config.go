package config

import "os"

// DBUser is string from os environment variable.
// It represents DB username.
var DBUser = os.Getenv("DB_USERNAME")

// DBPassword is string from os environment variable.
// It represents DB password.
var DBPassword = os.Getenv("DB_PASSWORD")

// DBName is string from os environment variable.
// It represents DB name.
var DBName = os.Getenv("DB_NAME")

// ClientURL is string from os environment variable.
// It is frontend endpoint.
var ClientURL = os.Getenv("CLIENT_URL")

// JWTKey is key for JWT signiture.
var JWTKey = "secret"
