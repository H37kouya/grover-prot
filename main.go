package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/H37kouya/grover-prot/helpers"
	"github.com/golang/freetype/truetype"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func init() {
	rand.Seed(int64(0))
	//import japanese fonts
	b, err := ioutil.ReadFile(`C:\WINDOWS\FONTS\BIZ-UDGOTHICR.TTC`)
	if err != nil {
		panic(err)
	}
	ft, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	fontName := "BIZ-UDGOTHICR"
	vg.AddFont(fontName, ft)

	//default font
	plot.DefaultFont = fontName
	plotter.DefaultFont = fontName
}

func main() {
	// csv読み込み
	f, err := os.Open("inputs/target.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	var labelX []string
	dataY := plotter.Values{}

	for i := 0; i < 200; i++ {
		record, err := r.Read()
		if i == 0 {
			i++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		labelX = append(labelX, record[0])
		floatVal, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			panic(err)
		}
		dataY = append(dataY, floatVal)
	}

	// グラフのプロット
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	bar, err := plotter.NewBarChart(dataY, vg.Points(20))
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Grover Alcoholism"
	p.X.Label.Text = "回数"
	p.Y.Label.Text = "出力値"
	p.Add(bar)

	// Save the plot to a PNG file.
	timeForFilename := helpers.GetTimeForFilename()
	if err := p.Save(12*vg.Inch, 12*vg.Inch, "outputs/"+timeForFilename+".png"); err != nil {
		panic(err)
	}
}
