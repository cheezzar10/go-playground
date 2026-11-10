package jar

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	VendorSuffix = "-axiom"
)

func GetTomeeHome() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return filepath.Join(userHome, "Applications", "lcat8", "lib"), nil
}

func TestJar(t *testing.T) {
	tomeeHome, err := GetTomeeHome()
	if err != nil {
		t.Fatal("failed to find tomee home:", err)
	}
	jarPath := filepath.Join(tomeeHome, "tomcat-api.jar")
	err = New(jarPath)
	if err != nil {
		t.Error("failed to read jar:", err)
	}
}

func TestHasSuffix(t *testing.T) {
	tomeeHome, err := GetTomeeHome()
	if err != nil {
		t.Fatal("failed to find tomee home:", err)
	}

	jarWithSuffixPath := filepath.Join(tomeeHome, "geronimo-transaction-3.1.5-7-axiom.jar")
	hasSuffix := HasSuffix(jarWithSuffixPath, VendorSuffix)
	if !hasSuffix {
		t.Error("has suffix")
	}

	jarWithoutSuffixPath := filepath.Join(tomeeHome, "tomcat-api.jar")
	hasSuffix = HasSuffix(jarWithoutSuffixPath, VendorSuffix)
	if hasSuffix {
		t.Error("doesn't have suffix")
	}
}
