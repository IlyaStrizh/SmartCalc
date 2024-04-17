package calc

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

const (
	plotWidth  = 400 // Ширина графика в пикселях
	plotHeight = 400 // Высота графика в пикселях
)

func (v *View) plot() {
	if v.numberLabel.Text != "0" {
		v.writeHistory()
	}

	myWindow := fyne.CurrentApp().NewWindow("Plot Window")

	xEntryMin, xEntryMax := widget.NewEntry(), widget.NewEntry()
	xEntryMin.SetText("-50")
	xEntryMax.SetText("100")
	yEntryMin, yEntryMax := widget.NewEntry(), widget.NewEntry()
	yEntryMin.SetText("-50")
	yEntryMax.SetText("100")
	xLabelMin := widget.NewLabel("xMin:")
	xLabelMax := widget.NewLabel("xMax:")
	yLabelMin := widget.NewLabel("yMin:")
	yLabelMax := widget.NewLabel("yMax:")

	plotCanvas := v.initPlotter(xEntryMin, xEntryMax, yEntryMin, yEntryMax)

	plotContent := container.New(layout.NewCenterLayout(), plotCanvas)

	refreshButton := widget.NewButton("Refresh Plot", func() {
		newCanvas := v.initPlotter(xEntryMin, xEntryMax, yEntryMin, yEntryMax)

		plotContent.Add(newCanvas)
		myWindow.Canvas().Refresh(plotContent)

	})

	refreshButton.Importance = widget.HighImportance

	xRow := container.NewHBox(
		xLabelMin,
		xEntryMin,
		xLabelMax,
		xEntryMax,
	)
	yRow := container.NewHBox(
		yLabelMin,
		yEntryMin,
		yLabelMax,
		yEntryMax,
	)

	hBox := container.NewHBox(
		xRow,
		yRow,
		refreshButton,
	)

	vBox := container.NewVBox(
		plotContent,
		hBox,
	)

	myWindow.SetContent(vBox)

	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.SetFixedSize(true)
	myWindow.Show()
}

func (v *View) initPlotter(xEntryMin, xEntryMax, yEntryMin, yEntryMax *widget.Entry) fyne.CanvasObject {
	p := plot.New()
	hGrid := plotter.NewGrid()
	//hGrid.Horizontal.Width = vg.Length(1)
	p.Add(hGrid)
	vGrid := plotter.NewGrid()
	//vGrid.Vertical.Width = vg.Length(1)
	p.Add(vGrid)
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	minValueX, _ := strconv.ParseFloat(xEntryMin.Text, 64)
	maxValueX, _ := strconv.ParseFloat(xEntryMax.Text, 64)
	minValueY, _ := strconv.ParseFloat(yEntryMin.Text, 64)
	maxValueY, _ := strconv.ParseFloat(yEntryMax.Text, 64)

	p.X.Min = minValueX
	p.X.Max = maxValueX
	p.Y.Min = minValueY
	p.Y.Max = maxValueY

	n := math.Min(p.X.Min, p.Y.Min)
	m := math.Max(p.X.Max, p.Y.Max)

	err := plotutil.AddLinePoints(p,
		"Formula", v.initPoints(n, m))
	//"Second", initPoints(10),
	//"Third", initPoints(10))
	if err != nil {
		panic(err)
	}

	newCanvas, err := plotToCanvas(p)
	if err != nil {
		panic(err)
	}

	return newCanvas
}

func (v *View) initPoints(n float64, m float64) plotter.XYs {
	len := 500
	pts := make(plotter.XYs, len)
	interval := (m - n) / float64(len)

	for i := 0; i < len; i++ {
		pts[i].X = n
		n += interval
		s := v.numberLabel.Text
		y, err := v.presenter.calcPlot(&s, fmt.Sprint(n))
		if err != nil {
			newPts := make(plotter.XYs, len)
			copy(newPts, pts[:i])
			copy(newPts[i+1:], pts[i:])
			pts = newPts
			continue
		}

		pts[i].Y = y
	}

	return pts
}

func plotToCanvas(p *plot.Plot) (fyne.CanvasObject, error) {
	// Создаем временный файл для сохранения графика
	f, err := os.CreateTemp("", "plot.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Сохраняем график в файл
	if err := p.Save(plotWidth, plotHeight, f.Name()+".png"); err != nil {
		return nil, err
	}

	// Загружаем изображение графика из файла
	img := canvas.NewImageFromFile(f.Name() + ".png")
	img.FillMode = canvas.ImageFillOriginal

	return img, nil
}
