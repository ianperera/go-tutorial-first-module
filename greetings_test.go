package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Ian"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ian")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ian") = %q, %v want to match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
