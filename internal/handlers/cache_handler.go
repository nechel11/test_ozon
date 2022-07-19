package handlers

import (
	"errors"
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/models"
	"github.com/nechel11/test_ozon/internal/utils"
)


func Cache_handler(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl, map_long_key map[string]models.JsonUrl){
	if r.Method == "POST"{
		short_output_handle_Cache(w, r, map_short_key, map_long_key)
	} else if r.Method == "GET"{
		long_output_handle_Cache(w, r, map_short_key)
	} else{
		utils.If_error_response(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

func short_output_handle_Cache(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl, map_long_key map[string]models.JsonUrl){
	var url_req models.JsonUrl
	var encoded_string models.JsonUrl
	
	if err := utils.Decoder_json(&url_req, r.Body); err != nil{	
		utils.If_error_response(w, err, http.StatusBadRequest)
		return
	}
	encoded_string.Url = utils.Hash_func(url_req.Url)
	map_long_key[url_req.Url] = encoded_string
	map_short_key[encoded_string.Url] = url_req

	log.Println(map_long_key, map_short_key)
	utils.Send_response(w, encoded_string)
}

func long_output_handle_Cache(w http.ResponseWriter, r *http.Request, map_short_key map[string]models.JsonUrl){
	var url_req models.JsonUrl
	
	if err_json := utils.Decoder_json(&url_req, r.Body); err_json != nil{	
		utils.If_error_response(w, err_json, http.StatusBadRequest)
	}
	utils.Send_response(w, map_short_key[url_req.Url])
}