package cmdline

import (
	"strings"
	"testing"
)

func TestGetVersionString(t *testing.T) {
	versionString := GetVersionString()

	if len(versionString) < 1 {
		t.Error("GetVersionString() is empty")
	}

	if !strings.Contains(versionString, " version ") {
		t.Error("GetVersionString() does not contain ' version '")
	}
}
