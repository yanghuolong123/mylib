package help

import (
	qrcode "github.com/skip2/go-qrcode"
)

func QrImg(url string) string {

	png, _ := qrcode.Encode(url, qrcode.High, 400)
	return string(png)
}
