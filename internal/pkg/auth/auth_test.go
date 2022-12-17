package auth

import (
	"testing"
)

func TestCompare(t *testing.T) {
	if err := Compare("$2a$10$/GbRsTwWLggK4vSj4Tgo5OldPMicZR6NnSNV/AaIycjyQXYsNZmi2", "123456"); err != nil {
		t.Fatal(err)
	}
}
