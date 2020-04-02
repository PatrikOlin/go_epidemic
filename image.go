package main

import (
	// "bytes"							
	// "encoding/base64"
	"fmt"
	"image"
	"image/png"
	// "log"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
	 "github.com/BurntSushi/xgbutil"		
	// "github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"

)

// draw the cells
func draw(w int, h int, cells []Cell) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(dest)
	for _, cell := range cells {
		gc.SetFillColor(cell.Color)
		gc.MoveTo(float64(cell.X), float64(cell.Y))
		gc.ArcTo(float64(cell.X), float64(cell.Y),
			float64(cell.R/2), float64(cell.R/2), 0, 6.283185307179586)
		gc.Close()
		gc.Fill()
	}
	return dest
}

// Print the image to iTerm2 terminal
func printImage(X *xgbutil.XUtil, img image.Image) {
	ximg := xgraphics.NewConvert(X, img)

	ximg.XShowExtra("PANDEMIC!", true)

	// xevent.Main(X)						
	// var buf bytes.Buffer
	// png.Encode(&buf, img)
	// imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	// fmt.Printf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a", imgBase64Str)
}

// save the image
func saveImage(filePath string, rgba *image.RGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot create file:", err)
	}

	png.Encode(imgFile, rgba.SubImage(rgba.Rect))
}
