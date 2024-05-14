package migration

import (
	"log"
	"main_module/model"

	"gorm.io/gorm"
)

func MigrateAllPostgresql(db *gorm.DB) {

	MigrateUtilisateurPostgresql(db)
	MigrateCoursPostgresql(db)
	MigrateEtudiantPostgresql(db)
	MigrateProfesseurPostgresql(db)

}

func MigrateCoursPostgresql(db *gorm.DB) {
	err := db.AutoMigrate(&model.Cours{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}
}

func MigrateEtudiantPostgresql(db *gorm.DB) {
	err := db.AutoMigrate(&model.Etudiant{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}
}

func MigrateProfesseurPostgresql(db *gorm.DB) {
	err := db.AutoMigrate(&model.Professeur{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}
}

func MigrateUtilisateurPostgresql(db *gorm.DB) {
	err := db.AutoMigrate(&model.Utilisateur{})
	if err != nil {
		log.Fatalf("Erreur de migration de la base de données: %v", err)
	}
}
