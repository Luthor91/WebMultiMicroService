package main

import (
	"log"
	"main_module/server"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	/*
	rd := database.ConnectToRedis()
	database.CreatePostgresDatabase()
	db := database.ConnectToPostgres()
	migration.MigrateAllPostgresql(db)
	sync.SyncData(rd, db)
	*/
	
	server.DefineRoutes()
}
