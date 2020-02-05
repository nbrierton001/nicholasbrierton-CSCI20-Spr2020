package main

import "fmt"

func main() {
  fmt.Println("We will selected a number between range of 10 and 49, we select 33")
  fmt.Println("Subtract 33 from 99, resulting in 66. We will designate the number 66 as our 'factor'")
  fmt.Println("We will select another number between range of 50 and 99, we select 67")
  fmt.Println("Add the factor to 67 (results in 133), now subtract 99, resulting in 34, 67 - 34 results in the number 33 we originally selected.")
  fmt.Println(67-(((99-33)+67)-99))
}