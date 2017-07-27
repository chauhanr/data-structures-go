package trees

import (
	"errors"
	"fmt"
)

type BinaryTree struct {
	Root *TNode
	fill int
}

type TNode struct {
	Item   int
	parent *TNode
	left   *TNode
	right  *TNode
}

func (tree *BinaryTree) Initialize() {
	tree.Root = nil
}

func (tree *BinaryTree) Size() int {
	return tree.fill
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		node := TNode{}
		node.Item = value
		tree.Root = &node
	} else {
		if tree.Root.Item >= value {
			tree.Root.left = tree.insertNode(tree.Root, tree.Root.left, value)
		} else {
			tree.Root.right = tree.insertNode(tree.Root, tree.Root.right, value)
		}
	}
	tree.fill++
}

/**
  Internal method of the tree to handle insert apart from the root element.
  this will check of the current node is null then it add the node to the current place.
  if not then it will search for the best place.
*/
func (tree *BinaryTree) insertNode(parent *TNode, curr *TNode, item int) *TNode {
	n := TNode{}
	n.Item = item
	if curr == nil {
		n.parent = parent
		return &n
	} else {
		if curr.Item >= item {
			curr.left = tree.insertNode(curr, curr.left, item)
			return curr
		} else {
			curr.right = tree.insertNode(curr, curr.right, item)
			return curr
		}
	}
}

func (tree *BinaryTree) SearchItem(item int) *TNode {
	if tree.Root == nil {
		return nil
	}
	if tree.Root.Item == item {
		return tree.Root
	} else if tree.Root.Item > item {
		return tree.searchTree(item, tree.Root.left)
	} else {
		return tree.searchTree(item, tree.Root.right)
	}
}

func (tree *BinaryTree) searchTree(item int, node *TNode) *TNode {
	if node == nil {
		return nil
	}
	if node.Item == item {
		return node
	} else if node.Item > item {
		return tree.searchTree(item, node.left)
	} else {
		return tree.searchTree(item, node.right)
	}
}

/**
just a testing function
*/
func (tree *BinaryTree) inOrderTraveral(node *TNode, output []int) []int {
	if node == nil {
		// do nothing
	} else {
		output = tree.inOrderTraveral(node.left, output)
		output = append(output, node.Item)
		output = tree.inOrderTraveral(node.right, output)
	}
	//fmt.Printf("%v", output)
	return output
}

/**
The following function will try and delete nodes of a tree it will handle 3 types of deletion.
1. delete a node that has no children - easiest just remove the node and do nothing.
2. delete a node with one child - delete the node and replace it with its child.
3. delete a node with two children - take the left most child of the right child.
*/
func (tree *BinaryTree) DeleteNodeByValue(item int) {
	// first search the item and find the node
	nodeToDelete := tree.SearchItem(item)
	if nodeToDelete == nil {
		return
	}
	// if the node to delete is is found check the following
	//1. does the node have no child.
	if nodeToDelete.left == nil && nodeToDelete.right == nil {
		if nodeToDelete.parent.left == nodeToDelete {
			nodeToDelete.parent.left = nil
		} else {
			nodeToDelete.parent.right = nil
		}
		nodeToDelete.parent = nil
	} else if nodeToDelete.left != nil && nodeToDelete.right != nil {
		// check for the test were we have both children. First get the right child
		rightChild := nodeToDelete.right
		leftMostChildItem, err := tree.deleteLeftMostNodeReturnValue(rightChild)
		if err != nil {
			fmt.Printf("Something is wrong.")
		}
		//fmt.Printf("Replacing %d with %d", nodeToDelete.Item, leftMostChildItem)
		nodeToDelete.Item = leftMostChildItem
	} else {
		// this means the node has only one child.
		if nodeToDelete.left != nil {
			if nodeToDelete.parent.left == nodeToDelete {
				nodeToDelete.parent.left = nodeToDelete.left
			} else {
				nodeToDelete.parent.right = nodeToDelete.left
			}
			nodeToDelete.left.parent = nodeToDelete.parent
			nodeToDelete.parent = nil
			nodeToDelete.left = nil
		} else {
			if nodeToDelete.parent.left == nodeToDelete {
				nodeToDelete.parent.left = nodeToDelete.right
			} else {
				nodeToDelete.parent.right = nodeToDelete.right
			}
			nodeToDelete.right.parent = nodeToDelete.parent
			nodeToDelete.parent = nil
			nodeToDelete.right = nil
		}
	}
}

func (tree *BinaryTree) FindMinimumElement() (int, error) {
	if tree.Root == nil {
		return -1, errors.New("Cannot find minimum element on a empty tree")
	}
	minElement := tree.Root
	for minElement.left != nil {
		minElement = minElement.left
	}
	return minElement.Item, nil
}

func (tree *BinaryTree) FindMaxElement() (int, error) {
	if tree.Root == nil {
		return -1, errors.New("Cannot find max element on a empty tree")
	}
	maxElement := tree.Root
	for maxElement.right != nil {
		maxElement = maxElement.right
	}
	return maxElement.Item, nil
}

func (tree *BinaryTree) deleteLeftMostNodeReturnValue(node *TNode) (int, error) {
	if node == nil {
		return -1, errors.New("Cannot find left most node for the current nose as it is nil")
	} else {
		current := node
		leftMostNode := current
		for current != nil {
			leftMostNode = current
			current = current.left
		}
		leftMostNodeValue := leftMostNode.Item
		//fmt.Printf("LeftMost Item is %d\n", leftMostNodeValue)

		if leftMostNode.right == nil {
			//fmt.Printf("Right node if left most node is nil \n")
			if leftMostNode.parent.left == leftMostNode {
				leftMostNode.parent.left = nil
			} else {
				leftMostNode.parent.right = nil
			}
			leftMostNode.parent = nil
		} else {
			// right element is not nil therefore replace right with current left most element
			if leftMostNode.parent.left == leftMostNode {
				leftMostNode.parent.left = leftMostNode.right
			} else {
				leftMostNode.parent.right = leftMostNode.right
			}
			leftMostNode.right.parent = leftMostNode.parent
			leftMostNode.parent = nil
		}
		return leftMostNodeValue, nil
	}
}

// function to compare the two trees
func CompareBSTs(t1 *BinaryTree, t2 *BinaryTree) bool {
	if t1 != nil && t2 != nil {
		if t1.Size() == t2.Size() {
			//fmt.Printf("Trees sizes are equal and size is %d\n", t1.Size())
			if t1.Size() != 0 {
				r1 := t1.Root
				r2 := t2.Root
				return compareTreeNodes(r1, r2)
			} else {
				return true
			}
		} else {
			//fmt.Printf("Tree sizes are not equal tree 1 size is %d and tree 2 is %d\n", t1.Size(), t2.Size())
			return false
		}
	} else if t1 == nil && t2 == nil {
		return true
	} else {
		return false
	}
}

func compareTreeNodes(n1 *TNode, n2 *TNode) bool {
	if n1 != nil && n2 != nil {
		if !compareTreeNodes(n1.left, n2.left) {
			return false
		}
		if n1.Item != n2.Item {
			//fmt.Printf("comparing %d from tree 1 , %d from tree2 \n", n1.Item, n2.Item)
			return false
		}
		if !compareTreeNodes(n1.right, n2.right) {
			return false
		}
		return true
	} else if n1 == nil && n2 == nil {
		return true
	} else {
		return false
	}
}
