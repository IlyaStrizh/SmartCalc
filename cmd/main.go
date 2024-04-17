package main

import (
	calc "SmartCalc_v3.0/internal"
	"SmartCalc_v3.0/internal/model"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	view := calc.NewCalcMenu(myApp)
	presenter := calc.NewPresenter(
		view, model.NewModel("../Resources/calc.so"))
	view.InitPresenter(presenter)

	myApp.Run()
}
