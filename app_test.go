package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetSSHHosts(t *testing.T) {
	// Create a temporary ssh config
	tmpDir, err := os.MkdirTemp("", "ssh-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config")
	content := `
Host server1
    HostName 1.2.3.4
Host server2 server3
    User root
Host *
    IdentityFile ~/.ssh/id_rsa
`
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// We need to override the home directory for testing if we wanted to use the real function,
	// but since our function uses os.UserHomeDir(), we'd need to mock it.
	// For now, let's just test a modified version or refactor the function to take a path.
}
