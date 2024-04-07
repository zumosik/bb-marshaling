package marshalling

import (
	"encoding/binary"
	"io"
	"reflect"
)

// An Encoder writes bytes values to an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new Encoder with the given io.Writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes the marshalled representation of v to the underlying io.Writer.
func (e *Encoder) Encode(v interface{}) error {
	val := reflect.ValueOf(v)
	if !val.IsValid() || !(val.Kind() == reflect.Struct) {
		return ErrInvalidType
	}

	// Iterate through struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		err := e.EncodeField(field)
		if err != nil {
			return err
		}
	}

	return nil
}

// EncodeField writes the marshalled representation of a single field to the underlying io.Writer.
func (e *Encoder) EncodeField(field reflect.Value) error {
	var err error

	switch field.Type().Kind() {
	// --- String ---
	case reflect.String:
		err = e.encodeStr(field.String())
	// --- Bool ---
	case reflect.Bool:
		err = e.encodeBool(field.Bool())
	// --- Int ---
	case reflect.Int8:
		err = e.encodeInt8(int8(field.Int()))
	case reflect.Int16:
		err = e.encodeInt16(int16(field.Int()))
	case reflect.Int32:
		err = e.encodeInt32(int32(field.Int()))
	case reflect.Int64:
		err = e.encodeInt64(field.Int())
	// --- Uint ---
	case reflect.Uint8:
		err = e.encodeUint8(uint8(field.Uint()))
	case reflect.Uint16:
		err = e.encodeUint16(uint16(field.Uint()))
	case reflect.Uint32:
		err = e.encodeUint32(uint32(field.Uint()))
	case reflect.Uint64:
		err = e.encodeUint64(field.Uint())
	// --- Float ---
	case reflect.Float32:
		err = e.encodeFloat32(float32(field.Float()))
	case reflect.Float64:
		err = e.encodeFloat64(field.Float())
	// --- Slice/Array ---
	case reflect.Array, reflect.Slice:
		err := e.encodeArraySlice(field)
		if err != nil {
			return err
		}
	default:
		return ErrInvalidValue
	}

	return err
}

func (e *Encoder) encodeArraySlice(field reflect.Value) error {
	length := field.Len()

	// Write length of array/slice as uint32
	err := e.encodeUint32(uint32(length))
	if err != nil {
		return err
	}

	// Encode each element of array/slice
	for i := 0; i < length; i++ {
		elem := field.Index(i)

		err := e.EncodeField(elem)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encodeStr(s string) error {
	// Encode string length as uint32 followed by string bytes
	err := e.encodeUint32(uint32(len(s)))
	if err != nil {
		return err
	}

	_, err = e.w.Write([]byte(s))
	return err
}

func (e *Encoder) encodeBool(b bool) error {
	// Encode boolean as a single byte (0 for false, 1 for true)
	var val byte
	if b {
		val = 1
	}
	return binary.Write(e.w, binary.BigEndian, val)
}

func (e *Encoder) encodeInt8(i int8) error {
	// Encode int8
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeInt16(i int16) error {
	// Encode int16
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeInt32(i int32) error {
	// Encode int32
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeInt64(i int64) error {
	// Encode int64
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeUint8(i uint8) error {
	// Encode uint8
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeUint16(i uint16) error {
	// Encode uint16
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeUint32(i uint32) error {
	// Encode uint32
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeUint64(i uint64) error {
	// Encode uint64
	return binary.Write(e.w, binary.BigEndian, i)
}

func (e *Encoder) encodeFloat32(f float32) error {
	// Encode float32
	return binary.Write(e.w, binary.BigEndian, f)
}

func (e *Encoder) encodeFloat64(f float64) error {
	// Encode float64
	return binary.Write(e.w, binary.BigEndian, f)
}
