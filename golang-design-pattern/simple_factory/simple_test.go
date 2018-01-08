package simple_factory

import "testing"

// TestType1 test get hiapi with factory
func TestType1(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("TOM")
	if s != "Hi, TOM"{
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T)  {
	api := NewAPI(2)
	s := api.Say("TOM")
	if s != "Hello, TOM"{
		t.Fatal("Type2 test fail")
	}
}