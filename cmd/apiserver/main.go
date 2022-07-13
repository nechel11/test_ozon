package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal"
)

var db *sql.DB

func handlefunc(){
	http.HandleFunc("/short", Short_url)
	http.HandleFunc("/short", Long_url)
	http.ListenAndServe("localhost:8070", nil)
}

func main(){
	db = Db_connect()
	log.Println("Connected to db")
	defer db.Close()
	handlefunc()
}