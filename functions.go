// Program to define a function

package main
import "fmt"

// define a function
func greet() {
  fmt.Println("Good Morning")

// function to add two numbers
func addNumbers(n1 int, n2 int) {
  sum := n1 + n2
  fmt.Println("Sum:", sum)
}
 

func main() {
  // function call
  addNumbers()
}
