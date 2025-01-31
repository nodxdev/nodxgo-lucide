package minify

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// HTMLDir minifies all the HTML's inside a directory
//
// This function works only inside the devcontainer
func HTMLDir(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".svg" {
			return nil
		}

		return HTML(path)
	})

	if err != nil {
		return fmt.Errorf("error walking the path %s: %w", dirPath, err)
	}

	return nil
}

// HTML minifies an HTML file
//
// This function works only inside the devcontainer
func HTML(filePath string) error {
	scriptPath := "/workspaces/nodxgo-lucide/internal/minify/html.ts"

	cmd := exec.Command("deno", "run", "-A", scriptPath, filePath, filePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error minifying file %s: %w", filePath, err)
	}

	return nil
}
