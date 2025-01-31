package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"strings"
)

type IconInfo struct {
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}

func generateComponent(fileName string, funcName string, svgBytes []byte) string {
	fullSvg := string(svgBytes)

	// Find where the children starts
	start := strings.Index(fullSvg, ">") + 1
	if start == -1 {
		log.Fatalf("could not find the start of the svg tag in %s", fileName)
	}

	// Remove the svg tag and format the svg
	svg := fullSvg[start:]
	svg = strings.ReplaceAll(svg, "</svg>", "")
	svg = strings.TrimSpace(svg)

	previewURL := repoIconsDir + "/" + fileName

	fn := `
		// ` + funcName + ` icon: ` + previewURL + `
		func ` + funcName + `(children ...nodx.Node) nodx.Node {
			return nodxgoSvgWrapper(
				nodx.Group(children...),
				nodx.Raw(` + "`" + svg + "`" + `),
			)
		}
	`

	return fn
}

func generateInfo(fileName string, name string, funcName string, jsonBytes []byte) string {
	var info IconInfo
	if err := json.Unmarshal(jsonBytes, &info); err != nil {
		log.Fatalf("could not unmarshal %s: %v", fileName, err)
	}

	tpl := `{
		Name: "%s",
		Icon: %s,
		Tags: []string{%s},
		Categories: []string{%s},
	},`

	tags := ""
	for _, tag := range info.Tags {
		tags += `"` + tag + `",`
	}

	categories := ""
	for _, category := range info.Categories {
		categories += `"` + category + `",`
	}

	return fmt.Sprintf(tpl, name, funcName, tags, categories)
}

func generatePackageDef() string {
	return `
		// Code generated by github.com/nodxdev/nodxgo-lucide build task. DO NOT EDIT.
		// v` + version + `

		package lucide
	`
}

func generateIconsFile(components []string) []byte {
	pkg := generatePackageDef()
	pkg += "\n"
	pkg += `import nodx "github.com/nodxdev/nodxgo"`
	pkg += "\n"
	pkg += strings.Join(components, "\n\n")

	b, err := format.Source([]byte(pkg))
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func generateInfoFile(infos []string) []byte {
	pkg := generatePackageDef() + `

		import nodx "github.com/nodxdev/nodxgo"

		// IconInfo represents the information of an icon.
		type IconInfo struct {
			Name       string
			Icon       func (children ...nodx.Node) nodx.Node
			Tags       []string
			Categories []string
		}

		// IconsInfo is a list of all the icons information.
		var IconsInfo = []IconInfo{
	`
	pkg += strings.Join(infos, "\n") + "\n"
	pkg += `}`

	b, err := format.Source([]byte(pkg))
	if err != nil {
		log.Fatal(err)
	}

	return b
}
