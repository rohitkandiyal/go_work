// Structs here are like dicts. But we can create methods(check 4 below) on a struct and in that way it behaves like a class.
// As it can have attributes and methods both.

package main

import (
	"fmt"
)

type Wheel struct {
	radius   int
	material string
}

// STRUCT Method
func (wheel *Wheel) printWheelDetails() {
	// ==> CONVENIENCE 1: Within method we NEED NOT TO DEREF a pointer. Therefore below both values are valid.
	fmt.Println("Radius is:", (*wheel).radius, "and Material is :", wheel.material)
}

func main() {
	cycleWheel := Wheel{22, "iron"}
	busWheel := Wheel{44, "aluminium"}

	// Way 1 to call method
	(&cycleWheel).printWheelDetails()

	// ==> CONVENIENCE 2: Way 2
	busWheel.printWheelDetails()

	// ==> CONVENIENCE 3: Now I don't want to define vars but directly call method for scooterWheel
	(*Wheel).printWheelDetails(&Wheel{15, "Platic"})
}
