package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/joho/godotenv"
    "os"
)

var jwtKey []byte

// Load the JWT secret key from the .env file
func init() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Retrieve the JWT secret key
    jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

func main() {
    // setupCORS is a middleware to enable CORS
    http.HandleFunc("/", setupCORS(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the server %s\n", r.URL.Path[1:])
    }))
    
    // Generate a new JWT token
    http.HandleFunc("/generateJWT", generateJWT)
    
    // APIs that require authentication
    http.HandleFunc("/departments", authenticate(departments))
    http.HandleFunc("/patients", authenticate(patients))
    http.HandleFunc("/patientCentric", authenticate(patientCentric))

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", nil))
}