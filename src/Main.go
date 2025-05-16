package main

import "fmt"

type Ergebnis struct {
	Zahl          int
	Eigenschaften map[string]string // [key]value
}

func main() {

	var eingabe int

	fmt.Print("Zahl: ")
	_, err := fmt.Scan(&eingabe) // pointer = & | ausserdem _ = zweiter Rückgabewert

	if err != nil {
		fmt.Println("Fehler bei der Eingabe:", err)
		return
	}

	ergebnis := analysiereZahl(eingabe)

	fmt.Println("Zahl:", ergebnis.Zahl)
	for k, v := range ergebnis.Eigenschaften { // range = gehe durch map (Eigenschaften)
		fmt.Println(k+":", v)
	}
}

func analysiereZahl(n int) Ergebnis {

	eigenschaften := make(map[string]string)

	if n > 0 {
		eigenschaften["Vorzeichen"] = "positiv"
	} else if n < 0 {
		eigenschaften["Vorzeichen"] = "negativ"
	} else {
		eigenschaften["Vorzeichen"] = "null"
	}

	return Ergebnis{ // zahl mit map geben
		Zahl:          n,
		Eigenschaften: eigenschaften,
	}
}
