package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func renderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	jsonResponse := getUsersFromDB()

	if jsonResponse == nil {
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func insertUser(response http.ResponseWriter, request *http.Request) {
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
		if userDetails.Name == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Lname == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Country == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "Country can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isInserted := insertUserInDB(userDetails)
			if isInserted {
				getUsers(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	userID := mux.Vars(request)["id"]
	if userID == "" {
		httpError.Message = "User id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		isdeleted := deleteUserFromDB(userID)
		if isdeleted {
			getUsers(response, request)
		} else {
			returnErrorResponse(response, request, httpError)
		}
	}
}

func updateUser(response http.ResponseWriter, request *http.Request) {
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
		if userDetails.Name == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.ID == 0 {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "user Id can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Lname == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Country == "" {
			httpError.Code = http.StatusBadRequest
			httpError.Message = "Country can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isUpdated := updateUserInDB(userDetails)
			if isUpdated {
				getUsers(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
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
