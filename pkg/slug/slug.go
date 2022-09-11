package slug

import (
	"regexp"
	"strings"
)

func Slugify(s string) string {
	// format s to only allow alphanumeric and hyphens
	pattern := regexp.MustCompile("[^a-zA-Z0-9-]+")
	s = pattern.ReplaceAllString(s, "")
	// remove leading and trailing hyphens
	s = strings.Trim(s, "-")
	// convert to lowercase
	s = strings.ToLower(s)

	return s
}
