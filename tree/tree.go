package tree

import (
	"fmt"
	
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

// Tree struct represents a working item tree object
type Tree struct {
	NodeTable map[int]*Node
	Roots []*Node
	ItemIDs []int
}

// NewWorkItemTree creates new Tree object
func NewWorkItemTree(items []workitemtracking.WorkItemLink) *Tree {
	tree := &Tree{
		NodeTable: make(map[int]*Node), 
		Roots: []*Node{},
		ItemIDs: []int{},
	}

	for _, workItem := range items {
		var parentID int

		if workItem.Source != nil {
			parentID = *workItem.Source.Id
		} else {
			parentID = -1
		}

		tree.ItemIDs = append(tree.ItemIDs, *workItem.Target.Id)
		tree.add(*workItem.Target.Url, *workItem.Target.Id, parentID)
	}

	return tree
}

// Show method prints a working item tree
func (tree *Tree) Show() {
	fmt.Printf("WORK ITEM TREE:\n\n")
	
	for _, branch := range tree.Roots {
		branch.show("")
	}
}

// MergeTitles method adds titles for each Node in working item tree from workitemtracking.WorkItem array
func (tree *Tree) MergeTitles(workItems []workitemtracking.WorkItem){
	for _, item := range workItems {
		title := fmt.Sprintf("%v", (*item.Fields)["System.Title"])
		
		tree.NodeTable[*item.Id].Title = title
	}
}

func (tree *Tree) add(URL string, id, parentID int) {
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