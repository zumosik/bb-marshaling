# BB marshaling
Library for encoding and decoding data into binary format
## Supported types for fields
- string
- bool
- int8, int16, int32, int64
- uint8, uint16, uint32, uint64
- float32, float64
- Array/Slice of any type upper
## Not supported types for fields
- int 
- uint
- pointer
- func
- interface{}
- map
- chan
## Basic example
### Encode
```go
package main

import (
	"github.com/zumosik/bb-marshaling"
	"os"
)

type TestStruct struct {
	Num         int64
	Flag        bool
	ArrOfFloats []float64
}

func main() {
	f, _ := os.Create("file.bin")

	enc := bb.NewEncoder(f)

	data := TestStruct{
		Num:         2,
		Flag:        true,
		ArrOfFloats: []float64{1.1, 2.2, 3.3, 4.4},
	}

	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
}
```
### Decode
```go
package main

import (
	"github.com/zumosik/bb-marshaling"
	"os"
)

type TestStruct struct {
	Num         int64
	Flag        bool
	ArrOfFloats []float64
}

func main() {
	f, _ := os.Open("file.bin")

	dec := bb.NewDecoder(f)

	data := TestStruct{
		Num:         2,
		Flag:        true,
		ArrOfFloats: []float64{1.1, 2.2, 3.3, 4.4},
	}

	err := dec.Decode(&data)
	if err != nil {
		panic(err)
	}
}
```