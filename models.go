package main

import (
	"errors"
)

type Rule struct {
	Identifier string
	Operator   Operator
	Value      string
	Result     string
}

// func (r *Rule) Evaluate(factValue string) (bool, error) {
// 	// TODO: Implement evaluation logic based on operator type and values
// 	return false, nil
// }

func (r *Rule) String() string {
	return "If " + r.Identifier + " " + r.Operator.String() + " " + r.Value + " then " + r.Result
}

func CreateRule(identifier, operator, value, result string) *Rule {
	if !isRuleValid(identifier, operator, value, result) {
		return nil
	}

	operatorValue, err := OperatorFromString(operator)

	if err != nil {
		operatorValue = Equal
	}

	return &Rule{
		Identifier: identifier,
		Operator:   operatorValue,
		Value:      value,
		Result:     result,
	}
}

func isRuleValid(identifier, operator, value, result string) bool {
	return identifier != "" && operator != "" && value != "" && result != ""
}

type Operator int

const (
	Equal Operator = iota
	GreaterThan
	LessThan
	NotEqual
	GreaterThanOrEqual
	LessThanOrEqual
	And
	Or
)

func (o Operator) String() string {
	switch o {
	case Equal:
		return "=="
	case GreaterThan:
		return ">"
	case LessThan:
		return "<"
	case NotEqual:
		return "!="
	case GreaterThanOrEqual:
		return ">="
	case LessThanOrEqual:
		return "<="
	case And:
		return "AND"
	case Or:
		return "OR"
	default:
		return "unknown"
	}
}

func OperatorFromString(s string) (Operator, error) {
	if op, found := operatorMap[s]; found {
		return op, nil
	}

	return 0, errors.New("unknown operator: " + s)
}

var operatorMap = map[string]Operator{
	"==":  Equal,
	">":   GreaterThan,
	"<":   LessThan,
	"!=":  NotEqual,
	">=":  GreaterThanOrEqual,
	"<=":  LessThanOrEqual,
	"&":   And,
	"|":   Or,
	"AND": And,
	"OR":  Or,
}

type Fact struct {
	Identifier string
	Value      string
}

type RuleBase struct {
	Rules []Rule
	Facts []Fact
}

func (rb *RuleBase) AddRule(rule Rule) {
	rb.Rules = append(rb.Rules, rule)
}

func (rb *RuleBase) AddFact(fact Fact) {
	rb.Facts = append(rb.Facts, fact)
}

func (rb *RuleBase) GetFactValue(identifier string) (string, bool) {
	for _, fact := range rb.Facts {
		if fact.Identifier == identifier {
			return fact.Value, true
		}
	}
	return "", false
}
