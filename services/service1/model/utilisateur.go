package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Utilisateur struct {
	gorm.Model
	ID         int    `gorm:"column:id"`           // Correspond à la colonne "id" dans la table des utilisateurs
	Nom        string `gorm:"column:nom"`          // Correspond à la colonne "nom" dans la table des utilisateurs
	Email      string `gorm:"column:email"`        // Correspond à la colonne "email" dans la table des utilisateurs
	MotDePasse string `gorm:"column:mot_de_passe"` // Correspond à la colonne "mot_de_passe" dans la table des utilisateurs
}

type UtilisateurData struct {
	Liste_utilisateur []Utilisateur
}

// SetMotDePasse définit le mot de passe de l'utilisateur en le chiffrant avec bcrypt
func (u *Utilisateur) SetMotDePasse(local_password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(local_password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.MotDePasse = string(hashed_password)
	return nil
}

// CompareMotDePasse compare le mot de passe fourni avec le mot de passe de l'utilisateur chiffré
func (u *Utilisateur) CompareMotDePasse(local_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.MotDePasse), []byte(local_password))
	return err == nil
}
