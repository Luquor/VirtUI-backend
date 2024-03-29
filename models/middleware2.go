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
	Password string `json:"password"`
}

type Token struct {
	Token string
}

func enregistreToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//fmt.Println("HEADER_AUTHORIZATION=", r.Header.Get("Auhtorization"))

		if r.Header.Get("Authorization") == "" && r.RequestURI != "/auth" {
			rw.Header().Add("WWW-Authenticate", "Bearer")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		} else if r.Header.Get("Authorization") != "" && r.RequestURI != "/auth" {
			token := strings.Split(r.Header.Get("Authorization"), " ")
			if !verify_token(token[1]) {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(rw, r)
	})
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	var token string
	file, err := os.Open("users.txt")
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
			break
		}

	}
	if connec {
		token, err = generateRandomToken()
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
	var tokenJson = Token{Token: token}
	render.JSON(w, r, tokenJson)

}
func authFailed() {
	fmt.Println("failed connection")
}
func verify_token(token string) bool {
	_, exist := cc.Get(token)
	return exist
}
func deconnection(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	supprimerToken(token)
	fmt.Println(token + "token supprimé")
}
