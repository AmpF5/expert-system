package ui

import (
	"expert-system/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createFactsListView() *tview.List {
	factsList := tview.NewList().
		ShowSecondaryText(false).
		SetHighlightFullLine(true).
		SetMainTextColor(tcell.ColorWhite).
		SetSelectedTextColor(tcell.ColorBlack).
		SetSelectedBackgroundColor(tcell.ColorGreen)

	factsList.SetBorder(true).SetTitle("Facts")

	var originalTextColor tcell.Color = tcell.ColorBlack
	var originalBgColor tcell.Color = tcell.ColorGreen

	factsList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
	})

	factsList.SetBlurFunc(func() {
		originalTextColor = tcell.ColorBlack
		originalBgColor = tcell.ColorGreen
		factsList.SetSelectedTextColor(tcell.ColorWhite) // Same as main text
		factsList.SetSelectedBackgroundColor(tcell.ColorDefault)
	})

	factsList.SetFocusFunc(func() {
		factsList.SetSelectedTextColor(originalTextColor)
		factsList.SetSelectedBackgroundColor(originalBgColor)
	})

	return factsList
}

func createAddFactButton(app *tview.Application, factsList *tview.List, pages *tview.Pages) *tview.Button {
	b := tview.NewButton("+Add new fact").
		SetSelectedFunc(func() {
			form := createFactForm(factsList, pages)

			modal := tview.NewGrid().
				SetColumns(0, 40, 0).
				SetRows(0, 13, 0).
				AddItem(form, 1, 1, 1, 1, 0, 0, true)

			pages.AddPage("add_fact_form", modal, true, true)

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

func createFactForm(factsList *tview.List, pages *tview.Pages) *tview.Form {
	form := tview.NewForm()

	form.AddInputField("Identifier", "", 20, nil, nil)
	form.AddInputField("Operator", "", 10, nil, nil)
	form.AddInputField("Value", "", 20, nil, nil)
	form.AddInputField("Result", "", 20, nil, nil)

	form.AddButton("Save", func() {
		conditionField := form.GetFormItemByLabel("Identifier").(*tview.InputField)
		operatorField := form.GetFormItemByLabel("Operator").(*tview.InputField)
		valueField := form.GetFormItemByLabel("Value").(*tview.InputField)
		resultField := form.GetFormItemByLabel("Result").(*tview.InputField)

		identifier := conditionField.GetText()
		operator := operatorField.GetText()
		value := valueField.GetText()
		result := resultField.GetText()

		rule := models.CreateRule(identifier, operator, value, result)

		if rule == nil {
			pages.RemovePage("add_fact_form")
			return
		}

		AddFact(factsList, rule)

		pages.RemovePage("add_fact_form")
	})

	form.AddButton("Cancel", func() {
		pages.RemovePage("add_fact_form")
	})

	form.SetBorder(true).
		SetTitle("Add New Rule").
		SetTitleAlign(tview.AlignCenter)

	form.SetButtonsAlign(tview.AlignCenter)

	return form
}

func AddFact(factList *tview.List, rule *models.Rule) {
	factList.AddItem(rule.String(), "", 0, nil)
}

func addFact(factsList *tview.List, rule *models.Rule) {
	factsList.AddItem(rule.String(), "", 0, nil)
}
