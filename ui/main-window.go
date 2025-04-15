package ui

import (
	"github.com/rivo/tview"
)

func InitMainWindow() *tview.Application {
	app := tview.NewApplication()

	pages := tview.NewPages()

	codeEditor := createCodeEditor()
	position := createPositionView()
	rulesList := createRulesListView()
	factsList := createFactsListView()

	addRuleButton := createAddRuleButton(app, rulesList, pages)
	addFactButton := createAddFactButton(app, factsList, pages)

	setupCodeEditorEvents(codeEditor, position)

	flex := setupLayout(codeEditor, position, rulesList, addRuleButton, addFactButton, factsList)

	pages.AddPage("main", flex, true, true)

	return app.SetRoot(pages, true).EnableMouse(true)
}

func setupLayout(
	codeEditor *tview.TextArea,
	position *tview.TextView,
	rulesList *tview.List,
	addRuleButton *tview.Button,
	addFactButton *tview.Button,
	factsList *tview.List) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(codeEditor, 0, 1, true).
			AddItem(position, 1, 0, false), 0, 2, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("TODO"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Answers"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Question"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(factsList, 0, 15, false).
			AddItem(addFactButton, 5, 1, false).
			AddItem(rulesList, 0, 15, false).
			AddItem(addRuleButton, 5, 1, false), 0, 2, false)

	return flex
}
