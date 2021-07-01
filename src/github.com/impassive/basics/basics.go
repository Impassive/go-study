//* this package is executable
package main

import "fmt"

//* executable func
func main() {
	// simple call
	fmt.Println("Hello, World!")
	//auto var
	var a int //integer
	// the initial is zero
	fmt.Println(a)
	//assign
	a = a + 5
	a -= 2
	var b = 4
	fmt.Println(a, b)

	// sizes of int

	// var a : int8  // -127 127
	// var b : uint8 // 0 255
	// var c : int16 //
	// var d : uint16
	// var e int // 32 bits or 64 bits depending on machine

	// floats

	// var f1 float32
	// var f2 float64

	//short typing + types mismatch
	c := 5
	d := 3.6 // 32 bit or 64 float depenging on machine

	fmt.Println(c + int(d))

	// overload

	var a1 uint8 = 255
	var a2 uint8 = 4
	fmt.Println(a1, "+", a2, "=", a1+a2, "uint!")

	// strage decimal

	b1 := 2.0
	b2 := 3.0

	fmt.Println(b1 / b2)

	//booleans

	truf := true

	var bln bool = false

	fmt.Println(truf, bln)

	var a3 string = "hello"
	fmt.Println(a3)

	//loops
	for {
		fmt.Println("Infinite Loop")
		break // exit
	}

	//loop with statement
	i := 0

	for i < 4 {
		fmt.Print(i)
		i++
	}
	fmt.Print(" ")

	// all in one
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}

	// conditionals
	i = 0
	if i > 1 {
		fmt.Print("Wrong")
	} else if i > 10 {
		fmt.Print("Yep")
	} else {
		fmt.Print("Nope")
	}

	bln = false

	if !bln {
		fmt.Println("NewLine")
	}

	switch bln {
	case true:
		fmt.Println("true")
		break
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}
}
