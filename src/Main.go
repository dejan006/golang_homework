package main

import "fmt"

func main() {

	var eingabe int

	fmt.Print("Type: ")
	_, err := fmt.Scan(&eingabe) // pointer = & + _ = zweiter Rückgabewert

	if err != nil {
		fmt.Println("Fehler bei der Eingabe:", err)
		return // Programm abbrechen
	}

	// TODO: Fehlerbehandlung bei der Eingabe

	// TODO: Eingabe verarbeiten (z. B. Wort zählen, Text analysieren)

	// TODO: Ergebnis strukturieren (z. B. mit struct, map oder slice)

	// TODO: Ergebnisse ausgeben (formatierte Ausgabe)
}

// TODO: Eigene Funktion(en) zur Analyse oder Berechnung implementieren
