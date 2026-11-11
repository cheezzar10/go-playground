package rt

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestPrintStack(t *testing.T) {
	PrintStack("*** local vulnerabilities enricher ***")
}

func TestPrintCallerStack(t *testing.T) {
	PrintCallerStack(os.Stdout)
}

func TestPlatform(t *testing.T) {
	fmt.Printf("go OS: %s\n", runtime.GOOS)
}
