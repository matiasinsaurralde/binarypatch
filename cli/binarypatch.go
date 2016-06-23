package main

import(
  "fmt"

  "github.com/matiasinsaurralde/binarypatch"
)

func main() {
  fmt.Println("binarypatch")

  b := binarypatch.New( "../sample/sample_darwin64", "darwin" )
  index := b.Locate()
  b.Write( []byte("hello"), index )

  b.WriteFile( "patched")

  b2 := binarypatch.New( "./patched", "darwin" )
  index2 := b2.Locate()
  data := b2.Read(index2, 5)

  fmt.Println("Stuff: ", data, string(data))

  myself := binarypatch.ReadMyself(5, -1, "darwin")

  fmt.Println("Myself: ", myself, string(myself) )

}
