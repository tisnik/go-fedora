// Seriál "Programovací jazyk Go"
//
// Třetí část
//    Datové typy v programovacím jazyku Go (2.část)
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go-2-cast/
//
// Demonstrační příklad číslo 5:
//    Kontrola uživatelsky definovaných typů.

package main

import "fmt"

// ID of user
type ID uint32

// Name of user
type Name string

// Surname of user
type Surname string

func main() {
	var i ID
	i = 0
	fmt.Println(i)

	var str = "common string"

	var n Name = str
	var s Surname = str

	n = "Jan"
	s = "Novák"

	fmt.Println(n)
	fmt.Println(s)
}
