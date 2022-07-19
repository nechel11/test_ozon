package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"github.com/nechel11/test_ozon/internal/base"
	"github.com/nechel11/test_ozon/internal/utils"
	"github.com/nechel11/test_ozon/internal/models"
)

func PG_handler(w http.ResponseWriter, r *http.Request){
	var db *sql.DB = data.Db_connect()
	if r.Method == "POST"{
		short_output_handle_PG(w, r, db)
	} else if r.Method == "GET"{
		long_output_handle_PG(w, r, db)
	} else{
		utils.If_error_response(w, errors.New("invalid method"), http.StatusBadRequest)
		return
	}
}

func short_output_handle_PG(w http.ResponseWriter, r *http.Request, db *sql.DB){
	
	var url_req models.JsonUrl
	var encoded_string models.JsonUrl
	var if_short_exists bool

	log.Println("Connected to db")
	if err := utils.Decoder_json(&url_req, r.Body); err != nil{	
		utils.If_error_response(w, err, http.StatusBadRequest)
		return
	}
	if err := database.Db_if_data_exists(&if_short_exists, url_req.Url, db); err != nil{
		utils.If_error_response(w, errors.New("db check for existence error"), http.StatusInternalServerError)
		return
	}
	if if_short_exists{
		var err_response error
		if encoded_string.Url, err_response = database.Db_get_short_url(url_req.Url, db); err_response != nil{
			utils.If_error_response(w, errors.New("getting short url from db error"), http.StatusBadRequest)
			return
		}
		log.Println("encoded string has been sent")
	} else {
			encoded_string.Url = utils.Hash_func(url_req.Url)
			if err := db_insert_url(encoded_string.Url, url_req.Url, db); err != nil{
				utils.If_error_response(w, errors.New("db adding data error"), http.StatusInternalServerError)
				return
			}
		log.Print(url_req.Url, " added to db with shortlink ", encoded_string, "\n")
	}
	utils.Send_response(w, encoded_string)	
}

func long_output_handle_PG(w http.ResponseWriter, r *http.Request, db *sql.DB){
	var url_req models.JsonUrl
	var decoded_string models.JsonUrl
	var err error

	if err_json := utils.Decoder_json(&url_req, r.Body); err_json != nil{	
		utils.If_error_response(w, err_json, http.StatusBadRequest)
	}
	decoded_string.Url, err = Db_get_long_url(url_req.Url, db)
	if err != nil{
		utils.If_error_response(w, errors.New("getting long url from db error"), http.StatusBadRequest)
	}
	utils.Send_response(w, decoded_string)
}