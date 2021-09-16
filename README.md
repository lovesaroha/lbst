# lbst
This is a generalized binary search tree package with clean and transparent API for the Go language.

## Features
- Lightweight and Fast.
- Native Go implementation.
- Support all data types.

## Requirements
- Go 1.9 or higher. We aim to support the 3 latest versions of Go.

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get -u github.com/lovesaroha/lbst
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage

### Create Binary Search Tree

```Golang
  // Create a binary search tree.
  bst := lbst.Create()
  // Insert values.
  bst.Insert(6)
  bst.Insert(2)
  bst.Insert(4)
  
  // Search value (If not found return nil).
  bst.Search(4)

  // See values in order.
  bst.InOrder()
```
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/49.png)

### Binary Search Tree With Strings
```Golang
  // Create a binary search tree.
  bst := lbst.Create()
  // Insert values.
  bst.Insert("love")
  bst.Insert("saroha")
  bst.Insert("github")
  
  // Search value (If not found return nil).
  bst.Search("love")

  // See values in order.
  bst.InOrder()
``` 
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/50.png)

### Add Multiple Values

```Golang
  // Create a binary search tree.
  bst := lbst.Create()

  // Add values in bst.
  bst.InsertValues([]interface{}{5,3,1,6,8,7,0,9})

  // See values in order.
  bst.InOrder()
```
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/51.png)


### Binary Search Tree With Custom Data Type
```Golang 
  type user struct {
    name string
    age int
  }

  // Compare function for custom data type.
  func compare(a interface{} , b interface{}) bool {
    // For less than return true.
    return  a.(user).age <= b.(user).age
  }

  // Create a binary search tree.
  bst := lbst.Create()

  // Add values in bst.
  bst.InsertWith(user{name: "love" , age: 27} , compare)
  bst.InsertWith(user{name: "joe" , age: 30} , compare)
  bst.InsertWith(user{name: "doe" , age: 34} , compare)

  // See values in order.
  bst.InOrder()
```
![graph](https://raw.githubusercontent.com/lovesaroha/gimages/main/52.png)