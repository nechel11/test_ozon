package internal

import (
	"log"
	"net/http"
	"strconv"
)


func if_error_respose(w http.ResponseWriter, err error, number int){
	log.Println(err.Error(), number)
	http.Error(w, strconv.Itoa(number), number)
}