package main

import "fmt"

func main() {

	var eingabe int

	fmt.Print("Type: ")
	_, err := fmt.Scan(&eingabe) // pointer = & | ausserdem _ = zweiter Rückgabewert

	if err != nil {
		fmt.Println("Fehler bei der Eingabe:", err)
		return // Programm abbrechen
	}

	// TODO: Eingabe verarbeiten (z. B. Wort zählen, Text analysieren)

	// TODO: Ergebnis strukturieren (z. B. mit struct, map oder slice)

	fmt.Println("Ergebnis")
	fmt.Println(eingabe)
}

// TODO: Eigene Funktion(en) zur Analyse oder Berechnung implementieren
