package bb

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

// Marshall writes marshalled representation of struct v to a byte slice.
//
// # v must be struct
func Marshall(v interface{}) ([]byte, error) {
	val := reflect.ValueOf(v)
	if !val.IsValid() || !(val.Kind() == reflect.Struct) {
		return nil, ErrInvalidType
	}

	var buf bytes.Buffer

	// Iterate through struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		err := encodeField(&buf, field)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// encodeField writes the marshalled representation of a single field to the underlying io.Writer.
func encodeField(buff *bytes.Buffer, field reflect.Value) error {
	var err error

	switch field.Type().Kind() {
	// --- String ---
	case reflect.String:
		err = encodeStr(buff, field.String())
	// --- Bool ---
	case reflect.Bool:
		err = encodeBool(buff, field.Bool())
	// --- Floats, Int, Uint
	case reflect.Float64, reflect.Float32,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		return binary.Write(buff, binary.BigEndian, field.Interface())
	// --- Slice/Array ---
	case reflect.Array, reflect.Slice:
		err := encodeArraySlice(buff, field)
		if err != nil {
			return err
		}
	default:
		return ErrInvalidValue
	}

	return err
}

func encodeArraySlice(buff *bytes.Buffer, field reflect.Value) error {
	length := field.Len()

	// Write length of array/slice as uint32
	err := encodeSize(buff, uint32(length))
	if err != nil {
		return err
	}

	// Encode each element of array/slice
	for i := 0; i < length; i++ {
		elem := field.Index(i)

		err := encodeField(buff, elem)
		if err != nil {
			return err
		}
	}

	return nil
}

func encodeStr(buff *bytes.Buffer, s string) error {
	// Encode string length as uint32 followed by string bytes
	err := encodeSize(buff, uint32(len(s)))
	if err != nil {
		return err
	}

	_, err = buff.Write([]byte(s))
	return err
}

func encodeBool(buff *bytes.Buffer, b bool) error {
	// Encode boolean as a single byte (0 for false, 1 for true)
	var val byte
	if b {
		val = 1
	}
	return binary.Write(buff, binary.BigEndian, val)
}

func encodeSize(buff *bytes.Buffer, i uint32) error {
	return binary.Write(buff, binary.BigEndian, i)
}
