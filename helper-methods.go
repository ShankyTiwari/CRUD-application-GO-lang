package main

import (
	"encoding/json"
	"fmt"
)

func getUsersFromDB() []byte {
	var (
		user  User
		users []User
	)
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Lname, &user.Country)
		users = append(users, user)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func insertUserInDB(userDetails User) bool {
	stmt, err := db.Prepare("INSERT into users SET Name=?,Lname=?,Country=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.Name, userDetails.Lname, userDetails.Country)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func deleteUserFromDB(userID string) bool {
	stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func updateUserInDB(userDetails User) bool {
	stmt, err := db.Prepare("UPDATE users SET name=?,lname=?,country=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.Name, userDetails.Lname, userDetails.Country, userDetails.ID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
