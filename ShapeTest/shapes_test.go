package main

import "testing"
// import "math"

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0,10.0}
// 	got := Perimeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T){

// 	t.Run("rectangles", func (t *testing.T){
// 		rectangle := Rectangle{9.0 , 9.0}
// 		got := rectangle.Area()
// 		want := 81.0
	
// 		if got != want{
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}

// 	})

// 	t.Run("circles", func(t *testing.T){
	// 		circle := Circle{10}
	// 		got := circle.Area()
	// 		want:= 314.1592653589793
	
	// 		if(got != want){
		// 			t.Errorf("got %.2f want %.2f", got, want)
		// 		}
		// 	})
		// }


		////////////////////////////////////
		// Second Evolution/////////////////
		////////////////////////////////////


// func TestArea(t *testing.T){
			
// 	checkArea := func(t *testing.T, shape Shape, want float64){
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want{
// 			t.Errorf("got %.2f want %.2f",got ,want)
// 		}
// 	}


// 	t.Run("rectangles", func (t *testing.T){
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
		
// 	})

// 	t.Run("circles", func(t *testing.T){
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// type Shape interface{
// 	Area() float64
// }
// 	t.Run("circles", func(t *testing.T){
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want:= 314.1592653589793

// 		if(got != want){
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})
// }


/////////////////////////////////////////////
// third refactorization/////////////////////
/////////////////////////////////////////////
func TestArea(t *testing.T){

	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	}{
		//*****Before
		// {Rectangle{12.0,6.0}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12,6}, 36},


		// *****After
		{name:"Rectangle", shape:Rectangle{Width:12, Height:6}, hasArea: 72.0},
		{name: "Circle", shape:Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape:Triangle{Base:12, Height:6}, hasArea:37},
	}


	for _,tt := range areaTests{

		t.Run(tt.name, func(t *testing.T){
			got := tt.shape.Area()
	
			if got != tt.hasArea{
				t.Errorf(" %#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}

		})
	}
}

type Shape interface{
	Area() float64
}
