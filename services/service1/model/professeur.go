package model

import (
	"gorm.io/gorm"
)

type Professeur struct {
	gorm.Model
	ID            int     `gorm:"column:id"`          // Correspond à la colonne "id" dans la table des professeurs
	Identifiant   string  `gorm:"column:identifiant"` // Correspond à la colonne "identifiant" dans la table des professeurs
	Nom           string  `gorm:"column:nom"`         // Correspond à la colonne "nom" dans la table des professeurs
	Cours         []Cours `gorm:"many2many:professeur_cours"`
	UtilisateurID uint    `gorm:"column:utilisateur_id;foreignKey:UtilisateurID"` // Clé étrangère qui pointe vers la colonne "id" dans la table des utilisateurs
	Utilisateur   Utilisateur
}

type ProfesseurData struct {
	Liste_professeur []Professeur
}
