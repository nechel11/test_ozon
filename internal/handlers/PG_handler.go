package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/database"
)

func PG_handler(w http.ResponseWriter, r *http.Request){
	var db *sql.DB = Db_connect()
	if r.Method == "POST"{
		short_output_handle_PG(w, r, db)
	} else if r.Method == "GET"{
		long_output_handle_PG(w, r, db)
	} else{
		if_error_response(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

func short_output_handle_PG(w http.ResponseWriter, r *http.Request, db *sql.DB){
	
	var url_req JsonUrl
	var encoded_string JsonUrl
	var if_short_exists bool

	log.Println("Connected to db")
	if err := decoder_json(&url_req, r.Body); err != nil{	
		if_error_response(w, err, http.StatusBadRequest)
		return
	}
	if err := db_if_data_exists(&if_short_exists, url_req.Url, db); err != nil{
		if_error_response(w, errors.New("db check for existence error"), http.StatusInternalServerError)
		return
	}
	if if_short_exists{
		var err_response error
		if encoded_string.Url, err_response = db_get_short_url(url_req.Url, db); err_response != nil{
			if_error_response(w, errors.New("getting short url from db error"), http.StatusBadRequest)
			return
		}
		log.Println("encoded string has been sent")
	} else {
			encoded_string.Url = hash_func(url_req.Url)
			if err := db_insert_url(encoded_string.Url, url_req.Url, db); err != nil{
				if_error_response(w, errors.New("db adding data error"), http.StatusInternalServerError)
				return
			}
		log.Print(url_req.Url, " added to db with shortlink ", encoded_string, "\n")
	}
	send_response(w, encoded_string)	
}

func long_output_handle_PG(w http.ResponseWriter, r *http.Request, db *sql.DB){
	var url_req JsonUrl
	var decoded_string JsonUrl
	var err error

	if err_json := decoder_json(&url_req, r.Body); err_json != nil{	
		if_error_response(w, err_json, http.StatusBadRequest)
	}
	decoded_string.Url, err = db_get_long_url(url_req.Url, db)
	if err != nil{
		if_error_response(w, errors.New("getting long url from db error"), http.StatusBadRequest)
	}
	send_response(w, decoded_string)
}