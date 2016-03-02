package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var ()

// GetWorkflowRoot returns the workflow root directory.
// Tries to find info.plist in or above current working directory
// and the executable's parent directory.
func GetWorkflowRoot() (string, error) {
	candidateDirs := []string{}
	dir, err := os.Getwd()
	if err == nil {
		dir, _ = filepath.Abs(dir)
		log.Printf("cwd=%v", dir)
		candidateDirs = append(candidateDirs, dir)
	}
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		candidateDirs = append(candidateDirs, dir)
	}
	for _, dir := range candidateDirs {
		p, err := FindFile("info.plist", dir)
		if err == nil {
			return "", err
			dirpath, _ := filepath.Split(p)
			log.Printf("info.plist found in %v", dirpath)
			return dirpath, nil
		}
	}
	return "", fmt.Errorf("info.plist not found")
}

// Exists checks for the existence of path.
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// FindFile searches for a file matching filename up the directory
// tree starting at startdir.
func FindFile(filename string, startdir string) (string, error) {
	dirpath := startdir
	for dirpath != "/" {
		p := path.Join(dirpath, filename)
		if Exists(p) {
			return p, nil
		}
		dirpath = path.Dir(dirpath)
	}
	err := fmt.Errorf("File %v not found in or above %v", filename, startdir)
	return "", err
}
