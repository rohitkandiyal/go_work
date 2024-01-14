// Interface is a collection of abstract method signatures. The data type which implements this interface must define these methods as concrete.

package main
import (
	"fmt"
)

// Defining interface

type shape interface {
	area() float64
	perimeter() float64
}


// Now rectange below implements this interface

type rectange struct {
	length, breadth float64
}
func (r rectange) area() float64 {				// Struct method
	return r.length * r.breadth
}
func (r rectange) perimeter() float64 {			// Struct method
	return 2 * (r.length + r.breadth)
}

// Now circle below too can similarily implement shape interface

func main()  {
	my_rect := rectange {
		length: 6.50,
		breadth: 8.50,
	}
	fmt.Println(my_rect.area())
	fmt.Println(my_rect.perimeter())
}



// LEARN ABOUT TYPE ASSERTIONS AND TYPE SWITCH... left as of now