package log

import (
	"fmt"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	var l Logger
	l = NewPrinter("stdout", fmt.Printf)
	l.Infof(nil, "Hello, %s!", "world")
}
