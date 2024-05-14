package database

import (
	"fmt"
	"log"

	"os"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func CreatePostgresDatabase() bool {
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_ADMIN_PASSWORD")
	user := os.Getenv("DB_ADMIN_USER")
	host := os.Getenv("DB_HOST")
	sslMode := os.Getenv("DB_SSL_MODE")

	connexion := fmt.Sprintf("user=%s password=%s host=%s sslmode=%s", user, password, host, sslMode)

	db_admin, err := sql.Open("postgres", connexion)
	if err != nil {
		log.Fatalf("Erreur lors de la connexion en tant qu'administrateur: %v", err)
	}
	defer db_admin.Close()

	stmt := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, _ = db_admin.Exec(stmt)

	// Vérifier si la base de données a été créée avec succès
	dbExists, err := CheckDatabaseExists(dbName)
	if err != nil {
		log.Fatalf("Erreur lors de la vérification de l'existence de la base de données: %v", err)
	}
	return dbExists
}

func CheckDatabaseExists(dbName string) (bool, error) {
	password := os.Getenv("DB_ADMIN_PASSWORD")
	user := os.Getenv("DB_ADMIN_USER")
	host := os.Getenv("DB_HOST")
	sslMode := os.Getenv("DB_SSL_MODE")

	connexion := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", user, password, host, dbName, sslMode)

	db, err := sql.Open("postgres", connexion)
	if err != nil {
		return false, err
	}
	defer db.Close()

	var dbFound string
	err = db.QueryRow("SELECT datname FROM pg_database WHERE datname = $1", dbName).Scan(&dbFound)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// ConnectToPostgres initialise la connexion à la base de données PostgreSQL
func ConnectToPostgres() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	connexion := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", user, password, host, dbName, sslMode)
	db, err := gorm.Open(postgres.Open(connexion), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur lors de la création de la base de données: %v", err)
	}

	return db
}
