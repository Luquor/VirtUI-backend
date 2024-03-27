package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/patrickmn/go-cache"
)

// cache pour le serveur
var cc cache.Cache
var cache_tokens []string

func CreateCache() {
	cc = *cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
}

// Add_Token définit un token dans le cache ou le change
func Add_Token(token string) {
	cache_tokens = append(cache_tokens, token)
	cc.Set(token, cache_tokens, cache.DefaultExpiration)
}

// Cette fonction recherche le token
func GetToken(token string) bool {
	token_est_la := false
	for i := range cache_tokens {
		if cache_tokens[i] == token {
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
