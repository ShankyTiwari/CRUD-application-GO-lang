package main

// User is Interface for user details.
type User struct {
	ID      int
	Name    string
	Lname   string
	Country string
}

// ErrorResponse is interface for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}
