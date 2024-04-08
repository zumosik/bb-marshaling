package bb

import (
	"reflect"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		sample interface{}
	}{
		{
			name: "Test 1",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  3,
				Flag: true,
			},
		},
		{
			name: "Test 2",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  8,
				Flag: false,
			},
		},
		{
			name: "Test 3",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  5,
				Flag: true,
			},
		},
		{
			name: "Test 4",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  6,
				Flag: false,
			},
		},
		{
			name: "Test 5",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  1,
				Flag: true,
			},
		},
		{
			name: "Test 6",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  3,
				Flag: true,
			},
		},
		{
			name: "Test 7",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  8,
				Flag: false,
			},
		},
		{
			name: "Test 8",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  5,
				Flag: true,
			},
		},
		{
			name: "Test 9",
			sample: struct {
				Num  int64
				Flag bool
			}{
				Num:  6,
				Flag: false,
			},
		},
		{
			name: "Test 10",
			sample: struct {
				Num int64
			}{
				Num: 1,
			},
		},
		// test cases with arrays
		{
			name: "Test 11",
			sample: struct {
				ArrBool []bool
			}{
				ArrBool: []bool{true, false, true},
			},
		},
		{
			name: "Test 12",
			sample: struct {
				ArrI32 []int32
			}{
				ArrI32: []int32{1, 2, 3},
			},
		},
		{
			name: "Test 13",
			sample: struct {
				ArrUI32 []uint32
			}{
				ArrUI32: []uint32{1, 2, 3},
			},
		},
		{
			name: "Test 14",
			sample: struct {
				ArrUint8 []uint8
			}{
				ArrUint8: []uint8{1, 2, 3},
			},
		},
		{
			name: "Test 15",
			sample: struct {
				ArrBool   [][]bool
				ArrUint16 [][]uint16
			}{
				ArrBool:   [][]bool{{true, false}, {true, false}},
				ArrUint16: [][]uint16{{1, 2}, {3, 4}},
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
