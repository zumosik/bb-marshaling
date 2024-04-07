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
