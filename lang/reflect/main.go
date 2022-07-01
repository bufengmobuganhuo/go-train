package main

import (
	"fmt"
	"os"

	"mengyu.com/gotrain/lang/reflect/display"
	"mengyu.com/gotrain/lang/reflect/sexpr"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
}

func main() {
	var movie = Movie{
		Title: "Casablanca", Year: 1942,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	}
	display.Display("movie", movie)
	display.Display("os.Stderr", os.Stderr)

	bytes, err := sexpr.Marshal(movie)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(bytes))
}
