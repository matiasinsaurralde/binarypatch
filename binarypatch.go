package binarypatch

import(
  "io/ioutil"
  "bytes"

  "github.com/kardianos/osext"
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

func( b *Binarypatch) WriteFile( name string ) (err error) {
  if name == "" {
    err = ioutil.WriteFile( b.Filename, b.Filedata, 0744 )
  } else {
    err = ioutil.WriteFile( name, b.Filedata, 0744 )
  }
  return
}

func ( b *Binarypatch) Read( index int, length int ) []byte {
  var data []byte
  data = b.Filedata[index:index+length]
  return data
}

func ReadMyself( length int, index int, arch string ) []byte {
  filename, err := osext.Executable()
  if err != nil {
    panic(err)
  }
  b := New( filename, arch )
  if index == -1 {
    index = b.Locate()
  }
  data := b.Read( index, length )
  return data
}
