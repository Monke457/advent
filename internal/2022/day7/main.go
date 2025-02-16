package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type file struct {
	directory directory
	path string
	size int
}

type directory struct {
	parent *directory
	path string
	dirs map[string]directory
	files map[string]file
}

func main() {
	data := reader.FileToArray("data/2022/day7.txt")

	

	root := newDirectory("/", nil)
	current := root

	for _, line := range data[1:] {
		if line[0] == '$' {
			if line[2:4] == "ls" {
				continue
			}
			arg := line[5:]
			if arg == ".." {
				current = *current.parent
			} else {
				current = current.dirs[arg]
			}
			continue
		}

		info, name, _ := strings.Cut(line, " ")
		if info == "dir" {
			if _, ok := current.dirs[name]; !ok {
				parent := current
				dir := newDirectory(name, &parent)
				current.dirs[name] = dir
			}
			continue
		}
		
		size, _ := strconv.Atoi(info)
		file := newFile(current, name, size)
		current.files[name] = file
	}

	sizemap := map[string]int{}
	mapSizes(root, &sizemap)

	sum := 0
	for _, size := range sizemap {
		if size > 100000 {
			continue
		}
		sum += size
	}

	fmt.Println(sizemap)
	fmt.Println("Sum:", sum)

	total := 70000000
	required := 30000000
	free := total - sizemap["/"]
	target := required - free
	fmt.Println(total, free, required, target)


	marked := markForDeletion(sizemap, target)
	size := math.MaxInt

	for _, path := range marked {
		if sizemap[path] < size {
			size = sizemap[path]
		}
	}

	fmt.Println("Size of directort to delete:", size)
}

func markForDeletion(sizes map[string]int, target int) []string {
	marked := []string{}
	for path, size := range sizes {
		if size >= target {
			marked = append(marked, path)
		}
	}
	return marked
}

func newFile(dir directory, name string, size int) file {
	return file {
		path: dir.path + name,
		size: size,
		directory: dir,
	}
}

func mapSizes(dir directory, sizes *map[string]int) {
	for _, child := range dir.dirs {
		mapSizes(child, sizes)
		(*sizes)[dir.path] += (*sizes)[child.path]
	}
	for _, file := range dir.files {
		(*sizes)[dir.path] += file.size
	}
}

func newDirectory(name string, parent *directory) directory {
	if parent != nil {
		if parent.parent != nil {
			name = "/" + name
		}
		name = parent.path + name
	}
	dir := directory{
		path: name,
		parent: parent, 
		dirs: map[string]directory{}, 
		files: map[string]file{},
	} 
	return dir
} 

