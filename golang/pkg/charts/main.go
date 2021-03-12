package main

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func main() {
	pie := chart.PieChart{
		Title: "Pizza",
		TitleStyle: chart.Style{
			Show:      false,
			FontSize:  25,
			FontColor: drawing.ColorFromHex("F0047F"),
		},
		Background: chart.Style{
			Show:      true,
			Padding:   chart.NewBox(100, 100, 100, 100),
			FillColor: drawing.ColorTransparent,
		},
		Canvas: chart.Style{
			Show:      true,
			FillColor: drawing.ColorTransparent,
		},
		Width:  800,
		Height: 800,
		Values: []chart.Value{
			{Label: "Bugs", Value: 120, Style: BugsPieStyle},
			{Label: "Melhorias", Value: 25, Style: MelhoriasPieStyle},
			{Label: "Apoios", Value: 25, Style: ApoiosPieStyle},
		},
	}

	buffer := bytes.NewBuffer([]byte{})

	err := pie.Render(chart.PNG, buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(buffer.Bytes()))
}

var (
	BugsPieStyle      = chart.Style{Show: true, FillColor: drawing.ColorFromHex("eba00c"), FontSize: 28, FontColor: drawing.ColorWhite}
	MelhoriasPieStyle = chart.Style{Show: true, FillColor: drawing.ColorFromHex("0c8eeb"), FontSize: 28, FontColor: drawing.ColorWhite}
	ApoiosPieStyle    = chart.Style{Show: true, FillColor: drawing.ColorFromHex("1fab2d"), FontSize: 28, FontColor: drawing.ColorWhite}
)
