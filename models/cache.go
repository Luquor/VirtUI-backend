package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// cache pour le serveur
var cc cache.Cache

func CreateCache() {
	cc = *cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
}

// Add_Token définit un token dans le cache ou le change
func Add_Token(token string) {
	cc.Set(token, token, 5*time.Minute)
}

// Cette fonction recherche le token
func GetToken(token string) bool {
	token_est_la := false
	for i := range cc.Items() {
		if i == token {
			token_est_la = true
		}
	}
	return token_est_la
}

// generateRandomToken génère un token d'authentification aléatoire
func generateRandomToken() (string, error) {
	// Créer un slice de 32 octets pour le token
	tokenBytes := make([]byte, 32)

	// Remplir le slice avec des octets aléatoires
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Convertir les octets en une chaîne base64
	token := base64.StdEncoding.EncodeToString(tokenBytes)
	fmt.Println("Token : " + token)
	return token, nil
}

// fonction supprimant le token afin de supprimer la session utilisateur
func supprimerToken(token string) {
	if GetToken(token) {
		cc.Delete(token)
	} else {
		panic("session introuvable")
	}
}
