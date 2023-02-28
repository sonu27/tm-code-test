package main

import (
	"fmt"
	"log"
	"os"
	"tm/internal/app"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("file path argument missing")
	}

	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := app.Start(file)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range output {
		fmt.Println(v)
	}
}
