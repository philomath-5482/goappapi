package main

import (
	"net/http"
	"os"
	"html/template"
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type ContactDetails struct{
	FirstName string
	LastName string
	Age string
}
var db *sql.DB

func home(w http.ResponseWriter, r *http.Request){

	t,err:=template.ParseFiles("signup.html")
	if err!=nil{
		fmt.Println(err)
	}
	if r.Method != http.MethodPost{
		t.Execute(w,nil)
		return
	
	}
}
func insertIntoDB(data ContactDetails) int64{
	result, err := db.Exec("INSERT INTO userdata VALUES (?, ?, ?)", data.FirstName, data.LastName, data.Age)
    if err != nil {
      fmt.Errorf("userdata: %v", err)
    } 
	rows, err := result.RowsAffected()
	if(err!=nil) {
		return 0
	}
	return rows
}
func signup(w http.ResponseWriter, r *http.Request){
	data:=ContactDetails{
		FirstName: r.FormValue("FirstName"),
		LastName: r.FormValue("LastName"),
		Age: r.FormValue("Age"),
	}
	fmt.Println(data)
	rows := insertIntoDB(data)
	if(rows>0){
		w.Write([]byte("rows inserted successfully"))
	}else {
		w.Write([]byte("rows cannot be inserted"))
	}
}

func main(){
	
	var err error
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", DB_USER, DB_PASSWORD,DB_NAME))
    if err != nil {
	    log.Printf("Error %s when opening DB\n", err)
	    return
	}
	// confirming the db connection
	err=db.Ping()
	if err==nil{
		fmt.Println("connected")
		fmt.Println(db)
	}else{
		fmt.Println(err)
	}

	http.HandleFunc("/",home)
	http.HandleFunc("/signup",signup)
	http.ListenAndServe(":8181",nil)
	//defer db.Close()
}