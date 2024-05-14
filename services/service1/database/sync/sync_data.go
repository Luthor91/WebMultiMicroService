package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"main_module/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Fonction pour synchroniser les données entre Redis et PostgreSQL
func SyncData(redisClient *redis.Client, db *gorm.DB) error {
	ctx := context.Background()

	// Récupérer tous les cours depuis Redis
	cours, err := redisClient.LRange(ctx, "cours:list", 0, -1).Result()
	if err != nil {
		return err
	}

	// Insérer chaque cours dans la table Cours de PostgreSQL
	for _, coursJSON := range cours {
		var coursData model.Cours
		err := json.Unmarshal([]byte(coursJSON), &coursData)
		if err != nil {
			return err
		}
		if err := db.Create(&coursData).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion du cours:", err)
			return err
		}
		fmt.Println("Cours inséré avec succès:", coursData.Titre)
	}

	// Récupérer tous les étudiants depuis Redis
	etudiants, err := redisClient.LRange(ctx, "etudiant:list", 0, -1).Result()
	if err != nil {
		return err
	}

	// Insérer chaque étudiant dans la table Etudiant de PostgreSQL
	for _, etudiantJSON := range etudiants {
		var etudiantData model.Etudiant
		err := json.Unmarshal([]byte(etudiantJSON), &etudiantData)
		if err != nil {
			return err
		}
		if err := db.Create(&etudiantData).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion de l'étudiant:", err)
			return err
		}
		fmt.Println("Étudiant inséré avec succès:", etudiantData.Nom)
	}

	// Récupérer tous les professeurs depuis Redis
	professeurs, err := redisClient.LRange(ctx, "professeur:list", 0, -1).Result()
	if err != nil {
		return err
	}

	// Insérer chaque professeur dans la table Professeur de PostgreSQL
	for _, professeurJSON := range professeurs {
		var professeurData model.Professeur
		err := json.Unmarshal([]byte(professeurJSON), &professeurData)
		if err != nil {
			return err
		}
		if err := db.Create(&professeurData).Error; err != nil {
			fmt.Println("Erreur lors de l'insertion du professeur:", err)
			return err
		}
		fmt.Println("Professeur inséré avec succès:", professeurData.Nom)
	}

	fmt.Println("La synchronisation des données entre Redis et PostgreSQL a été effectuée avec succès.")
	return nil
}
