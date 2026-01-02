package abstractions

import "slices"

type SafetyProtocol struct {
	rules []OrderingRule

	/* Optimized rules map built out of the original rules */
	optimizedRules map[PageNumber][]PageNumber
}

func NewSafetyProtocol(
	rules []OrderingRule,
) *SafetyProtocol {
	return &SafetyProtocol{
		rules:          rules,
		optimizedRules: getOptimizedRules(rules),
	}
}

func (p SafetyProtocol) IsValid(
	pageNumber PageNumber,
	numbersBefore []PageNumber,
) bool {

	if p.optimizedRules[pageNumber] == nil {
		return true
	}

	rulesForPageNumber := p.optimizedRules[pageNumber]

	for _, pageNumberBefore := range numbersBefore {
		/*
			Finds out if the `pageNumberBefore` can actually be printed before `pageNumber`
			If the `pageNumberBefore` is found in the list of rules for the current page number that
			indicates which ones have to be printed after, then the security protocol is violated
		*/
		if slices.Contains(rulesForPageNumber, pageNumberBefore) {
			return false
		}
	}

	return true
}

func getOptimizedRules(
	rules []OrderingRule,
) map[PageNumber][]PageNumber {

	optimizedRules := make(map[PageNumber][]PageNumber)

	for _, rule := range rules {

		pageNumbers, exists := optimizedRules[rule.From]

		if !exists {
			newPageNumbers := make([]PageNumber, 0)
			newPageNumbers = append(newPageNumbers, rule.To)
			optimizedRules[rule.From] = newPageNumbers
		} else {
			pageNumbers = append(pageNumbers, rule.To)
			optimizedRules[rule.From] = pageNumbers
		}

	}

	return optimizedRules
}
