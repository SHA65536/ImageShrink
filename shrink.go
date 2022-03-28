package imageshrink

import (
	"errors"
	"image"
	"image/png"
	"os"
)

type OpaqueImg interface {
	Opaque() bool
}

func ShrinkImg(img image.Image) (image.Image, error) {
	var ctop, cbot, cleft, cright int
	if oim, ok := img.(OpaqueImg); !ok || oim.Opaque() {
		return nil, errors.New("image: image has no alpha")
	}
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y && ctop == 0; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0x0000 {
				ctop = y
				break
			}
		}
	}
	for x := rect.Min.X; x < rect.Max.X && cleft == 0; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0x0000 {
				cleft = x
				break
			}
		}
	}
	for y := rect.Max.Y - 1; y > rect.Min.Y && cbot == 0; y-- {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0x0000 {
				cbot = y
				break
			}
		}
	}
	for x := rect.Max.X - 1; x > rect.Min.X && cright == 0; x-- {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0x0000 {
				cright = x
				break
			}
		}
	}
	cropRect := image.Rect(cleft, ctop, cright+1, cbot+1)
	newImg := image.NewRGBA(cropRect)
	for y := cropRect.Min.Y; y < cropRect.Max.Y; y++ {
		for x := cropRect.Min.X; x < cropRect.Max.X; x++ {
			newImg.Set(x, y, img.At(x, y))
		}
	}
	return newImg, nil
}

func ShrinkFile(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(fd)
	fd.Close()
	if err != nil {
		return err
	}
	cropped, err := ShrinkImg(img)
	if err != nil {
		return err
	}
	fd, err = os.Create(path)
	if err != nil {
		return err
	}
	err = png.Encode(fd, cropped)
	return err
}
