package calc

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type mrc struct {
	value string
	flag  bool
}

type View struct {
	myWindow    fyne.Window
	numberLabel *widget.Label
	xLabel      *widget.Label
	label       *widget.Label
	presenter   *Presenter
	historyFile string
	memory      mrc
	counter     int
}

func NewCalcMenu(myApp fyne.App) *View {
	menu := View{
		myWindow:    myApp.NewWindow("SmartCalc_v3.0"),
		numberLabel: widget.NewLabel("0"),
		xLabel:      widget.NewLabel("0"),
		label:       widget.NewLabel("x:"),
		historyFile: "../Resources/history.txt",
		memory:      mrc{value: "0"},
		counter:     1,
	}

	menu.label.TextStyle = fyne.TextStyle{Bold: true}
	menu.xLabel.TextStyle = fyne.TextStyle{Italic: true}
	menu.numberLabel.TextStyle = fyne.TextStyle{Bold: false, Italic: true}
	menu.numberLabel.Alignment = fyne.TextAlignTrailing
	menu.myWindow.SetContent(menu.makeMenuBox())
	menu.myWindow.Resize(fyne.NewSize(430, 250))
	menu.myWindow.SetFixedSize(true)
	menu.myWindow.Show()

	return &menu
}

func (v *View) InitPresenter(p *Presenter) {
	v.presenter = p
}

func (v *View) makeMenuBox() *fyne.Container {
	plotButton := widget.NewButton("Plot", v.plot)

	labelHBox := container.NewHBox(
		v.label,
		v.xLabel,
	)

	buttonHBox := container.NewHBox(
		v.makeVBox(),
		v.makeVBox1(),
		v.makeVBox2(),
		v.makeVBox3(),
		v.makeVBox4(),
		v.makeVBox5(),
		v.makeVBox6(),
		v.makeVBox7(),
		plotButton,
	)

	scrollLabelHBox := container.NewHScroll(labelHBox)
	scrollLabelHBox.SetMinSize(fyne.NewSize(350, 40))
	scrollNumberLabel := container.NewHScroll(v.numberLabel)
	scrollNumberLabel.SetMinSize(fyne.NewSize(350, 40))

	Box := container.NewVBox(
		scrollNumberLabel,
		scrollLabelHBox,
		buttonHBox,
	)

	return Box
}

func (v *View) makeVBox() *fyne.Container {
	infoButton := widget.NewButtonWithIcon("", theme.HelpIcon(), v.info)
	historyButton := widget.NewButtonWithIcon("", theme.DocumentIcon(), v.history)
	mrcButton := widget.NewButton("MRC", v.mrcButton)
	mMinusButton := widget.NewButton("M-", v.mMinusButton)
	mPlusButton := widget.NewButton("M+", v.mPlusButton)

	buttonVBox := container.NewVBox(
		infoButton,
		historyButton,
		mrcButton,
		mMinusButton,
		mPlusButton,
	)

	return buttonVBox
}

func (v *View) makeVBox1() *fyne.Container {
	lnButton := widget.NewButton("ln", func() { v.printButton("ln(") })
	sqrtButton := widget.NewButton("sqrt", func() { v.printButton("sqrt(") })
	atanButton := widget.NewButton("atan", func() { v.printButton("atan(") })
	acosButton := widget.NewButton("acos", func() { v.printButton("acos(") })
	asinButton := widget.NewButton("asin", func() { v.printButton("asin(") })

	buttonVBox1 := container.NewVBox(
		lnButton,
		sqrtButton,
		atanButton,
		acosButton,
		asinButton,
	)

	return buttonVBox1
}

func (v *View) makeVBox2() *fyne.Container {
	logButton := widget.NewButton(" log  ", func() { v.printButton("log(") })
	tanButton := widget.NewButton("tan", func() { v.printButton("tan(") })
	cosButton := widget.NewButton("cos", func() { v.printButton("cos(") })
	sinButton := widget.NewButton("sin", func() { v.printButton("sin(") })
	inverseButton := widget.NewButton("+/-", v.inverseButton)

	buttonVBox2 := container.NewVBox(
		logButton,
		tanButton,
		cosButton,
		sinButton,
		inverseButton,
	)

	return buttonVBox2
}

func (v *View) makeVBox3() *fyne.Container {
	nullButton := widget.NewButton("0", func() { v.printButton("0") })
	oneButton := widget.NewButton("1", func() { v.printButton("1") })
	fourButton := widget.NewButton("4", func() { v.printButton("4") })
	sevenButton := widget.NewButton("7", func() { v.printButton("7") })
	powButton := widget.NewButton("   ^   ", func() { v.printOperator("^") })

	buttonVBox3 := container.NewVBox(
		powButton,
		sevenButton,
		fourButton,
		oneButton,
		nullButton,
	)

	return buttonVBox3
}

func (v *View) makeVBox4() *fyne.Container {
	twoButton := widget.NewButton("2", func() { v.printButton("2") })
	fiveButton := widget.NewButton("5", func() { v.printButton("5") })
	eightButton := widget.NewButton("8", func() { v.printButton("8") })
	pointerButton := widget.NewButton(".", v.pointerButton)
	delButton := widget.NewButton("   <-  ", v.delButton)

	buttonVBox4 := container.NewVBox(
		delButton,
		eightButton,
		fiveButton,
		twoButton,
		pointerButton,
	)

	return buttonVBox4
}

func (v *View) makeVBox5() *fyne.Container {
	threeButton := widget.NewButton("3", func() { v.printButton("3") })
	sixButton := widget.NewButton("6", func() { v.printButton("6") })
	nineButton := widget.NewButton("9", func() { v.printButton("9") })
	eButton := widget.NewButton("e", func() { v.printOperator("e") })
	acButton := widget.NewButton("  AC ", v.acButton)

	buttonVBox5 := container.NewVBox(
		acButton,
		nineButton,
		sixButton,
		threeButton,
		eButton,
	)

	return buttonVBox5
}

func (v *View) makeVBox6() *fyne.Container {
	divButton := widget.NewButton("   /   ", func() { v.printOperator("/") })
	divButton.Importance = widget.WarningImportance

	multButton := widget.NewButton("*", func() { v.printOperator("*") })
	multButton.Importance = widget.WarningImportance

	minusButton := widget.NewButton("-", func() { v.printButton("-") })
	minusButton.Importance = widget.WarningImportance

	plusButton := widget.NewButton("+", func() { v.printOperator("+") })
	plusButton.Importance = widget.WarningImportance

	evaluateButton := widget.NewButton("=", v.evaluateButton)
	evaluateButton.Importance = widget.WarningImportance

	buttonVBox6 := container.NewVBox(
		divButton,
		multButton,
		minusButton,
		plusButton,
		evaluateButton,
	)

	return buttonVBox6
}

func (v *View) makeVBox7() *fyne.Container {
	modButton := widget.NewButton("  %  ", func() { v.printOperator("%") })
	modButton.Importance = widget.HighImportance

	bracketButton := widget.NewButton("(", func() { v.printButton("(") })
	bracketButton.Importance = widget.HighImportance

	rBracketButton := widget.NewButton(")", func() { v.printButton(")") })
	rBracketButton.Importance = widget.HighImportance

	xInitButton := widget.NewButton("x<-", v.xInitButton)
	xInitButton.Importance = widget.HighImportance

	xButton := widget.NewButton("x", v.xButton)
	xButton.Importance = widget.HighImportance

	buttonVBox7 := container.NewVBox(
		modButton,
		bracketButton,
		rBracketButton,
		xInitButton,
		xButton,
	)

	return buttonVBox7
}

func (v *View) writeHistory() {
	file, err := os.OpenFile(v.historyFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n", v.numberLabel.Text))
	if err != nil {
		log.Println("Ошибка при записи в файл:", err)
		return
	}
}
