package main

import "fmt";

func main() {
  longest, seed := 0, 0
  for i:=2; i < 1000000; i++ {
    length := len(collatz(i))
    if length > longest {
      seed, longest = i, length
      fmt.Printf("Current longest Collatz sequence is %d for %d\n", longest, i)
    }
  }
  fmt.Printf("The final answer is: The number below 1 million that produces the longest Collatz sequence is %d\n", seed)
}

func nextCollatz(n int) int {
  if(n%2 == 0) {
    return n/2
  }
  return 3*n + 1
}

func collatz(n int) []int {
  if n == 1 {
    return []int{1}
  }
  
  return append([]int{n}, collatz(nextCollatz(n))...)
}
