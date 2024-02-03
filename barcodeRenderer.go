package main

import (
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"strconv"
)

type SVG struct {
	XMLName     xml.Name `xml:"svg"`
	Version     string   `xml:"version,attr"`
	BaseProfile string   `xml:"baseProfile,attr"`
	Width       string   `xml:"width,attr"`
	Height      string   `xml:"height,attr"`
	Xmlns       string   `xml:"xmlns,attr"`
	GZone       GZone    `xml:"g"`
}

type GZone struct {
	XMLName xml.Name `xml:"g"`
	Stroke  string   `xml:"stroke,attr"`
	Lines   []Line   `xml:"line"`
}

type Line struct {
	XMLName     xml.Name `xml:"line"`
	StrokeWidth string   `xml:"stroke-width,attr"`
	Y1          string   `xml:"y1,attr"`
	X1          string   `xml:"x1,attr"`
	Y2          string   `xml:"y2,attr"`
	X2          string   `xml:"x2,attr"`
}

func saveAsSvg(barcodeValue string, filePath string) {
	rootLevel := &SVG{Version: "1.1", BaseProfile: "full", Width: "700", Height: "200", Xmlns: "http://www.w3.org/2000/svg"}
	rootLevel.GZone.Stroke = "black"

	for i, v := range barcodeValue {
		if string(v) == "1" {
			rootLevel.GZone.Lines = append(rootLevel.GZone.Lines, Line{StrokeWidth: "4", Y1: "10", X1: strconv.FormatInt(int64((i+1)*4+10), 10), Y2: "60", X2: strconv.FormatInt(int64((i+1)*4+10), 10)})
		}

	}

	filename := filePath
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(rootLevel); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

func saveAsPng(eanValue string) {

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

	out, err := os.Create("./ean.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgRect := image.Rect(0, 0, 300, 100)
	img := image.NewGray(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for indexWr, value := range barcodeValue {
		fill := &image.Uniform{color.White}
		if string(value) == "1" {
			fill = &image.Uniform{color.Black}
		}
		draw.Draw(img, image.Rect(int((indexWr+1)*4), 10, int((indexWr+2)*4), 60), fill, image.ZP, draw.Src)
	}

	// ok, write out the data into the new PNG file

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to ean.png")

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
