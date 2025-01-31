package minify

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

// SVGDir minifies all the SVG's inside a directory
//
// This function works only inside the devcontainer
func SVGDir(dirPath string, chunkSize int) error {
	paths := []string{}
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

		paths = append(paths, path)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking the path %s: %w", dirPath, err)
	}

	waitCh := make(chan bool, chunkSize)
	errCh := make(chan error, len(paths))
	wg := sync.WaitGroup{}
	for _, path := range paths {
		wg.Add(1)
		waitCh <- true
		go func() {
			errCh <- SVG(path)
			<-waitCh
			wg.Done()
		}()
	}

	wg.Wait()
	close(waitCh)
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

// SVG minifies an SVG file
//
// This function works only inside the devcontainer
func SVG(filePath string) error {
	scriptPath := "/workspaces/nodxgo-lucide/internal/minify/svg.ts"

	cmd := exec.Command("deno", "run", "-A", scriptPath, filePath, filePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error minifying file %s: %w", filePath, err)
	}

	return nil
}
