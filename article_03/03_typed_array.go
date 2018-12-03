// Seriál "Programovací jazyk Go"
//
// Třetí část
//
// Demonstrační příklad číslo 3:
//    Pole s prvky definovaného typu.

package main

import "fmt"

type Mesic string

func main() {
	var mesice [12]Mesic

	fmt.Println(mesice)

	mesice[0] = "Leden"
	mesice[1] = "Únor"
	mesice[2] = "Březen"
	mesice[3] = "Duben"
	mesice[4] = "Květen"
	mesice[5] = "Červen"
	mesice[6] = "Červenec"
	mesice[7] = "Srpen"
	mesice[8] = "Září"
	mesice[9] = "Říjen"
	mesice[10] = "Listopad"
	mesice[11] = "Prosinec"

	fmt.Println(mesice)
}
