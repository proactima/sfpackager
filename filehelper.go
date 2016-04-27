package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindFile ..
func FindFile(ignoreDirs []string, matchString string, foundFile *string) filepath.WalkFunc {
	found := false
	return func(path string, f os.FileInfo, err error) error {
		if found {
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return nil
		}

		if f.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		}

		matched, err := filepath.Match(matchString, f.Name())
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if matched {
			*foundFile = path
			found = true
		}

		return nil
	}
}

// FindFiles ..
func FindFiles(ignoreDirs []string, matchString string, foundFiles *[]string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if f.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		}

		matched, err := filepath.Match(matchString, f.Name())
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if matched {
			*foundFiles = append(*foundFiles, path)
		}

		return nil
	}
}

// CheckPath ..
func CheckPath(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, fileInfo, nil
	}
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	return true, fileInfo, err
}
