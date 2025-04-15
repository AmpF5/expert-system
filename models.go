package main

// import (
// 	"errors"
// )

// type RuleBase struct {
// 	Rules []Rule
// 	Facts []Fact
// }

// func (rb *RuleBase) AddRule(rule Rule) {
// 	rb.Rules = append(rb.Rules, rule)
// }

// func (rb *RuleBase) AddFact(fact Fact) {
// 	rb.Facts = append(rb.Facts, fact)
// }

// func (rb *RuleBase) GetFactValue(identifier string) (string, bool) {
// 	for _, fact := range rb.Facts {
// 		if fact.Identifier == identifier {
// 			return fact.Value, true
// 		}
// 	}
// 	return "", false
// }
