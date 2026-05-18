package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetSSHHosts returns a list of hosts from ~/.ssh/config
func (a *App) GetSSHHosts() []string {
	home, err := os.UserHomeDir()
	if err != nil {
		return []string{}
	}

	configPath := filepath.Join(home, ".ssh", "config")
	file, err := os.Open(configPath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(strings.ToLower(line), "host ") {
			host := strings.TrimSpace(line[5:])
			if host != "*" && host != "" {
				parts := strings.Fields(host)
				hosts = append(hosts, parts...)
			}
		}
	}

	return hosts
}

// ConnectToHost launches alacritty with ssh
func (a *App) ConnectToHost(host string) string {
	cmd := exec.Command("alacritty", "-e", "ssh", host)
	err := cmd.Start()
	if err != nil {
		return fmt.Sprintf("Error starting alacritty: %v", err)
	}
	return "Connected"
}
