/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package lbst

import (
	"fmt"
	"strings"
)

// Binary search tree object structure.
type BinarySearchTreeObject struct {
	root *nodeObject
}

// Node object structure defined.
type nodeObject struct {
	value interface{}
	left  *nodeObject
	right *nodeObject
	red   bool
}

// This function insert multiple values in binary search tree.
func (bst *BinarySearchTreeObject) InsertValues(values []interface{}) {
	for _, value := range values {
		bst.Insert(value)
	}
}

// This function insert new value in binary search tree.
func (bst *BinarySearchTreeObject) Insert(value interface{}) {
	bst.root = put(bst.root, value, isLessThan)
}

// This function put value in node.
func put(node *nodeObject, value interface{}, compare func(valueA interface{}, valueB interface{}) bool) *nodeObject {
	if node == nil {
		// If node is nil than return new node.
		return &nodeObject{value: value, red: true}
	}
	if compare(value, node.value) {
		// Left subtree.
		node.left = put(node.left, value, compare)

	} else {
		// Right subtree.
		node.right = put(node.right, value, compare)
	}

	// Balance red black binary search tree.
	if isRed(node.right) && !isRed(node.left) {
		// Right node link is red so rotate left.
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		// Two red links in left so rotate right.
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		// Two red links change colors.
		changeColors(node)
	}
	return node
}

// This function return true if given node link is red.
func isRed(node *nodeObject) bool {
	if node == nil {
		return false
	}
	return node.red
}

// This function rotate given node left.
func rotateLeft(node *nodeObject) *nodeObject {
	var rightNode = node.right
	node.right = rightNode.left
	rightNode.left = node
	rightNode.red = node.red
	node.red = true
	return rightNode
}

// This function rotate given node right.
func rotateRight(node *nodeObject) *nodeObject {
	var leftNode = node.left
	node.left = leftNode.right
	leftNode.right = node
	leftNode.red = node.red
	node.red = true
	return leftNode
}

// This function change color of given node.
func changeColors(node *nodeObject) {
	node.left.red = false
	node.right.red = false
	node.red = true
}

// This function return if less than.
func isLessThan(valueA interface{}, valueB interface{}) bool {
	switch valueAType := valueA.(type) {
	case int:
		// For b.
		switch valueBType := valueB.(type) {
		case int:
			if valueAType > valueBType {
				return false
			}
		case float64:
			if valueAType > int(valueBType) {
				return false
			}
		}
	case float64:
		// For b.
		switch valueBType := valueB.(type) {
		case int:
			if valueAType > float64(valueBType) {
				return false
			}
		case float64:
			if valueAType > valueBType {
				return false
			}
		}
	case string:
		// For b.
		switch valueBType := valueB.(type) {
		case string:
			if strings.Compare(valueAType, valueBType) > 0 {
				return false
			}
		}
	}
	return true
}

// This function insert value with given condition.
func (bst *BinarySearchTreeObject) InsertWith(value interface{}, compare func(valueA interface{}, valueB interface{}) bool) {
	bst.root = put(bst.root, value, compare)
}

// This function search given value in binary search tree.
func (bst BinarySearchTreeObject) Search(value interface{}) interface{} {
	var currentNode = bst.root
	for {
		if currentNode == nil {
			// Not found.
			return nil
		}
		if currentNode.value == value {
			// Value match found.
			return currentNode.value
		}
		if isLessThan(value, currentNode.value) {
			// Search in left tree.
			currentNode = currentNode.left
		} else {
			// Search in right tree.
			currentNode = currentNode.right
		}
	}
}

// This function perform in order traversal of binary search tree.
func (bst BinarySearchTreeObject) InOrder() {
	inOrder(bst.root)
}

// This function perform in order traversal of binary search tree.
func inOrder(node *nodeObject) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Println(node.value)
	inOrder(node.right)
}

// This function create new bst.
func Create() BinarySearchTreeObject {
	return BinarySearchTreeObject{}
}
