package calc

import (
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) history() {
	myWindow := fyne.CurrentApp().NewWindow("History")

	v.showHistory(myWindow)
}

func (v *View) showHistory(myWindow fyne.Window) {

	content, err := os.ReadFile(v.historyFile)
	if err != nil {
		log.Println("Ошибка при чтении файла history.txt:", err)
		return
	}

	lines := strings.Fields(string(content))
	if len(lines) > 12 {
		lines = append(lines, "")
	}
	list := widget.NewList(
		func() int {
			return len(lines)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Button).SetText(lines[id])
			obj.(*widget.Button).OnTapped = func() {
				v.updateNumberLabel(lines[id])
				myWindow.Close()
			}
		},
	)

	scrollView := container.NewScroll(list)
	scrollView.SetMinSize(fyne.NewSize(600, 480))

	cleanHistoryButton := widget.NewButton("Clean History", func() {
		v.cleanHistoryButton(myWindow)
	})
	cleanHistoryButton.Importance = widget.HighImportance

	vBox := container.NewVBox(
		scrollView,
		cleanHistoryButton,
	)

	myWindow.SetContent(vBox)
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Show()
}

func (v *View) cleanHistoryButton(myWindow fyne.Window) {
	file, err := os.Create(v.historyFile)
	if err != nil {
		log.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	v.showHistory(myWindow)
}
