package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	//"flag"
)

//non-recursive renaming
// func main() {
// 	//Hardcoded filepath for samples
// 	dir := "C:\\Go\\Projects\\Go-Hello-World\\File Renaming Tool\\samples"
// 	files, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		panic(err)
// 	}
// 	count := 0
// 	var toRename []string
// 	for _, file := range files {
// 		if file.IsDir() {
// 		} else {
// 			_, err := match(file.Name(), 4)
// 			if err == nil {
// 				count++
// 				toRename = append(toRename, file.Name())
// 			}
// 		}
// 	}
// 	for _, origFilename := range toRename {
// 		origPath := filepath.Join(dir, origFilename)
// 		newFilename, err := match(origFilename, count)
// 		if err != nil {
// 			panic(err)
// 		}
// 		newPath := filepath.Join(dir, newFilename)
// 		fmt.Printf("mv %s => %s", origPath, newPath)
// 		err = os.Rename(origPath, newPath)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }

// Filepath.Walk
func main() {

	//Hardcoded filepath for samples
	walkDir := "C:\\Go\\Projects\\Go-Hello-World\\File Renaming Tool\\samples"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)

		if m, err := match(info.Name()); err == nil {
			key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming:", oldPath, newPath, err.Error())
			}
		}
	}
}

type matchResult struct {
	base  string
	index int
	ext   string
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(filename string) (*matchResult, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return nil, fmt.Errorf("%s didn't match our pattern", filename)
	}
	return &matchResult{strings.Title(name), number, ext}, nil
}
