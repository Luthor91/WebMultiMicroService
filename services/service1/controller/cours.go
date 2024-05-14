package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main_module/database"
	"main_module/model"
	"time"

	"github.com/redis/go-redis/v9"
)

// GetAllCours récupère tous les cours depuis Redis s'ils existent, sinon depuis la base de données PostgreSQL
func GetAllCours() ([]model.Cours, error) {
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Supprimer la clé "cours:list" si elle existe
	if err := rdb.Del(ctx, "cours:list").Err(); err != nil {
		return nil, err
	}

	// Si les cours existent dans la base de données PostgreSQL, les récupérer depuis PostgreSQL
	db := database.ConnectToPostgres()
	var cours []model.Cours
	result := db.Find(&cours)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convertir les cours en JSON
	coursBytes, err := json.Marshal(cours)
	if err != nil {
		return nil, err
	}

	// Stocker les cours dans Redis
	err = rdb.Set(ctx, "cours:list", string(coursBytes), 0).Err()
	if err != nil {
		return nil, err
	}

	return cours, nil
}

// GetCoursByProfesseurID récupère tous les cours associés à un utilisateur depuis Redis s'ils existent, sinon depuis la base de données PostgreSQL
func GetCoursByUserID(userID int, typeUser string) ([]model.Cours, error) {

	// Vérifie bien que l'utilisateur en question est bien un professeur ou un étudiant
	if typeUser != "cours" && typeUser != "professeur" {
		return nil, errors.New("user must be cours or professeur")
	}

	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Clé Redis pour les cours associés à un utilisateur
	key := fmt.Sprintf("%s:%d:cours", typeUser, userID)

	// Récupérer les cours associés a l'utilisateur depuis Redis s'ils existent
	coursJSON, err := rdb.Get(ctx, key).Bytes()
	if err == nil {
		var cours []model.Cours
		if err := json.Unmarshal(coursJSON, &cours); err != nil {
			return nil, err
		}
		return cours, nil
	} else if err != redis.Nil {
		return nil, err
	}

	// Si les cours associés a l'utilisateur n'existent pas dans Redis, les récupérer depuis la base de données PostgreSQL
	db := database.ConnectToPostgres()
	var cours []model.Cours
	query := fmt.Sprintf("fk_%s = ?", typeUser)
	if err := db.Where(query, userID).Find(&cours).Error; err != nil {
		return nil, err
	}

	// Si les cours ont été récupérés depuis PostgreSQL, les stocker dans Redis
	coursJSON, err = json.Marshal(cours)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, key, coursJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return cours, nil
}

// GetCoursById récupère un cours par son ID depuis Redis s'il existe, sinon depuis la base de données PostgreSQL
func GetCoursById(id int) (*model.Cours, error) {
	if id <= 0 {
		return nil, errors.New("ID must be strictly greater than 0")
	}
	// Connexion à Redis
	rdb := database.ConnectToRedis()
	defer rdb.Close()

	ctx := context.Background()

	// Récupérer le etudiant depuis Redis s'il existe
	coursJSON, err := rdb.Get(ctx, fmt.Sprintf("cours:%d", id)).Bytes()
	if err == nil {
		var cours model.Cours
		if err := json.Unmarshal(coursJSON, &cours); err != nil {
			return nil, err
		}
		return &cours, nil
	} else if err != redis.Nil {
		fmt.Println("Error Redis :", err)
		return nil, err
	}

	// Si le cours n'existe pas dans Redis, le récupérer depuis la base de données PostgreSQL
	db := database.ConnectToPostgres()
	var cours model.Cours
	if err := db.First(&cours, id).Error; err != nil {
		return nil, err
	}

	// Si le cours a été récupéré depuis PostgreSQL, le stocker dans Redis
	coursJSON, err = json.Marshal(cours)
	if err != nil {
		return nil, err
	}
	err = rdb.Set(ctx, fmt.Sprintf("cours:%d", id), coursJSON, 0).Err()
	if err != nil {
		return nil, err
	}

	return &cours, nil
}

// InsertCours insère un cours dans la base de données Postgresql et dans le cache Redis
func InsertCours(cours *model.Cours) error {

	rd := database.ConnectToRedis()
	db := database.ConnectToPostgres()

	// Insérer le cours dans la base de données PostgreSQL
	if err := db.Create(cours).Error; err != nil {
		return err
	}

	// Convertir le cours en JSON
	coursJSON, err := json.Marshal(cours)
	if err != nil {
		return err
	}

	// Insérer le cours dans Redis
	ctx := context.Background()
	key := fmt.Sprintf("cours:%d", cours.ID) // Utiliser l'id du cours comme clé Redis
	expiration := 24 * time.Hour
	err = rd.Set(ctx, key, coursJSON, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateCours met à jour un cours dans la base de données PostgreSQL et dans Redis
func UpdateCours(cours *model.Cours) error {
	db := database.ConnectToPostgres()

	// Mettre à jour le cours dans la base de données PostgreSQL
	if err := db.Model(&model.Cours{}).Where("id = ?", cours.ID).Updates(cours).Error; err != nil {
		return err
	}

	// Mettre à jour le cours dans Redis
	rd := database.ConnectToRedis()
	defer rd.Close()

	ctx := context.Background()

	// Convertir le cours en JSON
	coursJSON, err := json.Marshal(cours)
	if err != nil {
		return err
	}

	// Mettre à jour le cours dans Redis avec la clé correspondant à son ID
	err = rd.Set(ctx, fmt.Sprintf("cours:%d", cours.ID), coursJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
