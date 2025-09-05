package main

import "errors"

var bracketParenthesis = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

var bracket = map[rune]bool{
	'{': true,
	'[': true,
	'(': true,
	'}': true,
	']': true,
	')': true,
}

func BracketSequenceValidator(str string) (bool, error) {
	if len(str) == 0 {
		return true, nil
	}
	stack := NewEmptyStack[rune]()
	runes := []rune(str)
	for _, char := range runes {
		if !bracket[char] {
			return false, errors.New("only brackets are allowed")
		}
		if closeChar, ok := bracketParenthesis[char]; ok {
			stack.Push(closeChar)
			continue
		}
		if val, ok := stack.Pop(); ok && val != char {
			return false, nil
		}
	}
	if len(stack.items) == 0 {
		return true, nil
	}

	return false, nil
}
