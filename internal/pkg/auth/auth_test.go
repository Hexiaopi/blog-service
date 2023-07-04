package auth

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	password, err := Encrypt("123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(password)
	if err := Compare(password, "123456"); err != nil {
		t.Fatal(err)
	}
}

func TestCompare(t *testing.T) {
	if err := Compare("$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G", "123456"); err != nil {
		t.Fatal(err)
	}
}
