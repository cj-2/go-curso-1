package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicativo := app.Gerar()

	//error := aplicativo.Run(os.Args)
	// if error != nil {
	// 	log.Fatal(error)
	// }

	// Simplificado, comum do Go
	if error := aplicativo.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}
