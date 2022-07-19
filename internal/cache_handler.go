package internal

import (
	"errors"
	"log"
	"net/http"
)

func Cache_handler(w http.ResponseWriter, r *http.Request){
	log.Println("entered handker")
	map_short := make(map[string]JsonUrl)
	map_long := make(map[string]JsonUrl)
	if r.Method == "POST"{
		short_url_handle_Cache(w, r, map_short, map_long)
	} else if r.Method == "GET"{
		long_url_handle_Cache(w, r, map_short, map_long)
	} else{
		if_error_response(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

func short_url_handle_Cache(w http.ResponseWriter, r *http.Request, map_short map[string]JsonUrl, map_long map[string]JsonUrl){
	var url_req JsonUrl
	var encoded_string JsonUrl
	
	if err := decoder_json(&url_req, r.Body); err != nil{	
		if_error_response(w, err, http.StatusBadRequest)
		return
	}
	encoded_string.Url = hash_func(url_req.Url)
	map_long[url_req.Url] = encoded_string
	map_short[encoded_string.Url] = url_req
	log.Println(map_long, map_short)
}

func long_url_handle_Cache(w http.ResponseWriter, r *http.Request, map_short map[string]JsonUrl, map_long map[string]JsonUrl){
	var url_req JsonUrl
	

	if err_json := decoder_json(&url_req, r.Body); err_json != nil{	
		if_error_response(w, err_json, http.StatusBadRequest)
	}

	log.Println(url_req.Url, map_short, map_long)
}