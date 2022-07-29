package database

import (
	"os"
)

func Get_env() string{
	return os.Getenv("db_url")
}