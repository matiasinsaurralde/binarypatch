package main

import(
  "fmt"

  "github.com/matiasinsaurralde/binarypatch"
)

func main() {
  fmt.Println("Hello")

  myself := binarypatch.ReadMyself(5, -1, "darwin")
  fmt.Println("Reading from myself: ", myself, string(myself) )
}
