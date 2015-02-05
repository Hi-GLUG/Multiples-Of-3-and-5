package main

import "fmt"

func sum_of_multiples( upto int ) int {

  // iterate over all numbers from 1 to `upto`
  // NO REDUCE??? srs???
  sum := 0
  for i := 0; i < upto; i++ {
    if (i % 3 == 0) || (i % 5 == 0) {
      sum += i
    }
  }
  return sum;
}

func main() {
  fmt.Println("Solution for 10 is", sum_of_multiples(10));
  fmt.Println("Solution for 20 is", sum_of_multiples(20));
}