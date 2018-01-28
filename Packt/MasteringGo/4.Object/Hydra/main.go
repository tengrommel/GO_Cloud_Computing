package main

import (
	"github.com/tengrommel/GO_Cloud_Computing/Packt/MasteringGo/4.Object/Hydra/hlogger"
	"net/http"
	"fmt"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}
func sroot(writer http.ResponseWriter, request *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(writer, "Welcome to the Hydra software system")
	logger.Println("Received an http Get request on root url")
}

