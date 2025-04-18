package main

import (
	"fmt"
	"io/ioutil"

	"math/rand"

	"github.com/innv8/sampleapp/shape"
)

func main() {
	fmt.Println("Hello world")

	var rec = shape.Rectangle{Width: 4, Height: 6}
	doShapeStuff(rec)

	rand.Seed(1)
	_ = rand.Int()
	_ = rand.Int()

	_, _ = ioutil.ReadAll(nil)

	//const greeting = "Hello World"

	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")

	var sq = shape.Square{Side: 4}
	doShapeStuff(sq)
}

func doShapeStuff(s shape.Shape) {
	fmt.Printf("The area is: %.2f, and the perimeter is %.2f\n",
		s.Area(), s.Perimeter())
}
