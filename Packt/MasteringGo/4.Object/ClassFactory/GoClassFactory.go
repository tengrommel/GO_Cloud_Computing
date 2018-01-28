package main

import (
	"github.com/tengrommel/GO_Cloud_Computing/Packt/MasteringGo/4.Object/ClassFactory/Appliances"
	"fmt"
)

func main() {
	myAppliance, err := Appliances.CreateAppliance()

	if err !=nil{
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
