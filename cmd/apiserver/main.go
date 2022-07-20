package main

import (
	
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/handlers"
	"github.com/nechel11/test_ozon/internal/models"
	"flag"
)


func handlefunc(flag string){
	if flag == "pg"{
		log.Println("application works with postgres")
		http.HandleFunc("/", handlers.PG_handler)
	} else if flag == "cache" {
		log.Println("application works with internal memory")
		map_short_key := make(map[string]models.JsonUrl)
		map_long_key := make(map[string]models.JsonUrl)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		handlers.Cache_handler(w, r, map_short_key, map_long_key)
		})
	} else {
		log.Fatal("wrong storage")
		return
	}
	http.ListenAndServe("localhost:8070", nil)
}

func main(){
	var storageType string
	flag.StringVar(&storageType, "storage", "pg",
		"choose storage : pg for postgres, cache for internal memory")
	flag.Parse()
	log.Println("Server http://localhost:8070")
	handlefunc(storageType)
	
}