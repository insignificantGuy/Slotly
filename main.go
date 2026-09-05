package main

import (
	"fmt"

	"github.com/insignificantGuy/Slotly/internal/bootstrap/apiserver"
)

func main() {
	fmt.Println("Starting Slotly")
	if err := apiserver.Start(); err != nil {
		panic(err)
	}
}
