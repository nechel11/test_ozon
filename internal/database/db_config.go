package database

import (
	"os"
	"fmt"
	"strconv"
)

func Get_env() string{
	var res string

	db_host := os.Getenv("db_host")
	db_port,_ := strconv.Atoi(os.Getenv("db_port"))
	db_name := os.Getenv("db_name")
	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_password")
	res = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	return res
}