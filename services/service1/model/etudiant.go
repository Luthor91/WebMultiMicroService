package model

import (
	"gorm.io/gorm"
)

type Etudiant struct {
	gorm.Model
	ID            int     `gorm:"column:id"`          // Correspond à la colonne "id" dans la table des étudiants
	Identifiant   string  `gorm:"column:identifiant"` // Correspond à la colonne "identifiant" dans la table des étudiants
	Nom           string  `gorm:"column:nom"`         // Correspond à la colonne "nom" dans la table des étudiants
	Cours         []Cours `gorm:"many2many:etudiant_cours"`
	UtilisateurID uint    `gorm:"column:utilisateur_id;foreignKey:UtilisateurID"` // Clé étrangère qui pointe vers la colonne "id" dans la table des utilisateurs
	Utilisateur   Utilisateur
}

type EtudiantData struct {
	Liste_etudiant []Etudiant
}
