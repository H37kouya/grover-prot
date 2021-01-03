package main

import (
	"image/color"
	"io/ioutil"
	"math/rand"

	"gonum.org/v1/plot/plotutil"

	"github.com/H37kouya/grover-prot/helpers"
	"github.com/golang/freetype/truetype"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type xyAxis struct {
	x float64
	y float64
}

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

	xyAxises1, err := helpers.ReadCsv("inputs/result_1.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises2, err := helpers.ReadCsv("inputs/result_2.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises3, err := helpers.ReadCsv("inputs/result_3.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises4, err := helpers.ReadCsv("inputs/result_4.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises5, err := helpers.ReadCsv("inputs/result_5.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises6, err := helpers.ReadCsv("inputs/result_6.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises7, err := helpers.ReadCsv("inputs/result_7.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises8, err := helpers.ReadCsv("inputs/result_8.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises9, err := helpers.ReadCsv("inputs/result_9.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises10, err := helpers.ReadCsv("inputs/result_10.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises11, err := helpers.ReadCsv("inputs/result_11.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises12, err := helpers.ReadCsv("inputs/result_12.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises13, err := helpers.ReadCsv("inputs/result_13.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises14, err := helpers.ReadCsv("inputs/result_14.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises15, err := helpers.ReadCsv("inputs/result_15.csv", 200)
	if err != nil {
		panic(err)
	}
	xyAxises16, err := helpers.ReadCsv("inputs/result_16.csv", 200)
	if err != nil {
		panic(err)
	}

	pts1 := make(plotter.XYs, len(xyAxises1))
	for i, axis := range xyAxises1 {
		pts1[i].X = axis.X
		pts1[i].Y = axis.Y
	}
	pts2 := make(plotter.XYs, len(xyAxises2))
	for i, axis := range xyAxises2 {
		pts2[i].X = axis.X
		pts2[i].Y = axis.Y
	}
	pts3 := make(plotter.XYs, len(xyAxises3))
	for i, axis := range xyAxises3 {
		pts3[i].X = axis.X
		pts3[i].Y = axis.Y
	}
	pts4 := make(plotter.XYs, len(xyAxises4))
	for i, axis := range xyAxises4 {
		pts4[i].X = axis.X
		pts4[i].Y = axis.Y
	}
	pts5 := make(plotter.XYs, len(xyAxises5))
	for i, axis := range xyAxises5 {
		pts5[i].X = axis.X
		pts5[i].Y = axis.Y
	}
	pts6 := make(plotter.XYs, len(xyAxises6))
	for i, axis := range xyAxises6 {
		pts6[i].X = axis.X
		pts6[i].Y = axis.Y
	}
	pts7 := make(plotter.XYs, len(xyAxises7))
	for i, axis := range xyAxises7 {
		pts7[i].X = axis.X
		pts7[i].Y = axis.Y
	}
	pts8 := make(plotter.XYs, len(xyAxises8))
	for i, axis := range xyAxises8 {
		pts8[i].X = axis.X
		pts8[i].Y = axis.Y
	}
	pts9 := make(plotter.XYs, len(xyAxises9))
	for i, axis := range xyAxises9 {
		pts9[i].X = axis.X
		pts9[i].Y = axis.Y
	}
	pts10 := make(plotter.XYs, len(xyAxises10))
	for i, axis := range xyAxises10 {
		pts10[i].X = axis.X
		pts10[i].Y = axis.Y
	}
	pts11 := make(plotter.XYs, len(xyAxises11))
	for i, axis := range xyAxises11 {
		pts11[i].X = axis.X
		pts11[i].Y = axis.Y
	}
	pts12 := make(plotter.XYs, len(xyAxises12))
	for i, axis := range xyAxises12 {
		pts12[i].X = axis.X
		pts12[i].Y = axis.Y
	}
	pts13 := make(plotter.XYs, len(xyAxises13))
	for i, axis := range xyAxises13 {
		pts13[i].X = axis.X
		pts13[i].Y = axis.Y
	}
	pts14 := make(plotter.XYs, len(xyAxises14))
	for i, axis := range xyAxises14 {
		pts14[i].X = axis.X
		pts14[i].Y = axis.Y
	}
	pts15 := make(plotter.XYs, len(xyAxises15))
	for i, axis := range xyAxises15 {
		pts15[i].X = axis.X
		pts15[i].Y = axis.Y
	}
	pts16 := make(plotter.XYs, len(xyAxises16))
	for i, axis := range xyAxises16 {
		pts16[i].X = axis.X
		pts16[i].Y = axis.Y
	}

	// グラフのプロット
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Grover Alcoholism"
	p.Title.Padding = vg.Length(10.0)
	p.Title.TextStyle.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.X.Label.Text = "回数"
	p.X.Padding = vg.Length(10.0)
	p.X.Label.Padding = vg.Length(10.0)
	p.X.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.X.LineStyle.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.X.Label.TextStyle.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.Y.Label.Text = "出力値"
	p.Y.Max = 1.0
	p.Y.Min = 0
	p.Y.Label.Padding = vg.Length(10.0)
	p.Y.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.Legend.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.Y.LineStyle.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.Y.Label.TextStyle.Color = color.RGBA{
		R: uint8(242),
		G: uint8(242),
		B: uint8(242),
		A: 0,
	}
	p.BackgroundColor = color.RGBA{
		R: uint8(49),
		G: uint8(49),
		B: uint8(49),
		A: uint8(5),
	}
	err = plotutil.AddLinePoints(p, pts1, pts2, pts3, pts4, pts5, pts6, pts7, pts8, pts9, pts10, pts11, pts12, pts13, pts14, pts15, pts16)

	// Save the plot to a PNG file.
	timeForFilename := helpers.GetTimeForFilename()
	if err := p.Save(12*vg.Inch, 12*vg.Inch, "outputs/"+timeForFilename+".png"); err != nil {
		panic(err)
	}
}
