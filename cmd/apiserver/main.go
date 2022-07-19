package main

import (
	
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal"
	"os"
)


func handlefunc(flag bool){
	log.Println("entered main")
	if !flag{
		http.HandleFunc("/", internal.PG_handler)
	} else {
		http.HandleFunc("/", internal.Cache_handler)
	}
	http.ListenAndServe("localhost:8070", nil)
}

func main(){
	var flag bool
	if os.Args[1] == "Cache"{
		flag = true
	}
	log.Println("Server http://localhost:8070")
	handlefunc(flag)
	
}