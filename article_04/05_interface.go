// Seriál "Programovací jazyk Go"
//
// Čtvrtá část
//
// Demonstrační příklad číslo 5:
//    Definice trojice jednoduchých rozhraní.

package main

// Shape je uživatelsky definovaná datová struktura
// představující geometrický tvar
type Shape interface {
}

// OpenShape je uživatelsky definovaná datová struktura
// představující otevřené tvary (úsečka, oblouk, křivka)
type OpenShape interface {
	length() float64
}

// ClosedShape je uživatelsky definovaná datová struktura
// představující uzavřené tvary (obdélník, kružnice, elipsa)
type ClosedShape interface {
	area() float64
	perimeter() float64
}

func main() {
}
