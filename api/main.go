package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

type application struct{
	port int
	jwtSecret string
	models database.Models
}


func main (){
	db,err := sql.Open("sqlite3","./data.db")
	if err !=nil{
		log.Fatal(err)
	}
	defer db.Close()
	models := database.NewModels(db)
	app:= &application("PORT",8080),
	jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-123456"),
	models: models,
}
if err := app.serve();err !=nil 
log.Fatal(err)
}}