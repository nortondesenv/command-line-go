package main

import (
	"app-cmd/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Initialize...")

	cli := app.Gerar()
	err := cli.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
