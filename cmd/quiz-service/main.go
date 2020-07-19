package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	pokedexResponse, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	pokedexData, err := ioutil.ReadAll(pokedexResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(pokedexData))

	pokedexHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Got them all!\n")
	}

	http.HandleFunc("/hello", pokedexHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
