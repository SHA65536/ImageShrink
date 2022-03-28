# ImageShrink
A small Go tool to crop transparent padding from an image according to content.

## Install
Use `go get` to install this package
```
go get github.com/sha65536/imageshrink
```
## Example
This example takes an image file, shrinks it, and overwrites the file with the new image.
```go
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
```