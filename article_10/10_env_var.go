// Seriál "Programovací jazyk Go"
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 10:
//    Přístup k proměnným prostředí.

package main

import (
	"fmt"
	"os"
)

func main() {
	env_vars := os.Environ()

	for i, env_var := range env_vars {
		fmt.Printf("%02d\t%s\n", i, env_var)
	}

	term, found := os.LookupEnv("TERM")
	if found {
		fmt.Printf("\n\n\nSelected TERM = %s", term)
	} else {
		fmt.Printf("\n\n\nThe TERM environment variable is not set")
	}
}
