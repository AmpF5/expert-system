package ui

import (
	"expert-system/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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

	return rulesList
}

func createAddRuleButton(app *tview.Application, rulesList *tview.List, pages *tview.Pages) *tview.Button {
	b := tview.NewButton("+Add new rule").
		SetSelectedFunc(func() {
			form := createRuleForm(rulesList, pages)

			modal := tview.NewGrid().
				SetColumns(0, 40, 0).
				SetRows(0, 13, 0).
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
			pages.RemovePage("add_rule_form")
			return
		}

		AddRule(rulesList, rule)

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

func AddRule(rulesList *tview.List, rule *models.Rule) {
	rulesList.AddItem(rule.String(), "", 0, nil)
}

func addRule(rulesList *tview.List, rule *models.Rule) {
	rulesList.AddItem(rule.String(), "", 0, nil)
}
