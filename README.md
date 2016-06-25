# binarypatch

Storing & retrieving arbitrary data from program headers

## Usage

Appending some data to an OSX/Darwin program:

```
package main

import(
  "github.com/matiasinsaurralde/binarypatch"
)

func main() {
  b := binarypatch.New( "myprogram", "darwin" )
  index := b.Locate()
  b.Write( []byte("hello"), index )

  b.WriteFile( "myprogram_with_data")
}
```

```myprogram``` does the following:

```
package main
import(
  "fmt"
  "github.com/matiasinsaurralde/binarypatch"
)

func main() {
  b := binarypatch.ReadMyself(5, -1, "darwin")
  fmt.Println("Reading from myself: ", b, string(b) )
}
```

The output of the original program is:

```
% ./myprogram
Reading from myself:  [0 0 0 0 0] 
```

Then, the output of the "patched" program:
```
% ./myprogram_with_data
Reading from myself:  [104 101 108 108 111] hello
```

## Supported architectures

* OSX/Darwin
* Windows


## Related projects

* [binaryExpand](https://github.com/MrMaxB/binaryExpand)
* [go-liora](https://github.com/guitmz/go-liora)
* [payload-injector](https://github.com/sivu22/payload-injector)

## License

[MIT](https://github.com/matiasinsaurralde/binarypatch/blob/master/LICENSE)
