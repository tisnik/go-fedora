// Seriál "Programovací jazyk Go"
//
// Dvacátá osmá část
//
// Demonstrační příklad číslo 4:
//    Vytvoření streamů z polí struktur.

package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type User struct {
	id      uint32
	name    string
	surname string
}

func main() {
	var users = [3]User{
		User{
			id:      1,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      2,
			name:    "Pepek",
			surname: "Vyskoč"},
		User{
			id:      3,
			name:    "Josef",
			surname: "Vyskočil"},
	}
	fmt.Println(users)

	fmt.Printf("input:  %v\n", users)

	stream := koazee.StreamOf(users)
	fmt.Printf("stream: %v\n", stream.Out().Val())
}
