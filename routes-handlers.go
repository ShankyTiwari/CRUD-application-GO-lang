package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

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

func renderHome(response http.ResponseWriter, request *http.Request) {
	if strings.Contains(request.URL.Path, ".") {
		chttp.ServeHTTP(response, request)
	} else {
		http.ServeFile(response, request, "views/index.html")
	}
}

func getUsers(response http.ResponseWriter, request *http.Request) {

	var (
		user  User
		users []User
	)
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.Name, &user.Lname, &user.Country)
			users = append(users, user)
		}
		defer rows.Close()

		jsonResponse, err := json.Marshal(users)
		if err != nil {
			returnErrorResponse(response, request, httpError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonResponse)
		}

	}
}

func inertUsers(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var userDetails User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()

	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		stmt, err := db.Prepare("INSERT into users SET id=?,name=?,lname=?,country=?")
		if err != nil {
			returnErrorResponse(response, request, httpError)
		} else {
			_, err := stmt.Exec("", userDetails.Name, userDetails.Lname, userDetails.Country)
			if err != nil {
				returnErrorResponse(response, request, httpError)
			} else {
				getUsers(response, request)
			}
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
