package utils

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"errors"
	"github.com/nechel11/test_ozon/internal/models"
)

// error handler

func If_error_response(w http.ResponseWriter, err error, number int){
	log.Println(err.Error(), number)
	http.Error(w, strconv.Itoa(number) + " " + err.Error(), number)
}

// send response to client fucntion

func Send_response(w http.ResponseWriter, response models.JsonUrl){
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		If_error_response(w, errors.New("encoding error"), http.StatusInternalServerError)
		return
	}
}
