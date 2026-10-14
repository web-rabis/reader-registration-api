package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"strings"

	"github.com/disintegration/imaging"
)

type ResizeMode string

const (
	ModeFit    ResizeMode = "fit"    // сохраняет пропорции, вписывая в рамку
	ModeResize ResizeMode = "resize" // жестко в указанный размер (без пропорций)
	ModeFill   ResizeMode = "fill"   // сохраняет пропорции + обрезает, чтобы было точно maxWidth x maxHeight
)

func ResizeImage(inputBase64 string, maxWidth, maxHeight int, mode ResizeMode, quality int) ([]byte, error) {
	commaIdx := strings.Index(inputBase64, ",")
	if commaIdx > 0 {
		inputBase64 = inputBase64[commaIdx+1:]
	}
	// Декодируем base64 в []byte
	data, err := base64.StdEncoding.DecodeString(inputBase64)
	if err != nil {
		return nil, err
	}

	// Декодируем картинку
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Проверяем размеры
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Если картинка меньше — возвращаем как есть
	if width <= maxWidth && height <= maxHeight {
		return data, nil
	}

	// Уменьшаем
	var resized image.Image
	switch mode {
	case ModeResize:
		resized = imaging.Resize(img, maxWidth, maxHeight, imaging.Lanczos)
	case ModeFill:
		resized = imaging.Fill(img, maxWidth, maxHeight, imaging.Center, imaging.Lanczos)
	default:
		resized = imaging.Fit(img, maxWidth, maxHeight, imaging.Lanczos)
	}

	// Кодируем обратно в JPEG (по умолчанию)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, resized, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
