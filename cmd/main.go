package main

import (
	"fmt"

	"github.com/sha65536/imageshrink"
)

func main() {
	fmt.Println(imageshrink.ShrinkFile("cmd/input.png"))
}
