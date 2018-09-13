// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.6 Test a simple library package main function

package main

import (
  "fmt"
  "test1"
)

// test functions DoubleVariable and TripleVariable by passing 8 in
func main() {
  out := test1.DoubleValue(8)            // call function in package
  fmt.Printf("Double, out = %d\n", out)     // should print "..out = 16"
//  out := test1.TripleVariable(8), call function in package
//  fmt.Printf("Triple, out = %d\n, out")   // should print "..out = 16"
}
