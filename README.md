# BB marshaling
> Library for marshalling data into binary format  

[![Built with Go](https://img.shields.io/badge/Built%20with-Go-00ADD8.svg)](https://golang.org/)
![GitHub License](https://img.shields.io/github/license/zumosik/bb-marshaling)

This library provides functionality to convert (marshal) a Go struct into a binary format and to convert (unmarshal) binary data back into a Go struct. The aim is to ensure that the binary representation is as concise and straightforward as possible.
## Supported types for fields
- string
- bool
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- Array/Slice of any type upper
- Nested struct with fields of any type upper
## Not supported types for fields
- pointer
- func
- interface{}
- map
- chan
## Basic example
### Marshall
```go
package main

import (
	"fmt"
	"github.com/zumosik/bb-marshaling"
)

type TestStruct struct {
	Num  int64
	Flag bool
}

func main() {
	binData, err := bb.Marshall(TestStruct{Num: 3, Flag: true})
	if err != nil {
		// handle error
	}

	fmt.Println(binData) // do something with the binary data
}

```
### Unmarshall
```go
package main

import (
	"fmt"
	"github.com/zumosik/bb-marshaling"
)

type TestStruct struct {
	Num  int64
	Flag bool
}

func main() {
	binData := []byte{0, 0, 0, 0, 0, 0, 0, 3, 1} // data from previous example
	var v TestStruct

	err := bb.Unmarshall(binData, &v)
	if err != nil {
		// handle error
	}

	fmt.Println(v) // // do something with the unmarshalled data
}

```