package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/nodxdev/nodxgo-lucide/internal/fileutil"
	"github.com/nodxdev/nodxgo-lucide/internal/minify"
	"github.com/nodxdev/nodxgo-lucide/internal/strutil"
)

// Check for new releases here: https://github.com/lucide-icons/lucide/releases

const (
	version               = "0.474.0"
	iconsURL              = "https://github.com/lucide-icons/lucide/archive/refs/tags/" + version + ".zip"
	repoIconsDir          = "https://raw.githubusercontent.com/lucide-icons/lucide/" + version + "/icons"
	tempDir               = "./tmp"
	infoOutputFilePath    = "./generated_info.go"
	iconsOutputFilePath   = "./generated_icons.go"
	includedIconsFilePath = "./generated_icons.txt"
)

func main() {
	os.RemoveAll(tempDir)
	err := os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Download icons
	iconsFile := path.Join(tempDir, "icons.zip")
	if err := fileutil.DownloadFile(iconsURL, iconsFile); err != nil {
		log.Fatalf("error downloading %s: %v", iconsURL, err)
	}

	// Unzip icons
	if err := fileutil.Unzip(iconsFile, tempDir); err != nil {
		log.Fatal(err)
	}

	// Minify SVG icons
	iconsDir := path.Join(tempDir, "lucide-"+version, "icons")
	if err := minify.SVGDir(iconsDir, 100); err != nil {
		log.Fatal(err)
	}

	// Read icons folder
	files, err := os.ReadDir(iconsDir)
	if err != nil {
		log.Fatal(err)
	}

	// Generate Go code from icons
	components := []string{}
	infos := []string{}
	includedIcons := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := path.Ext(file.Name())
		if ext != ".svg" && ext != ".json" {
			continue
		}

		kebabCaseName := strings.TrimSuffix(file.Name(), ext)
		funcName := strutil.KebabToUpperCamel(kebabCaseName)
		name := strutil.KebabToCapitalized(kebabCaseName)

		filePath := path.Join(iconsDir, file.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		if ext == ".svg" {
			component := generateComponent(file.Name(), funcName, b)
			components = append(components, component)
			includedIcons = append(includedIcons, funcName)
		}

		if ext == ".json" {
			info := generateInfo(file.Name(), name, funcName, b)
			infos = append(infos, info)
		}
	}
	iconsFileContents := generateIconsFile(components)
	infoFileContents := generateInfoFile(infos)

	// Write icons Go code to file
	err = os.WriteFile(iconsOutputFilePath, iconsFileContents, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Write info Go code to file
	err = os.WriteFile(infoOutputFilePath, infoFileContents, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Write icons list to file
	includedIconsFileContents := strings.Join(includedIcons, "\n")
	err = os.WriteFile(includedIconsFilePath, []byte(includedIconsFileContents), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Lucide icons generated successfully!")
}
