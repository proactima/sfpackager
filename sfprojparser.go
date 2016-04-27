package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/xmlpath.v2"
)

// SfApp ..
type SfApp struct {
	services        []SfService
	applicationName string
	appManifestFile string
}

// SfService ..
type SfService struct {
	projectFolder string
}

// Init ..
func (app *SfApp) Init(filePath string) {
	path := xmlpath.MustCompile("//ProjectReference/@Include")
	file, err := os.Open(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to open %s for reading", filePath)
		Exit(errMsg)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	root, err := xmlpath.Parse(reader)
	if err != nil {
		Exit("Oh no")
	}

	basePath := filepath.Dir(filePath)

	iter := path.Iter(root)
	for iter.Next() == true {
		foundPath := iter.Node().String()
		joined := filepath.Join(basePath, foundPath)
		var fullPath string
		fullPath, err = filepath.Abs(joined)
		if err != nil {
			Exit("Woops")
		}

		service := SfService{
			projectFolder: filepath.Dir(fullPath),
		}

		app.services = append(app.services, service)
	}

	var appManifest string
	ignoreDirs := []string{"bin", "obj", "pkg", "Backup"}
	err = filepath.Walk(basePath, FindFile(ignoreDirs, "ApplicationManifest.xml", &appManifest))
	if err != nil {
		Exit("Something went wrong looking for application manifests")
	}

	app.appManifestFile = appManifest
}

// ReadManifest ..
func (app *SfApp) ReadManifest() {
	path := xmlpath.MustCompile("//ApplicationManifest/@ApplicationTypeName")
	file, err := os.Open(app.appManifestFile)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to open %s for reading", app.appManifestFile)
		Exit(errMsg)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	root, err := xmlpath.Parse(reader)
	if err != nil {
		Exit("Oh no")
	}

	if appName, ok := path.String(root); ok {
		app.applicationName = appName
	}
}
