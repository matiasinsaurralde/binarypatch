package binarypatch

import(
  "io/ioutil"
  "bytes"
  // "fmt"
)

type Binarypatch struct {
  Filename string
  Filedata []byte
  Arch string
}

func New( filename string, arch string ) *Binarypatch {
  var err error
  b := Binarypatch{Filename: filename, Arch: arch}
  b.Filedata, err = ioutil.ReadFile( filename)

  if err != nil {
    panic(err)
  }

  return &b
}

func( b *Binarypatch ) Locate() ( index int) {
  switch b.Arch {
  case "darwin":
    str := []byte("__TEXT")
    for i, v := range b.Filedata {
      n := 0
      byteGroup := []byte{v}
      for n <= len(str) && i < len(b.Filedata) - len(str) {
        byteGroup = append(byteGroup, b.Filedata[i+n])
        n++
      }

      if bytes.Index( byteGroup, str ) == 0 {
        index = i + len(str)
        break
      }
    }
  }
  return index
}

func( b *Binarypatch) Write( data []byte, index int ) {
  for i, v := range data {
    position := index + i
    b.Filedata[position] = v
  }
}
