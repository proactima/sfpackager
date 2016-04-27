package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

// Package ...
var Package = func(c *cli.Context) {
	inputPath := ""
	if c.IsSet("sfprojects") {
		inputPath = c.String("sfprojects")
	} else {
		Exit("sfprojects must be set")
	}

	outputPath := ""
	if c.IsSet("output") {
		outputPath = c.String("output")
	} else {
		Exit("Output must be set")
	}

	prepareOutput(outputPath)

	projectFilesToProcess := findProjectFiles(inputPath)
	var sfApps []*SfApp

	fmt.Printf("Scanning %s for ServiceFabric applications ...\n", inputPath)
	for _, value := range projectFilesToProcess {
		sfApp := &SfApp{}
		sfApp.Init(value)
		sfApps = append(sfApps, sfApp)
	}

	for _, v := range sfApps {
		v.ReadManifest()
		fmt.Printf("%s - %s\n", v.applicationName, v.appManifestFile)
	}
}

func prepareOutput(path string) {
	exists, info, err := CheckPath(path)
	if err != nil {
		Exit("Something went wrong clearing output folder")
	}

	if exists && info.IsDir() {
		err = os.RemoveAll(path)
		if err != nil {
			Exit("Somthing went wrong clearing output folder")
		}
	} else if exists && !info.IsDir() {
		err = os.Remove(path)
		if err != nil {
			Exit("Somthing went wrong clearing output folder")
		}
	}

	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		Exit("Somthing went wrong creating output folder")
	}
}

func findProjectFiles(path string) []string {
	exist, fileInfo, err := CheckPath(path)

	if err != nil {
		Exit("Should not have happened")
	}

	var projectFilesToProcess []string

	if exist == true {
		if fileInfo.IsDir() {
			ignoreDirs := []string{".git", "bin", "obj", "pkg", "node_modules", "packages", "Backup"}
			err = filepath.Walk(path, FindFiles(ignoreDirs, "*.sfproj", &projectFilesToProcess))
		} else {
			projectFilesToProcess = append(projectFilesToProcess, "ye")
		}
	} else {
		Exit("Invalid path given")
	}

	return projectFilesToProcess
}
