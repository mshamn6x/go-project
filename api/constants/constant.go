package constants

import "os"

const (
	DbName    = "postgres://postgres:root@localhost:5432/project1?sslmode=disable"
	JWTIssuer = "project1"
)

var Secret = os.Getenv("JWT_SECRET")
