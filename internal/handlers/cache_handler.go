package handlers

import (
	"errors"
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/models"
)


func Cache_handler(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl, map_long_key map[string]models.JsonUrl){
	log.Println("entered handker")

	if r.Method == "POST"{
		short_output_handle_Cache(w, r, map_short_key, map_long_key)
	} else if r.Method == "GET"{
		long_output_handle_Cache(w, r, map_short_key, map_long_key)
	} else{
		if_error_response(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

func short_output_handle_Cache(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl, map_long_key map[string]JsonUrl){
	var url_req models.JsonUrl
	var encoded_string models.JsonUrl
	
	if err := decoder_json(&url_req, r.Body); err != nil{	
		if_error_response(w, err, http.StatusBadRequest)
		return
	}
	encoded_string.Url = hash_func(url_req.Url)
	map_long_key[url_req.Url] = encoded_string
	map_short_key[encoded_string.Url] = url_req

	log.Println(map_long_key, map_short_key)
	send_response(w, encoded_string)
}

func long_output_handle_Cache(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl, map_long_key map[string]JsonUrl){
	var url_req JsonUrl
	
	if err_json := decoder_json(&url_req, r.Body); err_json != nil{	
		if_error_response(w, err_json, http.StatusBadRequest)
	}
	send_response(w, map_short_key[url_req.Url])
}