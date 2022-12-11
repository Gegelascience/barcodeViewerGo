package main

import (
	"encoding/xml"
	"fmt"
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
			rootLevel.GZone.Lines = append(rootLevel.GZone.Lines, Line{StrokeWidth: "10", Y1: "10", X1: strconv.FormatInt(int64((i+1)*10), 10), Y2: "60", X2: strconv.FormatInt(int64((i+1)*10), 10)})
		}

	}

	filename := "ean.svg"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(rootLevel); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
