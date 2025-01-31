package main

import (
	"advent/internal/pkg/reader"
	"fmt"
)

type header struct {
	children int
	metadata int
}

type node struct {
	header header
	children []node
	metadata []int
}

func newNode() node {
	return node{
		header: header{},
		children: []node{},
		metadata: []int{},
	}
}

func main() {
	data := reader.FileToIntArrayByDivider("data/2018/day8.txt", " ")

	root := newNode()
	root.parseTree(data)

	res := root.addMetadata()
	fmt.Println("Sum of metadata:", res)
	res	= root.addMetadataIndexed()
	fmt.Println("Sum of metadata indexed:", res)
}

func (n node) addMetadataIndexed() int {
	res := 0
	if n.header.children == 0 {
		for _, data := range n.metadata {
			res += data
		}
		return res
	}
	for _, val := range n.metadata {
		if val > n.header.children {
			continue
		}
		res += n.children[val-1].addMetadataIndexed()
	}
	return res
}

func (n node) addMetadata() int {
	res := 0
	for _, data := range n.metadata {
		res += data
	}
	for _, child := range n.children {
		res += child.addMetadata()
	}
	return res
}

func (n *node) parseTree(data []int) []int {
	if len(data) < 3 {
		return data
	}
	n.header.children = data[0]
	n.header.metadata = data[1]
	data = data[2:]
	for range n.header.children {
		child := newNode()
		data = child.parseTree(data)
		n.children = append(n.children, child)
	}
	n.metadata = append(n.metadata, data[:n.header.metadata]...)
	return data[n.header.metadata:]
}
