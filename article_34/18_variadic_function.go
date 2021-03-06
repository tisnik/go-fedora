// Seriál "Programovací jazyk Go"
//
// Třicátá čtvrtá část
//    Programovací jazyk Go pro skalní céčkaře
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare/
//
// Demonstrační příklad číslo 18:
//    Variadické funkce.

package main

import "fmt"

func f1(msg string) {
	fmt.Printf("%s\n", msg)
}

func f2(parts ...string) {
	for _, val := range parts {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func f3(prefix string, parts ...string) {
	fmt.Printf("%s ", prefix)
	for _, val := range parts {
		fmt.Printf("%s ", val)
	}
	fmt.Println()
}

func main() {
	f1("Hello")
	f2("Hello", "world", "!")
	f3("Message:", "Hello", "world", "again", "!")
}
