package tree

import (
	"fmt"
	
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type Tree struct {
	NodeTable map[int]*Node
	Roots []*Node
	ItemIDs []int
}

func NewWorkItemTree(items []workitemtracking.WorkItemLink) *Tree {

	tree := &Tree{
		NodeTable: make(map[int]*Node), 
		Roots: []*Node{},
		ItemIDs: []int{},
	}

	for _, workItem := range items {

		var parentId int

		if workItem.Source != nil {
			parentId = *workItem.Source.Id
		} else {
			parentId = -1
		}

		tree.add(*workItem.Target.Url, *workItem.Target.Id, parentId)
	}

	return tree
}

func (tree *Tree) add(URL string, id, parentID int) {
	tree.ItemIDs = append(tree.ItemIDs, id)
	node := &Node{ID: id, URL: URL, Children: []*Node{}}

	if parentID == -1 {
		tree.Roots = append(tree.Roots, node)
	} else {
		parent, ok := tree.NodeTable[parentID]
		if !ok {
			fmt.Printf("add: parentId=%v: not found\n", parentID)
			return
		}

		parent.Children = append(parent.Children, node)
	}

	tree.NodeTable[id] = node
}