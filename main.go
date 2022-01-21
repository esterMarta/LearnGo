package main

import (
	"fmt"

	"estermarta.com/go-learn-1/mascot"

	"path"
)

func main() {
	fmt.Println(mascot.BestMascot())

	var speedx int

	fmt.Println(speedx)

	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println(dir, file)

	speed := 100
	force := 2.5

	//speed = speed * int(force) //200
	speed = int(float64(speed) * force) //250

	fmt.Println(speed)

}
