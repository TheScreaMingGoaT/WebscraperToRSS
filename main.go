package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	// Erstelle einen neuen Collector
	c := colly.NewCollector()

	// Warte-Gruppe für parallele Ausführung
	var wg sync.WaitGroup

	// Schlagwort, nach dem gesucht wird
	keyword := "Niko"

	// Definiere die Aktion, die beim Besuch jeder Seite ausgeführt wird
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, keyword) {
			fmt.Println(link)
		}
	})

	// Starte den Webscraper für mehrere Seiten parallel
	for i := 1; i <= 5; i++ {
		// Inkrementiere die Warte-Gruppe
		wg.Add(1)

		// Goroutine für jeden Besuch
		go func(pageNumber int) {
			// Verzögere die Ausführung der Warte-Gruppe, wenn die Goroutine beendet ist
			defer wg.Done()

			// Besuche die Seite mit dem Suchbegriff
			err := c.Visit(fmt.Sprintf("https://www.hltv.org/search?query=%s&page=%d", keyword, pageNumber))
			if err != nil {
				fmt.Printf("Fehler beim Besuch der Seite %d: %v\n", pageNumber, err)
			}
		}(i)
	}

	// Warte auf das Ende aller Goroutinen
	wg.Wait()
}
