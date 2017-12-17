package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test_handlerStudent_notImplemented(t *testing.T)  {
	// instantiate mock HTTP server
	// register our handlerStudent <-- actual logic

	ts := httptest.NewServer(http.HandlerFunc(handlerStudent))
	defer ts.Close()

	// create a request to our mock HTTP server
	// in our case it means to create DELETE request

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the DELETE request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil{
		t.Errorf("Error executing the DELETE request, %s", err)
	}

	// check if the response from the handler is what we expect
	if resp.StatusCode != http.StatusNotImplemented{
		t.Errorf("Expected StatusCode %d, received %d", http.StatusNotImplemented, resp.StatusCode)
	}
}