package main

import (
	
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal"
)


func handlefunc(){
	http.HandleFunc("/short", internal.Short_url)
	http.HandleFunc("/long", internal.Long_url)
	http.ListenAndServe("localhost:8070", nil)
}

func main(){
	
	log.Println("Server http://localhost:8070")
	handlefunc()
	
}