package calc

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (v *View) info() {
	myWindow := fyne.CurrentApp().NewWindow("Info")

	content, err := os.ReadFile("../Resources/info.md")
	if err != nil {
		fmt.Println("Ошибка при чтении файла info.md:", err)
		return
	}

	markdownContent := widget.NewRichTextFromMarkdown(string(content))

	scrollView := container.NewScroll(markdownContent)

	myWindow.SetContent(scrollView)
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.CenterOnScreen()
	myWindow.Show()
}
