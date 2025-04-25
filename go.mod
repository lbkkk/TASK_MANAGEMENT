module task-app

go 1.23.4

require (
	github.com/go-chi/chi/v5 v5.2.1
	github.com/go-chi/cors v1.2.1
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/mattn/go-sqlite3 v1.14.24
	golang.org/x/oauth2 v0.29.0
)

require cloud.google.com/go/compute/metadata v0.3.0 // indirect
