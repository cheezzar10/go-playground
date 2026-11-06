package jar

import (
	"os"
	"path/filepath"
	"testing"
)

func TestJar(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Error("failed to resolve user home dir:", err)
	}
	jarPath := filepath.Join(home, "Applications", "lcat8", "lib", "tomcat-api.jar")
	err = New(jarPath)
	if err != nil {
		t.Error("failed to read jar:", err)
	}
}
