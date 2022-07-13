package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const  (
	db_host = "127.0.0.1" 
	db_port=5432
	db_name="ozon"
	db_user = "zafar"
	db_password="12344321"
)

func Db_connect() (x *sql.DB){
	psqconn := fmt.Sprintf("host= %s port = %d user = %s password = %s dbname = %s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	db, err := sql.Open("postgres", psqconn)
	if err != nil {
		panic(err)
	}
	return db
}