// Seriál "Programovací jazyk Go"
//
// Dvacátá první část
//     Knihovny pro Go umožňující naplánování a spouštění periodických úloh
//     https://www.root.cz/clanky/knihovny-pro-go-umoznujici-naplanovani-a-spousteni-periodickych-uloh/
//
// Demonstrační příklad číslo 6:
//     Práce s objektem typu Duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("3h25m")
	e, _ := time.ParseDuration("15m")
	f := d.Round(e)

	fmt.Println(f.String())
	fmt.Printf("Hours:   %4.2f\n", f.Hours())
	fmt.Printf("Minutes: %2.0f\n", f.Minutes())
	fmt.Printf("Seconds: %4.0f\n", f.Seconds())
	fmt.Printf("ns:      %d\n", f.Nanoseconds())
}
