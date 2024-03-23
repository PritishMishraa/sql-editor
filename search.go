package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DatabaseInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var skipDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	// Add more directories to skip if needed
}

func searchSQLiteDatabases(rootDir string) []DatabaseInfo {
	var dbs []DatabaseInfo

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		if info.IsDir() && skipDirs[info.Name()] {
			return filepath.SkipDir
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".db") || strings.HasSuffix(path, ".sqlite") || strings.HasSuffix(path, ".sqlite3")) {
			db := DatabaseInfo{Name: info.Name(), Path: path}
			dbs = append(dbs, db)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the file system: %v\n", err)
	}

	return dbs
}
