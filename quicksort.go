// To execute Go code, please declare a func main() in a package "main"

package main

import (
  "fmt"
  "math/rand"
  "time"
)


func quicksort(arr []int) ([]int) {
  rand.Seed(time.Now().UnixNano())
  
  i := rand.Intn(len(arr))
  
  var smaller []int
  var larger []int
  
  
  for _, number := range arr {
    if number >= arr[i] {
      larger = append(larger, number)
    } else {
      smaller = append(smaller, number)
    }
    
  } 
 
  if len(smaller) != 0 {
    smaller = quicksort(smaller)
  }
  if len(larger) > 1 {
    larger = quicksort(larger)
  }
  
  return append(smaller, larger...)
}

func main() {

  // for i := 0; i < 5; i++ {
  //   fmt.Println("Hello, World!")
  // }
  var x []int
  
  for i := 0; i <= 50; i++ {
    x = append(x, rand.Int())
  }
  
  
  fmt.Println(quicksort(x))  
}
