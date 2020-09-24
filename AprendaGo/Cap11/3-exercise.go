package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourByFour bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle:    vehicle{2, "black"},
		fourByFour: true,
	}

	v2 := sedan{
		vehicle: vehicle{2, "red"},
		luxury:  false,
	}

	fmt.Println(v1)
	fmt.Println(v1.color)
	fmt.Println(v2)
	fmt.Println(v2.luxury)

}
