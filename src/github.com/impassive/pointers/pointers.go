package main

import "fmt"

//position x,y
type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(guy *badGuy) {
	x := guy.pos.x
	y := guy.pos.y
	fmt.Println("(", x, ", ", y, ")")
}

func addOne(x *int) {
	*x = *x + 1
}

func main() {

	x := 5
	fmt.Println(x)

	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)
	p := position{4, 2}
	badguy := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(&badguy)

}
