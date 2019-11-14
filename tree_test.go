package tree

import (
	"path/filepath"
	"testing"
)

func TestReturnedTree(t *testing.T) {
	testCases := []struct {
		name                string
		root                string
		exceptedChildLength int
	}{
		{"checking folder childern count", "fixtures/test-case-1", 5},
	}

	for _, testCase := range testCases {
		tree, err := Tree(testCase.root)
		if err != nil {
			t.Fatalf("Test case %s failed; %v", testCase.name, err)
		}

		childernLength := len(tree.Children)
		if childernLength != testCase.exceptedChildLength {
			t.Errorf("Test case %q failed; Expected %d childern but got %d in %s", testCase.name, testCase.exceptedChildLength, childernLength, testCase.root)
		}
	}
}

func TestTreeDeepth(t *testing.T) {
	const rootPath string = "fixtures/test-case-2"

	tree, err := Tree(rootPath)
	if err != nil {
		t.Fatalf("Diff tree test case failed; %v", err)
	}

	child := tree.Children[0].Children[0].Children[0]
	absolute, err := filepath.Abs(filepath.Join(rootPath, "sub", "folder", "sub-folder-file", "file"))
	if err != nil {
		t.Fatalf("Diff tree test case failed; %v", err)
	}

	if child.Parent == absolute {
		t.Errorf("Diff tree test case failed; Expected child depth of three path to match absolute %s", absolute)
	}
}
