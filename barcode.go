package barcode

import (
	"image"
	"image/color"
)

type CodeItemStruct struct {
	IsLine bool
	Width  int
}
type CodesStruct struct {
	Codes []CodeItemStruct
	Width int
}

func Encode(code *CodesStruct, size, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, code.Width*size, height))
	offset := 0
	for _, v := range code.Codes {
		for y := 0; y < height; y++ {
			for x := 0; x < v.Width*size; x++ {
				if v.IsLine {
					img.Set(x+offset, y, color.Black)
				} else {
					img.Set(x+offset, y, color.White)
				}
			}
		}
		offset = offset + v.Width*size
	}
	return img
}
