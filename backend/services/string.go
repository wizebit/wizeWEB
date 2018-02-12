package services

import "strings"

func Trim(s string) string {
	s = strings.TrimSuffix(s, "\n")
	return s
}