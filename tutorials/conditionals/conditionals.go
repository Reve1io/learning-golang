package conditionals

import "strings"

func DomainForLocale(domain, locale string) string {
	if len(locale) == 0 {
		return "en." + domain
	} else {
		return locale + "." + domain
	}
}

func ModifySpaces(s, mode string) string {

	switch {
	case mode == "dash":
		newS := strings.ReplaceAll(s, " ", "-")
		return newS
	case mode == "underscore":
		newS := strings.ReplaceAll(s, " ", "_")
		return newS
	default:
		newS := strings.ReplaceAll(s, " ", "*")
		return newS
	}
}
