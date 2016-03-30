package main

import (
  "fmt"
  "math/rand"
)

type tree struct {
  rightNode *tree
  leftNode *tree
  
  value int
}

func (t *tree) addNode(a int) {
  if a > t.value {
    if t.rightNode == nil {
      t.rightNode = &tree{value : a}
    } else {
      t.rightNode.addNode(a)
    }
  } else {
    if t.leftNode == nil {
      t.leftNode = &tree{value : a}
    } else {
      t.leftNode.addNode(a)
    }
  }
}

func (t *tree) show() {

  printed := false
  
  if t.leftNode == nil {
    fmt.Println(t.value)
    printed = true
  } else {
    t.leftNode.show()    
  }
  
  if t.rightNode == nil {
    if !printed {
      fmt.Println(t.value)
    }
  } else {
    t.rightNode.show()
  }

   
}

func main() {
  myTree := tree{value: 42}
  for i :=0; i<50; i++ {
    myTree.addNode(rand.Intn(100))
  }
  myTree.show()
}
