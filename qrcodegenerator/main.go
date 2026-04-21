package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	createQRCode("https://manuelarte.github.io/gophercamp-2026-create-your-first-linter", "qr-slides.jpeg")
	createQRCode("https://github.com/manuelarte", "qr-github-manuelarte.jpeg")
	createQRCode("https://github.com/manuelarte/gophercamp-2026-create-your-first-linter", "qr-github-manuelarte-gophercamp-2026-create-your-first-linter.jpeg")
}

func createQRCode(url string, filename string) {
	qrc, err := qrcode.New(url)
	if err != nil {
		logger.Error("could not generate QRCode", slog.Any("err", err))
		return
	}

	w, err := standard.New(fmt.Sprintf("../slides/public/%s", filename))
	if err != nil {
		logger.Error("standard.New failed", slog.Any("err", err))
		return
	}
	defer func(w *standard.Writer) {
		errC := w.Close()
		if errC != nil {
			logger.Error("w.Close failed", slog.Any("err", errC))
		}
	}(w)

	// save file
	if err = qrc.Save(w); err != nil {
		logger.Error("could not save image", slog.Any("err", err))
		return
	}
	logger.Info(
		"QR Code generated successfully",
		slog.String("url", url),
		slog.String("filename", filename),
	)
}
