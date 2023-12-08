package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type hm map[[16]byte][]string

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: duplicate path/to/search")
		os.Exit(1)
	}
	path := os.Args[1]

	m := make(hm)
	err := filepath.WalkDir(path, m.compare)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while comparing files:", err)
		os.Exit(1)
	}

	for _, paths := range m {
		if len(paths) > 1 {
			fmt.Println("Duplicates:", paths)
		}
	}
}

func (m hm) compare(path string, d fs.DirEntry, err error) error {
	if d.IsDir() {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	sum := md5.Sum(data)
	m[sum] = append(m[sum], path)
	return nil
}
