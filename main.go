package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var url string
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		fmt.Print("Entrez l'URL du site WordPress: ")
		fmt.Scanln(&url)
	}

	if isWordPressSite(url) {
		fmt.Printf("Le site %s est un site WordPress.\n", url)
	} else {
		fmt.Printf("Le site %s n'est pas un site WordPress.\n", url)
	}
}

func isWordPressSite(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP:", err)
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de l'analyse du document:", err)
		return false
	}

	// Vérifie la présence de balises spécifiques à WordPress
	if doc.Find("meta[name='generator'][content^='WordPress']").Length() > 0 {
		return true
	}

	return false
}
