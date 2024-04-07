package marshalling

import (
	"encoding/binary"
	"io"
	"reflect"
)

// A Decoder reads bytes from stream and writes them to values.
type Decoder struct {
	r io.Reader
}

// NewDecoder creates a new Decoder with the given io.Reader.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return ErrInvalidType
	}
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return ErrInvalidType
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			// Skip unexported fields
			continue
		}

		if err := d.DecodeField(field); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) DecodeField(field reflect.Value) error {

	switch field.Type().Kind() {
	// --- String ---
	case reflect.String:
		str, err := d.decodeStr()
		if err != nil {
			return err
		}
		field.SetString(str)
	// --- Bool ---
	case reflect.Bool:
		b, err := d.decodeBool()
		if err != nil {
			return err
		}
		field.SetBool(b)
	// --- Int ---
	case reflect.Int8:
		i, err := d.decodeInt8()
		if err != nil {
			return err
		}
		field.SetInt(int64(i))
	case reflect.Int16:
		i, err := d.decodeInt16()
		if err != nil {
			return err
		}
		field.SetInt(int64(i))
	case reflect.Int32:
		i, err := d.decodeInt32()
		if err != nil {
			return err
		}
		field.SetInt(int64(i))
	case reflect.Int64:
		i, err := d.decodeInt64()
		if err != nil {
			return err
		}
		field.SetInt(i)
	// --- Uint ---
	case reflect.Uint8:
		ui, err := d.decodeUint8()
		if err != nil {
			return err
		}
		field.SetUint(uint64(ui))
	case reflect.Uint16:
		ui, err := d.decodeUint16()
		if err != nil {
			return err
		}
		field.SetUint(uint64(ui))
	case reflect.Uint32:
		ui, err := d.decodeUint32()
		if err != nil {
			return err
		}
		field.SetUint(uint64(ui))
	case reflect.Uint64:
		ui, err := d.decodeUint64()
		if err != nil {
			return err
		}
		field.SetUint(ui)
	// --- Float ---
	case reflect.Float32:
		f, err := d.decodeFloat32()
		if err != nil {
			return err
		}
		field.SetFloat(float64(f))
	case reflect.Float64:
		f, err := d.decodeFloat64()
		if err != nil {
			return err
		}
		field.SetFloat(f)
	// --- Slice/Array ---
	case reflect.Array, reflect.Slice:
		if err := d.decodeArraySlice(field); err != nil {
			return err
		}
	default:
		return ErrInvalidValue
	}

	return nil
}

func (d *Decoder) decodeArraySlice(field reflect.Value) error {
	length, err := d.decodeUint32()
	if err != nil {
		return err
	}

	// Make a new slice with the given length and the same type as field's slice or array
	sliceType := reflect.SliceOf(field.Type().Elem())
	newSlice := reflect.MakeSlice(sliceType, int(length), int(length))

	for i := 0; i < int(length); i++ {
		elem := reflect.New(field.Type().Elem()).Elem() // Create a new element of the slice type
		if err := d.DecodeField(elem); err != nil {
			return err
		}
		newSlice.Index(i).Set(elem) // Set the decoded element to the slice
	}

	// Set the newly created slice to the field
	field.Set(newSlice)
	return nil
}

func (d *Decoder) decodeStr() (string, error) {
	length, err := d.decodeUint32()
	if err != nil {
		return "", err
	}
	strBytes := make([]byte, length)
	if _, err := io.ReadFull(d.r, strBytes); err != nil {
		return "", err
	}
	return string(strBytes), nil
}

func (d *Decoder) decodeBool() (bool, error) {
	var b [1]byte
	_, err := d.r.Read(b[:])
	if err != nil {
		return false, err
	}

	return b[0] != 0, nil
}

func (d *Decoder) decodeInt8() (int8, error) {
	var i int8
	if err := binary.Read(d.r, binary.BigEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func (d *Decoder) decodeInt16() (int16, error) {
	var i int16
	if err := binary.Read(d.r, binary.BigEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func (d *Decoder) decodeInt32() (int32, error) {
	var i int32
	if err := binary.Read(d.r, binary.BigEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func (d *Decoder) decodeInt64() (int64, error) {
	var i int64
	if err := binary.Read(d.r, binary.BigEndian, &i); err != nil {
		return 0, err
	}
	return i, nil
}

func (d *Decoder) decodeUint8() (uint8, error) {
	var u uint8
	if err := binary.Read(d.r, binary.BigEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (d *Decoder) decodeUint16() (uint16, error) {
	var u uint16
	if err := binary.Read(d.r, binary.BigEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (d *Decoder) decodeUint32() (uint32, error) {
	var u uint32
	if err := binary.Read(d.r, binary.BigEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (d *Decoder) decodeUint64() (uint64, error) {
	var u uint64
	if err := binary.Read(d.r, binary.BigEndian, &u); err != nil {
		return 0, err
	}
	return u, nil
}

func (d *Decoder) decodeFloat32() (float32, error) {
	var f float32
	if err := binary.Read(d.r, binary.BigEndian, &f); err != nil {
		return 0, err
	}
	return f, nil
}

func (d *Decoder) decodeFloat64() (float64, error) {
	var f float64
	if err := binary.Read(d.r, binary.BigEndian, &f); err != nil {
		return 0, err
	}
	return f, nil
}
