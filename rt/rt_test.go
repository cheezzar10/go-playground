package rt

import (
	"os"
	"testing"
)

func TestPrintStack(t *testing.T) {
	PrintStack("*** local vulnerabilities enricher ***")
}

func TestPrintCallerStack(t *testing.T) {
	PrintCallerStack(os.Stdout)
}
