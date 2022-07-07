package config

import (
	"path/filepath"
	"runtime"
)

func getCurrentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

func GetAbsRootPath() string {
	currDir := getCurrentDir()
	return filepath.Dir(currDir)
}

func AbsPathFromProjRoot(fileName string) string {
	return filepath.Join(GetAbsRootPath(), fileName)
}
