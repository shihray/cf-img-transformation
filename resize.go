package transform

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"

	"github.com/nfnt/resize"
)

const ResizePath = "resize"

func ResizeImg(ctx context.Context, s *storage.Client, e GCSEvent) (err error) {

	outputName := convertToResizePath(e.Name)
	if outputName == "" {
		return nil
	}

	inputBlob := s.Bucket(e.Bucket).Object(e.Name)
	r, err := inputBlob.NewReader(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	outputBlob := s.Bucket(e.Bucket).Object(outputName)
	w := outputBlob.NewWriter(ctx)
	defer w.Close()

	var (
		img   image.Image
		width uint = 400
	)
	if getType(e.Name) == "Banner" {
		width = uint(2500)
	}

	switch e.ContentType {
	case "image/png":
		img, err = png.Decode(r)
		if err != nil {
			log.Errorf(err.Error())
			return err
		}
	case "image/jpeg":
		img, err = jpeg.Decode(r)
		if err != nil {
			log.Errorf(err.Error())
			return err
		}
	}

	newImage := resize.Resize(width, 0, img, resize.Lanczos3)
	return jpeg.Encode(w, newImage, &jpeg.Options{75})
}

// name: 7000126/ModelCard/0fc78056-5c6b-4618-9e36-8948bfb6ea2c.gif
func convertToResizePath(name string) string {
	strArr := strings.Split(name, "/")

	if len(strArr) != 3 {
		log.Error("size error")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s", strArr[0], ResizePath, strArr[1], strArr[2])
}

func getType(name string) string {
	strArr := strings.Split(name, "/")

	if len(strArr) != 3 {
		log.Error("size error")
		return ""
	}
	return strArr[1]
}
