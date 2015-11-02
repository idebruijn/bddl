package bddl

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wsxiaoys/terminal/color"
)

func bddl(t *testing.T, lead string, args ...interface{}) {
	var msg bytes.Buffer
	fmt.Fprint(&msg, color.Sprintf("@{g}%s ", lead))
	for i, each := range args {
		if i%2 == 0 {
			fmt.Fprint(&msg, color.Sprintf("@{g}%v ", each))
		} else {
			fmt.Fprint(&msg, color.Sprintf("@{m}%v ", each))
		}
	}
	t.Log(msg.String())
}

func Feature(t *testing.T, args ...interface{}) {
	bddl(t, "Feature:", args...)
}

func Scenario(t *testing.T, args ...interface{}) {
	bddl(t, "Scenario:", args...)
}

func Given(t *testing.T, args ...interface{}) {
	bddl(t, "Given", args...)
}

func When(t *testing.T, args ...interface{}) {
	bddl(t, " When", args...)
}

func Then(t *testing.T, args ...interface{}) {
	bddl(t, " Then", args...)
}

func And(t *testing.T, args ...interface{}) {
	bddl(t, "  and", args...)
}

func Comment(t *testing.T, comment string, args ...interface{}) {
	t.Log(append(append([]interface{}{}, fmt.Sprintf(" ")+color.Sprintf("@{b}%s", "Comment:"+" "+comment)), args...)...)
}

func TODO(t *testing.T, comment string, args ...interface{}) {
	t.Log(append(append([]interface{}{}, fmt.Sprintf(" ")+color.Sprintf("@{R}%s", "TODO:"+" "+comment)), args...)...)
}
