package main

import (
	"fmt"
	"os"
	"strings"
)

func moo(animal string) string {
	msg := fmt.Sprintf("<%s says moo!>", animal)
	top := " " + strings.Repeat("_", len(msg))
	bot := " " + strings.Repeat("-", len(msg))
	return fmt.Sprintf("%s\n%s\n%s\n        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||", top, msg, bot)
}

func main() {
	animal := "Cow"
	if len(os.Args) > 1 {
		animal = os.Args[1]
	}
	fmt.Println(moo(animal))
}
