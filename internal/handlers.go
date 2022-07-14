package internal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var Db *sql.DB


func send_response(w http.ResponseWriter, short string){
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(short)
	if err != nil {
		if_error_respose(w, errors.New("encoding error"), http.StatusInternalServerError)
		return
	}
	
}

func Short_url(w http.ResponseWriter, r *http.Request){
	Db = Db_connect()
	log.Println("Connected to db")
	defer Db.Close()

	if r.Method != "POST"{
		if_error_respose(w, errors.New("only POST method"), http.StatusBadRequest)
	}
	var url_req JsonUrl
	if err := decoder_json(&url_req, r.Body); err != nil{
		if_error_respose(w, err, http.StatusBadRequest)
	}
	encoded_string := hash_func(url_req.Url)
	if err := insert_url(encoded_string, url_req.Url, Db); err != nil{
		if_error_respose(w, errors.New("data base adding error"), http.StatusInternalServerError)
		return
	}
	log.Print(url_req.Url, " added to db with shortlink ", encoded_string, "\n")
	send_response(w, encoded_string)
}

func Long_url(w http.ResponseWriter, r *http.Request){

}