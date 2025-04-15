package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	codeEditor := tview.NewTextArea().
		SetPlaceholder("Enter your code here...")

	codeEditor.SetTitle("Code Editor").SetBorder(true)

	codeEditor.SetBorderPadding(1, 1, 2, 2)

	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := codeEditor.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	codeEditor.SetMovedFunc(updateInfos)
	updateInfos()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(codeEditor, 0, 1, true).
			AddItem(position, 1, 0, false), 0, 2, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("TODO"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Answers"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Question"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Rules"), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
