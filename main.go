package main

import (
	"fmt"

	"github.com/insignificantGuy/Slotly/internal/database"
)

func main() {
	fmt.Println("Starting Slotly")
	database.InitMySQL()
}
