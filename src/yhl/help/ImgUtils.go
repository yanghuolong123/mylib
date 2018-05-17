package help

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func GetImgHW(f string) (h, w int) {
	file, err := os.Open(f)
	if err != nil {
		Error(err)
		return
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		Error(err)
		return
	}

	w, h = img.Width, img.Height

	return
}
