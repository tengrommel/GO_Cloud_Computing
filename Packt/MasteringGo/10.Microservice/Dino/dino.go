package main

import (
	"github.com/tengrommel/GO_Cloud_Computing/Packt/MasteringGo/10.Microservice/Dino/dynowebportal"
	"os"
	"log"
	"encoding/json"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil{
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}
