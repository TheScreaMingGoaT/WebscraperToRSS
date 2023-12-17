package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadHTML(url, filename string) error {
	// HTTP-Request, um die HTML-Seite herunterzuladen
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Datei erstellen, um die heruntergeladene HTML zu speichern
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Heruntergeladene HTML in die Datei schreiben
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("HTML-Seite wurde erfolgreich unter '%s' gespeichert.\n", filename)
	return nil
}

func main() {
	url := "https://www.hltv.org/"
	filename := "downloaded_page.html"

	err := downloadHTML(url, filename)
	if err != nil {
		fmt.Println("Fehler beim Herunterladen der HTML-Seite:", err)
	}
}
