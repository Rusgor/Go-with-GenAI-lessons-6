# Task 3 — AI blind test generation

## Function under test

`func ValidateEmail(s string) bool` from `validate/validate.go`.

## Prompt you gave the AI

> For an AI blind-testing exercise, consider ONLY this Go signature: `func ValidateEmail(s string) bool`. Do not inspect or search the workspace or use any context about its implementation. Generate a concise table-driven set of at least 10 test cases with input and expected bool, including ordinary valid/invalid values and edge cases. State the assumed validation policy behind expectations so any policy-dependent assertions are explicit. Return only cases plus assumptions; do not implement code.

## Edge cases the AI found that I had missed

The AI repeated two checks already in my initial table (empty input and a missing `@`) and added extra `@`, empty local/domain parts, consecutive local dots, malformed domain hyphens and dots, whitespace, Unicode, and a one-character top-level domain. I incorporated its generated cases into `emailCases` in `validate/validate_test.go`. I then ran them against the implementation with `go test -v -count=1 ./...`; all generated cases passed. The AI also made its assumptions explicit: conventional ASCII dot-atom syntax, multi-label domains, and a final label of at least two letters.

## Edge cases I had that the AI missed

The generated set did not test the local-part and whole-address length boundaries. I added tests for a 64-character local part at the 254-character total limit, a 65-character local part, and a 255-character address. Those boundary cases passed under the documented policy.

## Cases where the AI's expected output was wrong

None of its expected results conflicted with the policy chosen for this exercise. Some syntax choices—such as rejecting single-label domains, quoted local parts, or Unicode—are policy decisions rather than universal definitions of email validity, so I documented those limits in `validate.go` rather than treating them as universal rules.

## What I'd change about my own test-writing process after this

I will write down the accepted input policy before building test tables, then cover both malformed structure and boundary lengths. The AI's output was useful as a checklist, but its assumptions still needed review and alignment with the implementation. These tests check syntax only; they cannot prove that an address exists or can receive mail.

## Review status

This report and the implementation are ready for mentor review; mentor approval has not yet been recorded.
