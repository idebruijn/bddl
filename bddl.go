package bddl

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/wsxiaoys/terminal/color"
)

// PrintColorsEnabled can be changed to disable the use of terminal coloring.
// One usecase is to add a command line flag to your test that controls its value.
//
//	func init() {
//		flag.BoolVar(&bddl.PrintColorsEnabled, "color", true, "want colors?")
//	}
//
//	go test -color=false
//
//  or use the environment variable
//	TERMCOLORS=false go test

var PrintColorsEnabled = true

// Check for presence of the TERMCOLORS environment variable to set the TerminalColorsEnabled setting.
func init() {
	PrintColorsEnabled = os.Getenv("TERMCOLORS") != "false"
}

func bddl(t *testing.T, lead string, args ...interface{}) {
	var msg bytes.Buffer

	if PrintColorsEnabled {
		fmt.Fprint(&msg, color.Sprintf("@{g}%s ", lead))
		for i, each := range args {
			if i%2 == 0 {
				fmt.Fprint(&msg, color.Sprintf("@{g}%v ", each))
			} else {
				fmt.Fprint(&msg, color.Sprintf("@{m}%v ", each))
			}
		}
	} else {
		fmt.Fprint(&msg, " ", lead)
		for _, each := range args {
			fmt.Fprint(&msg, " ", each)
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
	if PrintColorsEnabled {
		t.Log(append(append([]interface{}{}, fmt.Sprintf(" ")+color.Sprintf("@{b}%s", "Comment:"+" "+comment)), args...)...)
	} else {
		t.Log(append(append([]interface{}{}, "Comment:"+" "+comment), args...)...)
	}
}

func TODO(t *testing.T, comment string, args ...interface{}) {
	if PrintColorsEnabled {
		t.Log(append(append([]interface{}{}, fmt.Sprintf(" ")+color.Sprintf("@{R}%s", "TODO:"+" "+comment)), args...)...)
	} else {
		t.Log(append(append([]interface{}{}, "Comment:"+" "+comment), args...)...)
	}
}
