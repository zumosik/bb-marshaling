package bb

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		sample interface{}
	}{
		{
			name: "TestStructFloat",
			sample: struct {
				Float float32
			}{
				Float: 1.552,
			},
		},
		{
			name: "TestStructInt",
			sample: struct {
				I32 int32
			}{
				I32: 42,
			},
		},
		{
			name: "TestStructString",
			sample: struct {
				Str  string
				Str1 string
				Str2 string
				Str3 string
			}{
				Str:  "hello",
				Str1: "world",
				Str2: "c++",
				Str3: "go",
			},
		},
		{
			name: "TestStructBool",
			sample: struct {
				Flag  bool
				Flag1 bool
				Flag2 bool
			}{
				Flag:  true,
				Flag1: false,
				Flag2: false,
			},
		},
		{
			name: "TestStructArr",
			sample: struct {
				ArrBool  []bool
				ArrFloat []float32
				ArrStr   []string

				ArrI32 []int32
				ArrI64 []int64
				ArrI16 []int16
				ArrI8  []int8
			}{
				ArrBool:  []bool{true, false, false},
				ArrFloat: []float32{1.1, 2.2, 3.3},
				ArrStr: []string{
					"hello", "world", "C++", "go",
				},
				ArrI32: []int32{1, 2, 3},
				ArrI64: []int64{1, 2, 3},
				ArrI16: []int16{1, 2, 3},
				ArrI8:  []int8{1, 2, 3},
			},
		},
		{
			name: "TestStructMoreFields",
			sample: struct {
				F32    []float32
				StrArr []string
			}{
				F32:    []float32{1.1, 2.2, 3.3},
				StrArr: []string{"hello", "world", "C++", "go"},
			},
		},
		{
			name: "TestAnotherStruct",
			sample: struct {
				S string
				B bool

				I8  int8
				I16 int16
				I32 int32
				I64 int64

				F32 float32
				F64 float64
			}{
				S:   "hello world",
				B:   false,
				I8:  -8,
				I16: -99,
				I32: 25,
				I64: 445,
				F32: -0.066,
				F64: -0.2577,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create an encoder
			var buf bytes.Buffer
			encoder := &Encoder{w: &buf}

			// Encode the sample struct
			if err := encoder.Encode(test.sample); err != nil {
				t.Fatalf("error encoding struct: %v", err)
			}

			t.Log(buf)

			// Create a new instance of the same type as test.sample
			decodedStruct := reflect.New(reflect.TypeOf(test.sample)).Interface()

			// Create a decoder
			decoder := &Decoder{r: bytes.NewBuffer(buf.Bytes())}

			// Decode into the new struct
			if err := decoder.Decode(decodedStruct); err != nil {
				t.Fatalf("error decoding struct: %v", err)
			}

			// Check if the decoded struct is the same as the original struct
			if !reflect.DeepEqual(test.sample, reflect.ValueOf(decodedStruct).Elem().Interface()) {
				t.Errorf("decoded struct does not match original struct.\nOriginal: %+v\nDecoded: %+v", test.sample, decodedStruct)
			}
		})
	}
}

/*
ArrBool  []bool
ArrI32   []int32
ArrUI32  []uint32
ArrUint8 []uint8
*/

// {[0 0 0 3
//1 0 0
//
//0 0 0 8
//0 0 0 5
//0 0 0 5
//0 0 0 5
//255 255 255 203
//0 0 0 3
//0 0 0 3
//0 0 0 3
//0 0 0 3
//
//
//0 0 0 5
//0 0 0 8
//0 0 0 8
//0 0 0 8
//0 0 0 8
//0 0 0 8

//0 0 0 6 1 2 3 4 5 6 0 0 0 3 63 140 204 205 64 12 204 205 64 83 51 51 0 0 0 4 0 0 0 5 104 101 108 108 111 0 0 0 5 119 111 114 108 100 0 0 0 3 67 43 43 0 0 0 2 103 111]
