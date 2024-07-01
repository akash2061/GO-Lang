package test

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	r, err := IsEmail("hello")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(r)
	}
	r, err = IsEmail("example@fuu.com")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(r)
	}
}
