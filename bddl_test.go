package bddl

import "testing"

// go test -v
func TestAll(t *testing.T) {
	Given(t, "given")
	And(t, "and")
	Comment(t, "comment")
	Then(t, "then")
	Then(t, "when")
	Feature(t, "feature")
	Scenario(t, "Scenario")
}
