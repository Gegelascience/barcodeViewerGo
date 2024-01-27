package main

import (
	"fmt"
	"image/color"
	"runtime/debug"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	var eanValue string
	fmt.Println("Please, write ean value you want to see")
	fmt.Scanln(&eanValue)
	// example values: 12345670 3666154117284

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			fmt.Printf("%v, %s", panicInfo, string(debug.Stack()))
		}
	}()

	if isCorrectEan(eanValue) {
		saveAsPng(eanValue)
		myApp := app.New()
		myWindow := myApp.NewWindow("BarcodeViewver")

		// Define a welcome text centered
		text := canvas.NewText(eanValue, color.Black)
		text.Alignment = fyne.TextAlignCenter

		// Display a vertical box containing text, image and button
		box := container.NewVBox(
			text,
			drawBarcode((eanValue)),
		)

		// Display our content
		myWindow.SetContent(box)

		// Close the App when Escape key is pressed
		myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

			if keyEvent.Name == fyne.KeyEscape {
				myApp.Quit()
			}
		})

		// Show window and run app
		myWindow.ShowAndRun()

	} else {
		fmt.Println("Invalid EAN")
	}

}

func drawBarcode(eanValue string) *canvas.Image {

	barcodeValue := ""

	switch len(eanValue) {
	case 8:
		barcodeValue = calculateEan8(eanValue)
		fmt.Println(barcodeValue)

	case 13:
		barcodeValue = calculateEan13(eanValue)

	default:
		panic("Invalid EAN")
	}

	saveAsSvg(barcodeValue, "ean.svg")

	imageEan := canvas.NewImageFromFile("ean.svg")
	imageEan.SetMinSize(fyne.NewSize(700, 200))
	imageEan.FillMode = canvas.ImageFillContain

	return imageEan
}
