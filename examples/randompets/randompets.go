// program listpets lists random pets from the petfinder API
package main

import (
	"flag"
	"fmt"
	"os"
	petfinder "github.com/tmc/go-petfinder"
)

var (
	apiKey     = flag.String("key", "", "Petfinder API Key")
)

func main() {
	flag.Parse()
	api, err := petfinder.NewAPI(*apiKey, "")
	if err != nil {
		fmt.Println("error creating api:", err)
		os.Exit(1)
	}
	pets, err := api.RandomPets()
	if err != nil {
		fmt.Println("error fetching pets:", err)
		os.Exit(1)
	}
	for _, pet := range pets {
		fmt.Println("Pet:", pet)
	}
}
