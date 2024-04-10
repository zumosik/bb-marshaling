package bb

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
)

// Unmarshall reads marshalled representation of struct v from a byte slice.
//
// # v must be pointer to struct
func Unmarshall(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)

	valPtr := reflect.ValueOf(v)
	if valPtr.Kind() != reflect.Ptr || valPtr.IsNil() {
		return ErrInvalidType
	}
	val := valPtr.Elem()
	if val.Kind() != reflect.Struct {
		return ErrInvalidType
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			// Skip unexported fields
			continue
		}

		if err := decodeField(buf, field); err != nil {
			return err
		}
	}

	return nil
}
func decodeField(buf *bytes.Buffer, field reflect.Value) error {
	switch field.Type().Kind() {
	// --- String ---
	case reflect.String:
		str, err := decodeStr(buf)
		if err != nil {
			return err
		}
		field.SetString(str)
	// --- Bool ---
	case reflect.Bool:
		b, err := decodeBool(buf)
		if err != nil {
			return err
		}
		field.SetBool(b)
	// --- Float, Int, Uint ---
	case reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		if err := binary.Read(buf, binary.BigEndian, field.Addr().Interface()); err != nil {
			return err
		}
	// --- Slice/Array ---
	case reflect.Array, reflect.Slice:
		if err := decodeArraySlice(buf, field); err != nil {
			return err
		}
	// --- Struct ---
	case reflect.Struct:
		for i := 0; i < field.NumField(); i++ {
			if err := decodeField(buf, field.Field(i)); err != nil {
				return err
			}
		}
	default:
		return ErrInvalidValue
	}

	return nil
}

func decodeArraySlice(buf *bytes.Buffer, field reflect.Value) error {
	length, err := decodeSize(buf)
	if err != nil {
		return err
	}

	// Make a new slice with the given length and the same type as field's slice or array
	sliceType := reflect.SliceOf(field.Type().Elem())
	newSlice := reflect.MakeSlice(sliceType, int(length), int(length))

	for i := 0; i < int(length); i++ {
		elem := reflect.New(field.Type().Elem()).Elem() // Create a new element of the slice type
		if err := decodeField(buf, elem); err != nil {
			return err
		}
		newSlice.Index(i).Set(elem) // Set the decoded element to the slice
	}

	// Set the newly created slice to the field
	field.Set(newSlice)
	return nil
}

func decodeStr(buf *bytes.Buffer) (string, error) {
	length, err := decodeSize(buf)
	if err != nil {
		return "", err
	}
	strBytes := make([]byte, length)
	if _, err := io.ReadFull(buf, strBytes); err != nil {
		return "", err
	}
	return string(strBytes), nil
}

func decodeBool(buf *bytes.Buffer) (bool, error) {
	var b [1]byte
	_, err := buf.Read(b[:])
	if err != nil {
		return false, err
	}

	return b[0] != 0, nil
}

func decodeSize(buf *bytes.Buffer) (uint32, error) {
	var i uint32
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}
