package utils

import (
	"testing"
)

func TestPathExists(t *testing.T) {

	p := IsExistsInGoPath("/beeweb/utils/functions.go")
	if !p {
		t.Fail()
	}
	t.Log(p, "OK")
}
