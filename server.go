package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/labstack/echo"
)
var db * sql.DB
var err error
var chttp = http.NewServeMux()

type User struct{
	Id int
    Name string
    Lname string
    Country string	
}

func main() {

	db,err = sql.Open("mysql","root:@/test")


	e := echo.New()

	chttp.Handle("/", http.FileServer(http.Dir("./")))


	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		if (strings.Contains(r.URL.Path,".")) {
			chttp.ServeHTTP(w,r)
		}else{
			http.ServeFile(w,r,"views/index.html")
		}
	})

	
	http.HandleFunc("/getUsers", func(response http.ResponseWriter,req *http.Request){

		var (
			user User
			users []User
			)
		
		rows, err:= db.Query("SELECT * FROM users")
		if err !=nil{
			panic(err)
		}

		for rows.Next(){
			
			err = rows.Scan(&user.Id, &user.Name, &user.Lname, &user.Country)
			users = append(users,user)

			if err !=nil{
				panic(err)
			}
		}
		defer rows.Close()

		jsonResponse , err := json.Marshal(users)
		if err !=nil{
			panic(err)
		}
		response.Header().Set("Content-Type","application/json")
		response.Write(jsonResponse)		
	})

	http.HandleFunc("/insertUsers",func(rw http.ResponseWriter,req *http.Request){
		
		decoder :=json.NewDecoder(req.Body)
		var t User
		err := decoder.Decode(&t)
		if err !=nil{
			panic(err)
		}
		defer req.Body.Close()

		stmt ,err := db.Prepare("INSERT into users SET id=?,name=?,lname=?,country=?")
		if err !=nil{
			panic(err)
		}


		res ,err := stmt.Exec("",t.Name,t.Lname,t.Country)
		if err !=nil{
			panic(err)
		}
		fmt.Println(res)
		
		var (
			user User
			users []User
			)
		
		rows, err:= db.Query("SELECT * FROM users")
		if err !=nil{
			panic(err)
		}

		for rows.Next(){
			
			err = rows.Scan(&user.Id, &user.Name, &user.Lname, &user.Country)
			users = append(users,user)

			if err !=nil{
				panic(err)
			}
		}
		defer rows.Close()

		jsonResponse , err := json.Marshal(users)
		if err !=nil{
			panic(err)
		}
		rw.Header().Set("Content-Type","application/json")
		rw.Write(jsonResponse)
	})
	
	e.Logger.Fatal(http.ListenAndServe(":1323",nil))
}
