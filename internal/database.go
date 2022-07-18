package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const  (
	db_host = "127.0.0.1" 
	db_port=5432
	db_name="ozon"
	db_user = "zafar"
	db_password="12344321"
)

func Db_connect() *sql.DB{
	psqconn := fmt.Sprintf("host= %s port = %d user = %s password = %s dbname = %s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	db, err := sql.Open("postgres", psqconn)
	if err != nil {
		log.Fatal("Can not access to DB", err)
	}
	return db
}


func insert_url_db(short_url, long_url string, db *sql.DB) error{
	_, err := db.Exec(fmt.Sprintf("INSERT INTO records (long_url, short_url) VALUES ('%s', '%s')", long_url, short_url))
	if err != nil {		
		return err
	}
	return nil
}

func if_data_exists(flag *bool,long_url string, db *sql.DB) (error){
	record, err := db.Query(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM records where long_url = '%s')", long_url))
	if err != nil{
		return  err
	}
	for record.Next(){
		err = record.Scan(flag)
		if err != nil {
			log.Fatal("Can not check if data exists", err)
			return err
		}
	}
	return nil
}

func get_short_url(long_url string, db *sql.DB) (string, error){
	var res string

	record, err := db.Query(fmt.Sprintf("SELECT short_url FROM records WHERE long_url = ('%s')", long_url))
	if err != nil{
		return res, err
	}
	for record.Next(){
		err = record.Scan(&res)
		if err != nil{
			return res, err
		}
	}
	return res, nil
}

func get_long_url(short_url string, db *sql.DB) (string, error){
	var res string

	record, err := db.Query(fmt.Sprintf("SELECT long_url FROM records WHERE short_url = ('%s')", short_url))
	if err != nil{
		return res, err
	}
	for record.Next(){
		err = record.Scan(&res)
		if err != nil{
			return res, err
		}
	}
	return res, nil
}