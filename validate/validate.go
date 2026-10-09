// Package validate provides deliberately limited syntactic validators for
// common user-input formats. These checks do not establish that an address
// exists or that a phone number is assigned.
package validate

import "strings"

// ValidateEmail accepts ASCII dot-atom email addresses with a multi-label
// domain and a final label of at least two letters. The local part is limited
// to 64 bytes, each domain label to 63 bytes, and the full address to 254 bytes.
// It rejects whitespace, Unicode, quoted local parts, address literals, and
// leading, trailing, or consecutive dots. It is not a full RFC validator.
func ValidateEmail(s string) bool {
	if len(s) == 0 || len(s) > 254 || strings.IndexFunc(s, isWhitespace) >= 0 {
		return false
	}

	at := strings.IndexByte(s, '@')
	if at <= 0 || at != strings.LastIndexByte(s, '@') || at == len(s)-1 {
		return false
	}

	local, domain := s[:at], s[at+1:]
	if len(local) > 64 || local[0] == '.' || local[len(local)-1] == '.' || strings.Contains(local, "..") {
		return false
	}
	for _, r := range local {
		if !isASCIIAlphaNumeric(r) && !strings.ContainsRune("!#$%&'*+-/=?^_`{|}~.", r) {
			return false
		}
	}

	labels := strings.Split(domain, ".")
	if len(labels) < 2 {
		return false
	}
	for _, label := range labels {
		if len(label) == 0 || len(label) > 63 || label[0] == '-' || label[len(label)-1] == '-' {
			return false
		}
		for _, r := range label {
			if !isASCIIAlphaNumeric(r) && r != '-' {
				return false
			}
		}
	}

	tld := labels[len(labels)-1]
	if len(tld) < 2 {
		return false
	}
	for _, r := range tld {
		if !isASCIIAlpha(r) {
			return false
		}
	}
	return true
}

// ValidatePhone accepts 8–15 ASCII digits, optionally prefixed by '+'.
// Separators and all other characters are rejected; this is a format check,
// not a check that the number is assigned or reachable.
func ValidatePhone(s string) bool {
	if len(s) > 0 && s[0] == '+' {
		s = s[1:]
	}
	if len(s) < 8 || len(s) > 15 {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\v' || r == '\f'
}

func isASCIIAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isASCIIAlphaNumeric(r rune) bool {
	return isASCIIAlpha(r) || r >= '0' && r <= '9'
}
