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

// Nested struct
type car struct {
	make       string
	model      string
	height     int
	frontwheel Wheel
	rearwheel  Wheel
}

// Embedded struct
type Truck struct {
	make  string
	model string
	Wheel
}

// Struct Methods
type auth_info struct {
	username string
	password string
}

func (auth_i auth_info) basic_auth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", auth_i.username, auth_i.password)
}

func main() {

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// 1) STRUCTS

	mycyclewheel := Wheel{}
	mycyclewheel.radius = 5
	mycyclewheel.material = "castIron"

	fmt.Printf("My cycle's wheel radius is %d amd it is made of %s \n", mycyclewheel.radius, mycyclewheel.material)

	mycar := car{}

	mycar.frontwheel.radius = 10
	mycar.frontwheel.material = "iron"

	fmt.Printf("My car's front wheel radius is %d amd it is made of %s \n", mycar.frontwheel.radius, mycar.frontwheel.material)

	// By default a struct's int values are 0 and strings are empty .... BUT ARE INITIALIZED

	fmt.Printf("My car's height is %d amd make is %s \n", mycar.height, mycar.make)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

	// 2) ANANONYMOUS STRUCT -- use when we need to crete just one instance of a struct.
	// More usage of anonomous struct is in nested structs... we can have an anonymous struct inside a struct

	mycar1 := struct {
		make  string
		model string
	}{
		make:  "jeep",
		model: "compass",
	}

	fmt.Printf("My car is %s %s \n", mycar1.make, mycar1.model)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

	// 3) EMBEDDED STRUCT : we insert a struct inside another completely and as it is

	mytruck := Truck{}

	mytruck.make = "Tata"
	mytruck.model = "Neo"
	mytruck.radius = 30 // so we are able to directly access wheel's properties unlike nested structs
	mytruck.material = "SolidIron"

	fmt.Printf("My truck is %s %s and wheel radius is %d which is made of %s \n", mytruck.make, mytruck.model, mytruck.radius, mytruck.material)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

	// 4) STRUCT METHODS - we can define a function on a struct which will manipulate the struct.

	myauth_info := auth_info{
		username: "rohit",
		password: "pass",
	}
	fmt.Println(myauth_info.basic_auth())
}
