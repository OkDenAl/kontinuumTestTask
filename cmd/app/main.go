package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"kontinuumTest/pkg"
	"strings"
)

func main() {
	a := app.New()
	window := a.NewWindow("Рассчет баллов")
	window.Resize(fyne.NewSize(400, 320))
	label := widget.NewLabel(
		"Введите даты, за которые ты хочешь получить результат\n" +
			"\t(даты нужно писать в СТРОГО в формате ГГГГ.ММ.ДД) \n\n" +
			"Опционально можно указать конкретные id учеников через пробел ")
	answer := widget.NewLabel("")
	entry := widget.NewEntry()
	btn := widget.NewButton("Рассчитать данные", func() {
		answer.SetText("")
		data := strings.Split(entry.Text, " ")
		info := pkg.SqlHandler(data)
		for _, text := range info {
			answer.SetText(answer.Text + text + "\n")
		}
	})
	close1 := widget.NewButton("Закрыть", func() {
		a.Quit()
	})
	window.SetContent(container.NewVScroll(container.NewVBox(
		label,
		entry,
		btn,
		answer,
		close1,
	)))
	window.ShowAndRun()
}
