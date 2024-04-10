package bb

import (
	"math"
	"reflect"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		sample interface{}
	}{
		{
			name: "Test string, bool, uint32",
			sample: struct {
				Str  string
				Flag bool
				Num  uint32
			}{
				Num:  8,
				Flag: false,
				Str:  "hello world!",
			},
		},
		{
			name: "Test int64 + bool",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  6,
				Flag: false,
			},
		},
		{
			name: "Test int64",
			sample: struct {
				Num int64
			}{
				Num: 1,
			},
		},
		// test cases with arrays
		{
			name: "Test arrays of bool",
			sample: struct {
				ArrBool []bool
			}{
				ArrBool: []bool{true, false, true},
			},
		},
		{
			name: "Test arrays of int32",
			sample: struct {
				ArrI32 []int32
			}{
				ArrI32: []int32{1, 2, 3},
			},
		},
		{
			name: "Test arrays of strings",
			sample: struct {
				ArrStr []string
			}{
				ArrStr: []string{"hello", "world"},
			},
		},
		{
			name: "Test arrays of uint32",
			sample: struct {
				ArrUI32 []uint32
			}{
				ArrUI32: []uint32{1, 2, 3},
			},
		},
		{
			name: "Test arrays of uint8",
			sample: struct {
				ArrUint8 []uint8
			}{
				ArrUint8: []uint8{1, 2, 3},
			},
		},
		{
			name: "Test matrix",
			sample: struct {
				ArrBool   [][]bool
				ArrUint16 [][]uint16
			}{
				ArrBool:   [][]bool{{true, false}, {true, false}},
				ArrUint16: [][]uint16{{1, 2}, {3, 4}},
			},
		},
		{
			name: "Test structs",
			sample: struct {
				Data struct {
					Str  string
					Num  int32
					Flag bool
				}
				Info struct {
					Id uint8
				}
			}{
				Data: struct {
					Str  string
					Num  int32
					Flag bool
				}{
					Str:  "hello world!",
					Num:  -4040,
					Flag: true,
				},
				Info: struct {
					Id uint8
				}{
					Id: 90,
				},
			},
		},
		{
			name: "Test structs with arrays",
			sample: struct {
				Data struct {
					StrArr []string
				}
				Info struct {
					IdArr []uint8
				}
			}{
				Data: struct {
					StrArr []string
				}{
					StrArr: []string{"hello", "world"},
				},
				Info: struct {
					IdArr []uint8
				}{
					IdArr: []uint8{1, 2, 3},
				},
			},
		},
		{
			name: "Test all simple types",
			sample: struct {
				Str  string
				Flag bool
				I8   int8
				I16  int16
				I32  int32
				I64  int64
				U8   uint8
				U16  uint16
				U32  uint32
				U64  uint64
				F32  float32
				F64  float64
			}{
				Str:  "hello world!",
				Flag: true,
				I8:   -8,
				I16:  -16,
				I32:  -32,
				I64:  -64,
				U8:   8,
				U16:  16,
				U32:  32,
				U64:  64,
				F32:  32.32,
				F64:  math.Inf(2),
			},
		},
		{
			name: "Test a lot of types",
			sample: struct {
				SimpleTypes struct {
					Str  string
					Flag bool
					I8   int8
					I16  int16
					I32  int32
					I64  int64
					U8   uint8
					U16  uint16
					U32  uint32
					U64  uint64
					F32  float32
					F64  float64
				}
				StructsWithArrays struct {
					Data struct {
						StrArr []string
					}
				}
				Arrays struct {
					ArrBool   []bool
					ArrUint16 []uint16
				}
			}{
				SimpleTypes: struct {
					Str  string
					Flag bool
					I8   int8
					I16  int16
					I32  int32
					I64  int64
					U8   uint8
					U16  uint16
					U32  uint32
					U64  uint64
					F32  float32
					F64  float64
				}{
					Str:  "hello world!",
					Flag: true,
					I8:   -8,
					I16:  -16,
					I32:  -32,
					I64:  -64,
					U8:   8,
					U16:  16,
					U32:  32,
					U64:  64,
					F32:  32.32,
					F64:  math.Inf(2),
				},
				StructsWithArrays: struct {
					Data struct {
						StrArr []string
					}
				}{
					Data: struct {
						StrArr []string
					}{
						StrArr: []string{"hello", "world"},
					},
				},
				Arrays: struct {
					ArrBool   []bool
					ArrUint16 []uint16
				}{
					ArrBool:   []bool{true, false},
					ArrUint16: []uint16{1, 2},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := Marshall(test.sample)
			if err != nil {
				t.Errorf("Marshall() error = %v", err)
			}

			// Create a new instance of the same type as test.sample
			structDecode := reflect.New(reflect.TypeOf(test.sample)).Interface()

			err = Unmarshall(b, structDecode)
			if err != nil {
				t.Errorf("Unmarshall() error = %v", err)
			}

			if !reflect.DeepEqual(test.sample,
				reflect.ValueOf(structDecode).Elem().Interface()) {
				t.Errorf("Expected: %v, got: %v", test.sample, reflect.ValueOf(structDecode).Elem().Interface())
			} else {
				t.Logf("Test passed")
			}
		})
	}
}
