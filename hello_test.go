package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Chris")
	want := "Hello Jojo"

	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}