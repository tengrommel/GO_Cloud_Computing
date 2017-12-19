package Webhooks

import (
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

type Message struct {
	Text string `json:"text"`
	Channel string `json:"channel, omitempty"`
	Username string `json:"username, omitempty"`
	Icon string `json:"icon_url, omitempty"`
}

func main()  {
	msg := Message{"Another Go example", "#general","",""}
	msh, err := json.Marshal(msg)
	if err!= nil{
		fmt.Errorf("Error during marshalling: %v", err.Error())
	}
	url := "https//"
	res, err := http.DefaultClient.Post(url, "application/json", bytes.NewReader(msh))
	if err != nil{
		fmt.Println("Error during POST: %v", err.Error())
	}

	if res.StatusCode != http.StatusOK{
		fmt.Errorf("Wrong")
	}

}