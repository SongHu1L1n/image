package app

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"net/http"
	"strings"

	"github.com/gogf/gf/os/gfile"
)

const (
	ImageJpeg = 1
	ImagePng  = 2
	ImageBmp  = 3
	ImageGif  = 4
	ImageWebp = 5
)

type Image struct {
	ImagePath string
	ImageType int
	Ext       string
	Data      []byte
	Width     int
	Height    int
}

func (img *Image) Open(path string) error {
	img.ImagePath = path
	img.Data = gfile.GetBytes(path)
	img.Ext = gfile.ExtName(path)
	contentType := http.DetectContentType(img.Data)
	var imgType int
	if strings.Contains(contentType, "jepg") {
		imgType = ImageJpeg
	} else if strings.Contains(contentType, "png") {
		imgType = ImagePng
	} else if strings.Contains(contentType, "bmp") {
		imgType = ImageBmp
	} else if strings.Contains(contentType, "gif") {
		imgType = ImageGif
	} else if strings.Contains(contentType, "webp") {
		imgType = ImageWebp
	}
	img.ImageType = imgType
	reader := bytes.NewReader(img.Data)
	imge, _, err := image.Decode(reader)
	if err != nil {
		log.Printf("fail to decode img, img type is %d, err : %v", img.ImageType, err)
		return fmt.Errorf("fail to decode img, img type is %d, err : %v", img.ImageType, err)
	}
	bound := imge.Bounds()
	img.Width = bound.Max.X
	img.Height = bound.Max.Y
	return nil
}

func (img *Image) Reset() {
	img.Data = nil
}
