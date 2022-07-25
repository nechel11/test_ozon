#!/bin/bash
if ! [ $1 ] && ! [ $2 ] && ! [ $3 ] && ! [ $4 ] && ! [ $5 ] ; then
	echo "enter pg_host, db_port, db_name, db_user and db_password"
	fi

if [ $1 ] && [ $2 ] && [ $3 ] && [ $4 ] && [ $5 ]; then
	export	"db_host=$1" 
	export	"db_port=$2"
	export	"db_name=$3"
	export	"db_user=$4"
	export	"db_password=$5"
	echo "database env has been imported"
	fi