# Task 2 — Email/phone validation (table-driven tests)

`ValidateEmail` and `ValidatePhone` are implemented in `validate.go` with
explicitly limited syntax policies. Both test tables in `validate_test.go`
contain at least eight cases, including positive, negative, and boundary
examples.

**Accepted formats**

- Email: ASCII dot-atom local part (maximum 64 bytes); multi-label domain
  (maximum 63 bytes per label); final label has at least two ASCII letters;
  full address maximum 254 bytes. Unicode, whitespace, quoted local parts,
  address literals, and malformed dots/hyphens are rejected. This is not a complete RFC validator.
- Phone: 8–15 ASCII digits with an optional leading `+`. Separators and all
  other characters are rejected.

Run locally with:

```bash
go test -v ./validate/...
```

Each case is its own subtest (`t.Run`), so a single wrong case never
hides the others in the output.
