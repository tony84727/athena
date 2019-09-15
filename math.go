package main

import (
	"regexp"
	"strings"
)

var (
	mathExtractPattern = regexp.MustCompile(`\(?\s*\d+\s*(\s*[+\-*/]?\s*\(?\s*\d+\s*\)?\s*)*`)
)

func extractMathExpressions(input string) []string {
	matches :=  mathExtractPattern.FindAllString(input, -1)
	if matches == nil {
		return nil
	}
	for i,m := range matches {
		matches[i] = strings.TrimSpace(m)
	}
	return matches
}

