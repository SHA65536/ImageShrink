package main

import (
	"github.com/sha65536/imageshrink"
)

func main() {
	err := imageshrink.ShrinkFile("input.png")
	if err != nil {
		panic(err)
	}
}
