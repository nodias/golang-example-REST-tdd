package main

import "os"

func main(){
	os.Setenv("TEST_DB_USERNAME","admin")
	os.Setenv("TEST_DB_PASSWORD","admin")
	os.Setenv("TEST_DB_NAME","postgres")
	os.Setenv("TEST_DB_HOST","54.180.2.92")
	os.Setenv("TEST_DB_PORT","5432")

	a:= App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		)
	a.Run(":8080")
}