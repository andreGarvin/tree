// Package tree is a package for recursively retreives and returns a tree of sorts of all the files and sub folders in the given directory
package tree

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Directory is a struct that reprents a node in the directory tree
type Directory struct {
	Parent   string
	Children []Directory
}

// Tree returns the tree of the given directory
func Tree(path string) (tree Directory, err error) {

	// getting the absolute path
	absoluteFilePath, err := filepath.Abs(path)
	if err != nil {
		return tree, err
	}

	// adds the absolute file path of the parent folder
	tree.Parent = absoluteFilePath

	// getting the files of that directory
	directory, err := ioutil.ReadDir(absoluteFilePath)
	if err != nil {
		return tree, err
	}

	// this initiates the recursive path collection
	result, err := createDictoryTree(tree.Parent, directory)
	if err != nil {
		return tree, err
	}

	// setting the children of that path to the root tree
	tree.Children = result

	return tree, err
}

// createDictoryTree recursively gets all files paths in a given directory
func createDictoryTree(rootPath string, files []os.FileInfo) (children []Directory, err error) {
	if len(files) == 0 {
		return children, nil
	}

	// getting the first file in the directory
	firstChild := files[0]

	// getting the file absolute file path
	childAbsolutePath, err := filepath.Abs(filepath.Join(rootPath, firstChild.Name()))
	if err != nil {
		return children, err
	}

	// create a new Directory struct for the file
	firstBranch := Directory{
		Parent: childAbsolutePath,
	}

	// if the file is directory then start a new recursive directory dive
	if firstChild.IsDir() {

		subdirectory, err := ioutil.ReadDir(childAbsolutePath)
		if err != nil {
			return nil, err
		}

		subTree, err := createDictoryTree(firstBranch.Parent, subdirectory)
		if err != nil {
			return children, nil
		}

		// adding all the return sub folder files from that directory or lower
		firstBranch.Children = append(firstBranch.Children, subTree...)
	}

	// adding the first file into the slice
	children = append(children, firstBranch)

	// starting a new directory search dive on the result of the files provided in that directory
	subTree, err := createDictoryTree(rootPath, files[1:])
	if err != nil {
		return children, nil
	}

	// adding the rest of the sub files into the slice
	children = append(children, subTree...)

	return children, nil
}
