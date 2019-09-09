package design_pattern

import "testing"

func TestApiV1(t *testing.T) {
	api := NewApi("v1")
	reply := api.Hello()
	if reply != "Hello, World." {
		t.Fatal("Create ApiV1 Failed.")
	}
}

func TestApiV2(t *testing.T) {
	api := NewApi("v2")
	reply := api.Hello()
	if reply != "Hello, Golang." {
		t.Fatal("Create ApiV2 Failed.")
	}
}
