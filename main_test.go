package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

//test de creation d'un contenaire
//On fait une requete en brut avec l'api lxd puis on compare à ce que l'on reçoit

func testCreationContainer() {

	fmt.Print(http.Get("127.0.0.1:8443/1.0/instance"))

}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

// Fonction principale
func main1() {
	// Liste de tests internes
	tests := []InternalTest{
		{"Test1", "abc", "abc", true},
		{"Test2", "123", "456", false},
		{"Test3", "go", "golang", false},
	}

	// Exécution des tests avec la fonction RunTests
	ok := RunTests(matchString, tests)

	// Affichage du résultat
	fmt.Println("Les tests se sont-ils bien déroulés ?", ok)
}

func matchString(pat, str string) (bool, error) {
	return pat == str, nil
}

// Structure représentant un test interne
type InternalTest struct {
	Nom              string
	resultat_entree  string
	resultat_attendu string
	comparaison      bool
}

// Fonction de test
func TestRunTests(t *testing.T) {
	// Liste de tests internes
	tests := []InternalTest{
		{"testCreationContainer", "0", "0", true},
		{"Test2", "123", "456", false},
		{"Test3", "go", "golang", true},
	}
	// Exécution des tests avec la fonction RunTests
	ok := RunTests(matchString, tests)

	// Vérification du résultat des tests
	if !ok {
		fmt.Print(http.Get("127.0.0.1:8443/1.0/instances"))

		t.Error("Les tests ont échoué.")
	}
}

// Fonction RunTests qui exécute les tests avec la fonction de correspondance de chaînes
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool) {
	for _, test := range tests {
		result, err := matchString(test.resultat_attendu, test.resultat_entree)
		if err != nil {
			// Gérer les erreurs si nécessaire
			fmt.Println("Erreur lors de l'exécution du test :", err)
			return false
		}

		if result != test.comparaison {
			// Le test a échoué
			fmt.Printf("Échec du test '%s'. Résultat attendu : %s, Résultat obtenu : %s\n", test.Nom, test.resultat_attendu, test.resultat_entree)
			return false
		}

		fmt.Printf("Test '%s' réussi\n", test.Nom)
	}

	// Tous les tests ont réussi

	return true
}

///////////////////////////
