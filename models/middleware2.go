package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/render"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password`
}

func enregistreToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//fmt.Println("HEADER_AUTHORIZATION=", r.Header.Get("Auhtorization"))

		if r.Header.Get("Authorization") == "" && r.RequestURI != "/auth" {
			rw.Header().Add("WWW-Authenticate", "Bearer")
			rw.WriteHeader(http.StatusUnauthorized)
			http.Redirect(rw, r, "/auth", http.StatusTemporaryRedirect)
		} else if r.Header.Get("Authorization") != "" {
			token := strings.Split(r.Header.Get("Authorization"), " ")
			if !verify_token(token[1]) {
				rw.WriteHeader(http.StatusUnauthorized)
			}
		}

		next.ServeHTTP(rw, r)
	})
}

func authenticate(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("fichier.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
	}
	defer file.Close()

	log.Print("Verification de votre identite")
	r.ParseForm()
	jsonResponse := User{}
	json.NewDecoder(r.Body).Decode(&jsonResponse)
	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)
	// Parcourir chaque ligne du fichier
	connec := false
	for scanner.Scan() {
		line := scanner.Text()
		credentials := strings.Split(line, ":")

		if credentials[0] == jsonResponse.Username && credentials[1] == jsonResponse.Password {
			connec = true
		}

	}
	if connec {
		token, err := generateRandomToken()
		fmt.Println(token)
		if err != nil {
			log.Panic("Erreur lors de la generation du token")
		} else {
			Add_Token(token)
			r.Header.Set("Authorization", "Bearer "+token)
			fmt.Println(cc)
			http.Redirect(w, r, "/", http.StatusAccepted)

		}
	} else {
		/// suite à une erreur l'auhtentification à echouer///////////////
		authFailed()
		//http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
	}
	render.JSON(w, r, jsonResponse)

}
func authFailed() {
	fmt.Println("failed connection")
}
func verify_token(token string) bool {
	_, exist := cc.Get(token)
	return exist
}
