package main

import "testing"

func TestHello(t *testing.T){
	
	//**********************first version********************
	// got := Hello("Jojo")
	// want := "Hello Jojo"

	// if got != want {
	// 	t.Errorf("want %s got %s", want, got)
	// }

	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	}
	t.Run("Saying Hello to people", func(t *testing.T){
		got := Hello("Jojo","")
		want := "Hello Jojo"
		assertCorrectMessage(t, got, want)

		
	})

	t.Run("Saying 'Hello World' when there is no argument in Hello()", func(t *testing.T){
		got := Hello("","")
		want := "Hello World"

		assertCorrectMessage(t, got, want)


	})

	t.Run("In Spanish", func(t *testing.T){
		got := Hello("Jose", "French")
		want := "Bonjour Jose"

		assertCorrectMessage(t,got,want)
	})

}