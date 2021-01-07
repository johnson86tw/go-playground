package main

import "fmt"

//  In the file system, there are two types of objects File and Folder.
//  There are cases when File and Folder are treated to be the same way.

type component interface {
	search(string)
}

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{name: "Folder1"}
	folder2 := &folder{name: "Folder2"}

	folder1.add(file1)
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	folder2.search("rose")
}
