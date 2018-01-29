package Appliances

import (
	"fmt"
	"errors"
)

type Appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance() (Appliance, error) {
	//Request the user to enter the appliance type
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove ")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	// use fmt.scan to retrieve the user's input
	var myType int
	fmt.Scan(&myType)

	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
		//new case added for microwaves
	case MICROWAVE:
		return new(Microwave), nil

	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}