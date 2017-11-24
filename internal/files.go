package internal

import (
	"os"
	"path/filepath"
)

// Directory returns a slice of filenames. If the given filename
// refers to a directory, return a slice of names of files that are in
// this directory. If the given filename does not refer to a
// directory, return a slice with this filename as the only entry.
func Directory(file string) (files []string, err error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return []string{filepath.Base(file)}, nil
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		nerr := f.Close()
		if err == nil {
			err = nerr
		}
	}()
	return f.Readdirnames(0)
}

// FullPathname returns the given filename if it is
// absolute. Otherwise, it joins it with the current working
// directory.
func FullPathname(filename string) (string, error) {
	if filepath.IsAbs(filename) {
		return filename, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, filename), nil
}
