package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting hydraChatClient...")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	/*
	confStruct := struct {
		Name string `name:"name"`
		RemoteAddr string `name:"remoteip"`
		TCP bool `name:"tcp"`
	}{}

	HydraConfigurator.GetConfiguration(HydraConfigurator.CUSTOM,&confStruct,"chat.conf")
	name = confStruct.Name
	proto := "tcp"
	if !confStruct.TCP{
		proto = "udp"
	}
	*/
	fmt.Println("Hello %s, connecting to the hydra chat system...... \n", name)

}
