// Task 2 uses table-driven cases for both validators. Each function has a
// deliberately documented syntax policy in validate.go; these tests specify
// the supported and rejected formats, not full address/number verification.
package validate

import (
	"strings"
	"testing"
)

const minCases = 8

// emailCases specifies the ASCII dot-atom and multi-label domain policy.
var emailCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid simple", "student@softserve.academy", true},
	{"valid basic address", "alice@example.com", true},
	{"valid subdomain and tag", "first.last+tag@example.co.uk", true},
	{"valid underscore local part", "a_b-c@example-domain.org", true},
	{"valid dot-atom punctuation", "a!#$%&'*+-/=?^_`{|}~@example.org", true},
	{"valid short labels", "x@ab.cd", true},
	{"missing at sign", "student-softserve.academy", false},
	{"empty string", "", false},
	{"empty local part", "@example.com", false},
	{"empty domain", "student@", false},
	{"multiple at signs", "student@@example.com", false},
	{"consecutive local dots", "alice..smith@example.com", false},
	{"domain label starts with hyphen", "alice@-example.com", false},
	{"consecutive domain dots", "alice@example..com", false},
	{"leading whitespace", " alice@example.com", false},
	{"single-letter TLD", "alice@example.c", false},
	{"unicode local part", "用户@example.com", false},
}

func TestValidateEmail(t *testing.T) {
	if len(emailCases) < minCases {
		t.Fatalf(
			"emailCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(emailCases), minCases,
		)
	}

	for _, tc := range emailCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateEmail(tc.input)
			if got != tc.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestValidateEmail_LengthBoundaries(t *testing.T) {
	local64 := strings.Repeat("a", 64)
	local65 := strings.Repeat("a", 65)
	domain189 := strings.Repeat("a", 63) + "." + strings.Repeat("b", 63) + "." + strings.Repeat("c", 61)
	domain190 := strings.Repeat("a", 63) + "." + strings.Repeat("b", 63) + "." + strings.Repeat("c", 62)

	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"maximum local part and total length", local64 + "@" + domain189, true},
		{"local part too long", local65 + "@example.com", false},
		{"total address too long", local64 + "@" + domain190, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ValidateEmail(tc.input); got != tc.want {
				t.Errorf("ValidateEmail() = %v, want %v", got, tc.want)
			}
		})
	}
}

// phoneCases specifies the optional-plus, digits-only phone policy.
var phoneCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid with plus", "+380501234567", true},
	{"valid digits only", "0501234567", true},
	{"valid minimum length", "12345678", true},
	{"valid maximum length", "123456789012345", true},
	{"empty string", "", false},
	{"contains letters", "050-abc-4567", false},
	{"too few digits", "1234567", false},
	{"too many digits", "1234567890123456", false},
	{"plus without digits", "+", false},
	{"double plus", "++380501234567", false},
	{"hyphen separator", "+380-50-123-4567", false},
	{"embedded whitespace", "050 123 4567", false},
	{"parentheses", "+(380)501234567", false},
	{"unicode digits", "+٣٨٠٥٠١٢٣٤٥٦٧", false},
}

func TestValidatePhone(t *testing.T) {
	if len(phoneCases) < minCases {
		t.Fatalf(
			"phoneCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(phoneCases), minCases,
		)
	}

	for _, tc := range phoneCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidatePhone(tc.input)
			if got != tc.want {
				t.Errorf("ValidatePhone(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
