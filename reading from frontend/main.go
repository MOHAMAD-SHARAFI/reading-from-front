package main

import (
	"log"
	"project/cmd"
	"project/internal/config"
)

func main() {
	if err := cmd.Run(config.Default); err != nil {
		log.Fatal(err)
	}
}
