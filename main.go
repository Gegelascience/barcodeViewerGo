package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(isCorrectEan("12345670"))

	myApp := app.New()
	myWindow := myApp.NewWindow("BarcodeViewver")

	// Define a welcome text centered
	text := canvas.NewText("12345670", color.Black)
	text.Alignment = fyne.TextAlignCenter

	text2 := canvas.NewText("3666154117284", color.Black)
	text2.Alignment = fyne.TextAlignCenter

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		drawBarcode(("12345670")),
		//text2,
		//drawBarcode(("3666154117284")),
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

}

func drawBarcode(eanValue string) *fyne.Container {

	barcodeValue := ""
	//barcodeContainer := container.NewHBox()
	barcodeContainer := container.NewMax()

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

	/*for _, v := range barcodeValue {
		switch string(v) {
		case "1":
			line1 := canvas.NewRectangle(color.Black)
			line1.SetMinSize(fyne.NewSize(10, 50))
			barcodeContainer.Add(line1)
		case "0":
			line1 := canvas.NewRectangle(color.White)
			line1.SetMinSize(fyne.NewSize(10, 50))
			barcodeContainer.Add(line1)
		}
	}*/

	return barcodeContainer
}

func calculateEan13(eanValue string) string {
	barcodeValue := "101"

	firstPartRaw := string(eanValue[1:7])
	secondPartRaw := string(eanValue[7:13])

	prefix := string(eanValue[0])

	for index, element := range firstPartRaw {
		setToApply := calculateSetFromPrefix(prefix, index)
		if setToApply == "A" {
			barcodeValue = barcodeValue + mapSetA(string(element))
		} else {
			barcodeValue = barcodeValue + mapSetB(string(element))
		}

	}

	barcodeValue = barcodeValue + "01010"

	for _, element := range secondPartRaw {
		barcodeValue = barcodeValue + mapSetC(string(element))
	}

	barcodeValue = barcodeValue + "101"

	return barcodeValue
}

func calculateEan8(eanValue string) string {
	barcodeValue := "101"

	firstPartRaw := string(eanValue[0:4])
	secondPartRaw := string(eanValue[4:8])

	for _, element := range firstPartRaw {

		barcodeValue = barcodeValue + mapSetA(string(element))

	}

	barcodeValue = barcodeValue + "01010"

	for _, element := range secondPartRaw {
		barcodeValue = barcodeValue + mapSetC(string(element))
	}

	barcodeValue = barcodeValue + "101"

	return barcodeValue
}

func mapSetA(rawCharacter string) string {
	switch rawCharacter {
	case "0":
		return "0001101"
	case "1":
		return "0011001"
	case "2":
		return "0010011"
	case "3":
		return "0111101"
	case "4":
		return "0100011"
	case "5":
		return "0110001"
	case "6":
		return "0101111"
	case "7":
		return "0111011"
	case "8":
		return "0110111"
	case "9":
		return "0001011"
	default:
		return ""

	}
}

func mapSetB(rawCharacter string) string {
	switch rawCharacter {
	case "0":
		return "0100111"
	case "1":
		return "0110011"
	case "2":
		return "0011011"
	case "3":
		return "0100001"
	case "4":
		return "0011101"
	case "5":
		return "0111001"
	case "6":
		return "0000101"
	case "7":
		return "0010001"
	case "8":
		return "0001001"
	case "9":
		return "0010111"
	default:
		return ""

	}
}

func mapSetC(rawCharacter string) string {
	switch rawCharacter {
	case "0":
		return "1110010"
	case "1":
		return "1100110"
	case "2":
		return "1101100"
	case "3":
		return "1000010"
	case "4":
		return "1011100"
	case "5":
		return "1001110"
	case "6":
		return "1010000"
	case "7":
		return "1000100"
	case "8":
		return "1001000"
	case "9":
		return "1110100"
	default:
		return ""

	}
}

func calculateSetFromPrefix(prefix string, index int) string {
	/*
		Found odd set (set A) or even set (set B) by prefix value
	*/

	if index == 0 {
		return "A"
	}
	switch prefix {
	case "0":
		return "A"
	case "1":
		if index == 1 || index == 3 {
			return "A"
		}
		return "B"
	case "2":
		if index == 1 || index == 4 {
			return "A"
		}
		return "B"

	case "3":
		if index == 1 || index == 5 {
			return "A"
		}
		return "B"

	case "4":
		if index == 2 || index == 3 {
			return "A"
		}
		return "B"

	case "5":
		if index == 3 || index == 4 {
			return "A"
		}
		return "B"

	case "6":
		if index == 4 || index == 5 {
			return "A"
		}
		return "B"

	case "7":
		if index == 2 || index == 4 {
			return "A"
		}
		return "B"

	case "8":
		if index == 2 || index == 5 {
			return "A"
		}
		return "B"

	case "9":
		if index == 3 || index == 5 {
			return "A"
		}
		return "B"
	default:
		return ""

	}

}
