package model

import (
	"gorm.io/gorm"
)

type Cours struct {
	gorm.Model
	ID              uint         `gorm:"column:id"`                  // Correspond à la colonne "id" dans la table des cours
	Identifiant     string       `gorm:"column:identifiant"`         // Correspond à la colonne "identifiant" dans la table des cours
	Titre           string       `gorm:"column:titre"`               // Correspond à la colonne "titre" dans la table des cours
	Resume          string       `gorm:"column:resume"`              // Correspond à la colonne "resume" dans la table des cours
	Niveau          string       `gorm:"column:niveau"`              // Correspond à la colonne "niveau" dans la table des cours
	PlaceDisponible uint         `gorm:"column:place_disponible"`    // Correspond à la colonne "place_disponible" dans la table des cours
	TempsExpiration uint         `gorm:"column:temps_expiration"`    // Correspond à la colonne "temps_expiration" dans la table des cours
	Professeurs     []Professeur `gorm:"many2many:professeur_cours"` // Relation many-to-many avec la table des professeurs
	Etudiants       []Etudiant   `gorm:"many2many:etudiant_cours"`   // Relation many-to-many avec la table des étudiants
}

type CoursData struct {
	Liste_cours []Cours
}
