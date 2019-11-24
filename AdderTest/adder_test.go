package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T){

	sum := Add(2,2)
	expected :=4

	if sum != expected{
		t.Errorf("got '%d' but expect '%d' ", sum , expected)
	}
	
}

func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	//Output : 6
}

func Add(x, y int) int{
	return x+y
}


