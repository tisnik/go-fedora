// Seriál "Programovací jazyk Go"
//
// První část
//
// Demonstrační příklad číslo 6:
//    Pokus o opětovnou deklaraci proměnných

package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	b := "hello"
	fmt.Println(b)
	c := true
	fmt.Println(c)

	a := 20
	fmt.Println(a)
	b := "world"
	fmt.Println(b)
	c := false
	fmt.Println(c)
}
