package test

import (
	myhwproj "github.com/KaperD/HSE-SD-MyHwProj/internal"
	"testing"
)

func TestTest(t *testing.T) {
	if !myhwproj.IsZeroValue("") {
		t.Fail()
	}
	if !myhwproj.IsZeroValue(nil) {
		t.Fail()
	}
	if !myhwproj.IsZeroValue(0) {
		t.Fail()
	}
}
