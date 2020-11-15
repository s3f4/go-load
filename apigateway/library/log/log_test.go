package log

import (
	"os"
	"testing"
)

func TestSetOutputPaths(t *testing.T) {
	os.Setenv(APP_ENV, "development")
	BuildLogger(os.Getenv(APP_ENV))
	Error("test", "test")
	Debug("test", "test")
}
