package internal

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)



func send_response(w http.ResponseWriter, response interface{}){
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		if_error_respose(w, errors.New("encoding error"), http.StatusInternalServerError)
		return
	}
}

func Short_url(w http.ResponseWriter, r *http.Request){
	var if_short_exists bool
	var url_req JsonUrl
	var db *sql.DB = Db_connect()
	var encoded_string JsonUrl

	log.Println("Connected to db")
	defer db.Close()	
	if r.Method != "POST"{
		if_error_respose(w, errors.New("only POST method"), http.StatusBadRequest)
		return
	}
	if err := decoder_json(&url_req, r.Body); err != nil{	
		if_error_respose(w, err, http.StatusBadRequest)
	}
	if err_db := if_data_exists(&if_short_exists, url_req.Url, db);err_db != nil{
		if_error_respose(w, errors.New("db check for existence error"), http.StatusInternalServerError)
		return
	}
	if if_short_exists{
		var err_response error
		if encoded_string.Url, err_response = get_short_url(url_req.Url, db); err_response != nil{
			if_error_respose(w, errors.New("getting short url from db error"), http.StatusBadRequest)
		}
		log.Println("encoded string has been sent")
		send_response(w, encoded_string)	
	} else {
			encoded_string.Url = hash_func(url_req.Url)
		if err := insert_url(encoded_string.Url, url_req.Url, db); err != nil{
			if_error_respose(w, errors.New("db adding data error"), http.StatusInternalServerError)
			return
		}
		log.Print(url_req.Url, " added to db with shortlink ", encoded_string, "\n")
		
		send_response(w, encoded_string)
	}
}

func Long_url(w http.ResponseWriter, r *http.Request){
	var url_req JsonUrl
	var decoded_string JsonUrl
	var db *sql.DB = Db_connect()
	var err error

	defer db.Close()
	if r.Method != "GET"{
		if_error_respose(w, errors.New("only GET method"), http.StatusBadRequest)
		return
	}
	if err_json := decoder_json(&url_req, r.Body); err_json != nil{	
		if_error_respose(w, err_json, http.StatusBadRequest)
	}
	decoded_string.Url, err = get_long_url(url_req.Url, db)
	if err != nil{
		if_error_respose(w, errors.New("getting long url from db error"), http.StatusBadRequest)
	}
	send_response(w, decoded_string)
}