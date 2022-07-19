package main

import (
	
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/handlers"
	"github.com/nechel11/test_ozon/internal/models"
	"os"
)


func handlefunc(flag bool){
	if !flag{
		http.HandleFunc("/", handlers.PG_handler)
	} else {
		map_short_key := make(map[string]models.JsonUrl)
		map_long_key := make(map[string]models.JsonUrl)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		handlers.Cache_handler(w, r, map_short_key, map_long_key)
		})
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