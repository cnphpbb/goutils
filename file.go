package goutils

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, e := os.Stat(path)
	return e == nil || os.IsExist(e)
}

func WriteFile(filename string, data []byte) error {
	os.MkdirAll(path.Dir(filename), os.ModePerm)
	return ioutil.WriteFile(filename, data, 0644)
}

func UserHomeDir() string {
	return os.Getenv("HOME")
}

func ConversionInPath(inPath string) string {
	log.Println("Func::ConversionInpath:inpath to:: ", inPath)
	if inPath == "~" || strings.HasPrefix(inPath, "~"+string(os.PathSeparator)) {
		inPath = UserHomeDir() + inPath[1:]
	}
	if inPath == "$HOME" || strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator)) {
		inPath = UserHomeDir() + inPath[5:]
	}

	inPath = os.ExpandEnv(inPath)
	return inPath
}
