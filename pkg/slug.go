package pkg

import "strings"

func GenerateSlug(name string) string {
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	return slug
}
