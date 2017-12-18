package main

import (
	"regexp"
	"strings"
	"time"
	"fmt"
)

func tryString() {
	input := "frexit, brexit, nexit, lexit"
	strings.Contains(input, "frexit")
	strings.Contains(input, "brexit")
	strings.Contains(input, "grexit")
	strings.Contains(input, "nexit")
}

func tryRegex()  {
	pattern := "(br|n|gr|l)exit"
	input := "frexit, brexit, nexit, lexit"
	rgx, _:= regexp.Compile(pattern)
	rgx.FindAllString(input, -1)
}

func track(start time.Time, name string)  {
	duration := time.Since(start)
	fmt.Printf("Duration for %s: %s", name, duration)
}

func main() {
	defer track(time.Now(), "Evaluation")

	for i:=0;i<10000;i++ {
		//tryString()
		tryRegex()
	}

}
