// Seriál "Programovací jazyk Go"
//
// Čtyřicátá druhá část
//    Testování aplikací psaných v Go s využitím knihoven Goblin a Frisby
//    https://www.root.cz/clanky/testovani-aplikaci-psanych-v-go-s-vyuzitim-knihoven-goblin-a-frisby/

package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestBefore(t *testing.T) {
	g := Goblin(t)
	x := 0

	g.Describe("Adder", func() {
		g.It("x+1", func() {
			g.Assert(x + 1).Equal(1)
		})
	})

	g.Describe("Adder", func() {
		g.It("x+1", func() {
			g.Assert(x + 1).Equal(11)
		})
		g.Before(func() {
			x = 10
		})
	})

	g.Describe("Adder", func() {
		g.It("x+1", func() {
			g.Assert(x + 1).Equal(2)
		})
		g.Before(func() {
			x = 1
		})
	})
}
