package log

import (
	"context"
	"fmt"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	l := NewPrinter("stdout", fmt.Printf)
	l.Infof(context.TODO(), "Hello, %s!", "world")
}
