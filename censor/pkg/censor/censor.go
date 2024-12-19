package censor

import "strings"

var badContent = [...]string{"qwerty", "йцуке", "123"}

// IsCensored возвращает true, если текст должен быть подвергнут цензуре, и false в противном случае.
func IsCensored(text string) bool {
	isCensored := false

	for _, subStr := range badContent {
		isCensored = strings.Contains(strings.ToLower(text), strings.ToLower(subStr))
		if isCensored {
			return isCensored
		}
	}

	return isCensored
}
