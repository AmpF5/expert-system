package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitUI() *tview.Application {
	app := tview.NewApplication()

	pages := tview.NewPages()

	codeEditor := createCodeEditor()
	position := createPositionView()
	rulesList := createRulesListView()

	addRuleButton := createAddRuleButton(app, rulesList, pages)

	setupCodeEditorEvents(codeEditor, position)

	flex := setupLayout(codeEditor, position, rulesList, addRuleButton)

	pages.AddPage("main", flex, true, true)

	return app.SetRoot(pages, true).EnableMouse(true)
}

func setupLayout(
	codeEditor *tview.TextArea,
	position *tview.TextView,
	rulesList *tview.List,
	addRuleButton *tview.Button) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(codeEditor, 0, 1, true).
			AddItem(position, 1, 0, false), 0, 2, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("TODO"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Answers"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Question"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(rulesList, 0, 15, false).
			AddItem(addRuleButton, 5, 1, false), 0, 2, false)

	return flex
}

func createRulesListView() *tview.List {
	rulesList := tview.NewList().
		ShowSecondaryText(false).
		SetHighlightFullLine(true).
		SetMainTextColor(tcell.ColorWhite).
		SetSelectedTextColor(tcell.ColorBlack).
		SetSelectedBackgroundColor(tcell.ColorGreen)

	rulesList.SetBorder(true).SetTitle("Rules")

	var originalTextColor tcell.Color = tcell.ColorBlack
	var originalBgColor tcell.Color = tcell.ColorGreen

	rulesList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
	})

	rulesList.SetBlurFunc(func() {
		originalTextColor = tcell.ColorBlack
		originalBgColor = tcell.ColorGreen
		rulesList.SetSelectedTextColor(tcell.ColorWhite) // Same as main text
		rulesList.SetSelectedBackgroundColor(tcell.ColorDefault)
	})

	rulesList.SetFocusFunc(func() {
		rulesList.SetSelectedTextColor(originalTextColor)
		rulesList.SetSelectedBackgroundColor(originalBgColor)
	})

	addRule(rulesList, "Rule 1: If temperature > 30°C then hot")
	addRule(rulesList, "Rule 2: If sky = cloudy AND humidity > 80% then rain")

	return rulesList
}

func createAddRuleButton(app *tview.Application, rulesList *tview.List, pages *tview.Pages) *tview.Button {
	b := tview.NewButton("+Add new rule").
		SetSelectedFunc(func() {
			form := createRuleForm(rulesList, pages)

			modal := tview.NewGrid().
				SetColumns(0, 40, 0).
				SetRows(0, 12, 0).
				AddItem(form, 1, 1, 1, 1, 0, 0, true)

			pages.AddPage("add_rule_form", modal, true, true)

			app.SetFocus(form)
		})

	b.SetBorder(true)
	backgroundColor := tcell.ColorCadetBlue
	textColor := tcell.ColorWhite

	b.SetBackgroundColor(backgroundColor)
	b.SetBackgroundColorActivated(backgroundColor)
	b.SetLabelColor(textColor)
	b.SetLabelColorActivated(textColor)

	return b
}

func createRuleForm(rulesList *tview.List, pages *tview.Pages) *tview.Form {
	form := tview.NewForm()

	form.AddInputField("Identifier", "", 20, nil, nil)
	form.AddInputField("Condition", "", 10, nil, nil)
	form.AddInputField("Value", "", 20, nil, nil)

	form.AddButton("Save", func() {
		conditionField := form.GetFormItemByLabel("Identifier").(*tview.InputField)
		operatorField := form.GetFormItemByLabel("Condition").(*tview.InputField)
		valueField := form.GetFormItemByLabel("Value").(*tview.InputField)

		identifier := conditionField.GetText()
		condition := operatorField.GetText()
		value := valueField.GetText()

		ruleText := fmt.Sprintf("%s %s %s", identifier, condition, value)

		AddRule(rulesList, ruleText)

		pages.RemovePage("add_rule_form")
	})

	form.AddButton("Cancel", func() {
		pages.RemovePage("add_rule_form")
	})

	form.SetBorder(true).
		SetTitle("Add New Rule").
		SetTitleAlign(tview.AlignCenter)

	form.SetButtonsAlign(tview.AlignCenter)

	return form
}

func AddRule(rulesList *tview.List, ruleText string) {
	rulesList.AddItem(ruleText, "", 0, nil)
}

func addRule(rulesList *tview.List, ruleText string) {
	rulesList.AddItem(ruleText, "", 0, nil)
}

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
