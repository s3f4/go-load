package log

import (
	"os"
	"testing"
)

func TestSetOutputPaths(t *testing.T) {
	os.Setenv(APPENV, "development")
	BuildLogger(os.Getenv(APPENV))
	Error("test", "test")
	Debug("test", "test")
}
