package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("could not start app, %v", err)
		os.Exit(1)
	}
}

func run() error {
	return fmt.Errorf("not implemented yet")
}
