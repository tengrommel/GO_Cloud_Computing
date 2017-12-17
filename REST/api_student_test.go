package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
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

// empty array back
func Test_handlerStudent_malformedURL(t *testing.T)  {
	ts := httptest.NewServer(http.HandlerFunc(handlerStudent))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/student/id/extra",
		ts.URL + "/stud/",
	}
	for _, tstring := range testCases{
		resp, err := http.Get(tstring)
		if err != nil{
			t.Errorf("Error making the GET request, %s", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("For route: %s, expected StatusCode %d, received %d", tstring,
				http.StatusBadRequest, resp.StatusCode)
			return
		}
	}
}


// GET /student/
// empty array back
func Test_handlerStudent_getAllStudents_empty(t *testing.T)  {
	ts := httptest.NewServer(http.HandlerFunc(handlerStudent))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/student/")
	if err != nil{
		t.Errorf("Error making the GET request, %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
		return
	}
	var a []interface{}
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil{
		t.Errorf("Error parsing the expected JSON body. Got error: %s", err)
	}
	if len(a) != 0 {
		t.Errorf("Excpected empty array. got %s", a)
	}
}

// GET /student/
// empty array back
func Test_handlerStudent_getAllStudents_Tom(t *testing.T)  {
	// add new student
	db = StudentsDB{}
	db.Init()
	testStudent := Student{"Tom", 21, "id0"}
	db.Add(testStudent)
	ts := httptest.NewServer(http.HandlerFunc(handlerStudent))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/student/")
	if err != nil{
		t.Errorf("Error making the GET request, %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
		return
	}
	var a []Student
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil{
		t.Errorf("Error parsing the expected JSON body. Got error: %s", err)
	}
	if len(a) != 1 {
		t.Errorf("Excpected array with one element, got %s", a)
	}
	if a[0].Name != testStudent.Name || a[0].Age != testStudent.Age || a[0].Id != testStudent.Id {
		t.Errorf("Students do not match! Got: %s, Expected: %s\n", a[0], testStudent)
	}
}

// GET /student/id0
// empty array back
func Test_handlerStudent_getStudents_Tom(t *testing.T)  {
	// add new student
	db = StudentsDB{}
	db.Init()
	testStudent := Student{"Tom", 21, "id0"}
	db.Add(testStudent)
	ts := httptest.NewServer(http.HandlerFunc(handlerStudent))
	defer ts.Close()

	// NotFound
	resp, err := http.Get(ts.URL + "/student/id1")
	if err != nil{
		t.Errorf("Error making the GET request, %s", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusNotFound, resp.StatusCode)
		return
	}


	// Tom
	resp, err = http.Get(ts.URL + "/student/id0")
	if err != nil{
		t.Errorf("Error making the GET request, %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusCode %d, received %d", http.StatusOK, resp.StatusCode)
		return
	}
	var a Student
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil{
		t.Errorf("Error parsing the expected JSON body. Got error: %s", err)
	}

	if a.Name != testStudent.Name || a.Age != testStudent.Age || a.Id != testStudent.Id {
		t.Errorf("Students do not match! Got: %s, Expected: %s\n", a, testStudent)
	}
}