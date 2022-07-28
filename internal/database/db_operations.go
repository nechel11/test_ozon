package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Db_connect() (*sql.DB, error){
	psqconn := Get_env()
	db, err := sql.Open("postgres", psqconn)
	if err != nil {
		return db, err
	}
	return db, nil
}

func Db_insert_url(short_url, long_url string, db *sql.DB) error{
	_, err := db.Exec(fmt.Sprintf("INSERT INTO records (long_url, short_url) VALUES ('%s', '%s')", long_url, short_url))
	if err != nil {		
		return err
	}
	return nil
}

func Db_if_data_exists(flag *bool,long_url string, db *sql.DB) (error){
	record, err := db.Query(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM records where long_url = '%s')", long_url))
	if err != nil{
		return  err
	}
	for record.Next(){
		err = record.Scan(flag)
		if err != nil {
			return err
		}
	}
	return nil
}

func Db_get_short_url(long_url string, db *sql.DB) (string, error){
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

func Db_get_long_url(short_url string, db *sql.DB) (string, error){
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