package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func createCodeEditor() *tview.TextArea {
	codeEditor := tview.NewTextArea().
		SetPlaceholder("Enter your code here...")

	codeEditor.SetTitle("Code Editor").SetBorder(true)
	codeEditor.SetBorderPadding(1, 1, 2, 2)

	return codeEditor
}

func createPositionView() *tview.TextView {
	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	return position
}

func setupCodeEditorEvents(codeEditor *tview.TextArea, position *tview.TextView) {
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
}
