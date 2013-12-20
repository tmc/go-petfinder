// program listbreeds lists animal breeds with the petfinder API
package main

import (
	"flag"
	"fmt"
	"os"
	petfinder "github.com/tmc/go-petfinder"
)

var (
	apiKey     = flag.String("key", "", "Petfinder API Key")
	animalName = flag.String("animal", "dog", "Animal name")
)

func main() {
	flag.Parse()
	api, err := petfinder.NewAPI(*apiKey, "")
	if err != nil {
		fmt.Println("error creating api:", err)
		os.Exit(1)
	}
	breeds, err := api.Breeds(*animalName)
	if err != nil {
		fmt.Println("error fetching breeds:", err)
		os.Exit(1)
	}
	for _, breed := range breeds {
		fmt.Println("Breed:", breed)
	}
}
