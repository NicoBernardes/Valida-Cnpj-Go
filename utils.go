package main

import "regexp"

func removeNonDigits(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllLiteralString(s, "")
}
