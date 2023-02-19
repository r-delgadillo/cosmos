package jsonutil

import (
	"fmt"
	"io"
	"math"
	"reflect"

	"github.com/fatih/structs"
	"github.com/goccy/go-json"
)

// SafeMarshal marshals a value to a JSON string, handling various
// serialization concerns in the process.
func SafeMarshal(val interface{}) ([]byte, error) {
	str, err := json.MarshalWithOption(val, json.DisableHTMLEscape())
	if e, ok := err.(*json.UnsupportedValueError); ok && e.Value.Kind() == reflect.Float64 {
		str, err = json.MarshalWithOption(replaceUnsupportedFloats(val), json.DisableHTMLEscape())
	}
	return str, err
}

// SafeEncoder is a JSON encoder that handles various serialization concerns.
type SafeEncoder json.Encoder

// NewSafeEncoder creates a JSON encoder that handles various serialization
// concerns.
func NewSafeEncoder(w io.Writer) *SafeEncoder {
	return (*SafeEncoder)(json.NewEncoder(w))
}

// Encode encodes a value using a JSON encoder, handling various serialization
// concerns in the process.
func (se *SafeEncoder) Encode(val interface{}) error {
	err := (*json.Encoder)(se).EncodeWithOption(val, json.DisableHTMLEscape())
	if e, ok := err.(*json.UnsupportedValueError); ok && e.Value.Kind() == reflect.Float64 {
		err = (*json.Encoder)(se).EncodeWithOption(replaceUnsupportedFloats(val), json.DisableHTMLEscape())
	}
	return err
}

// Scans a to-be-marshalled object and replaces unsupported float64 values.
// Since this function is potentially expensive, only call it if we already
// know there are unsupported floats present.
func replaceUnsupportedFloats(value interface{}) interface{} {
	r := reflect.ValueOf(value)
	switch r.Kind() {

	case reflect.Map:
		clean := make(map[string]interface{}, r.Len())
		i := r.MapRange()
		for i.Next() {
			clean[fmt.Sprint(i.Key())] = replaceUnsupportedFloats(i.Value().Interface())
		}
		return clean

	case reflect.Struct:
		// Coerce the struct to a map[string]interface{} using json tags
		mapper := structs.New(r.Interface())
		mapper.TagName = "json"
		clean := mapper.Map()
		for k, v := range clean {
			clean[k] = replaceUnsupportedFloats(v)
		}
		return clean

	case reflect.Slice:
		clean := make([]interface{}, r.Len())
		for i := 0; i < r.Len(); i++ {
			clean[i] = replaceUnsupportedFloats(r.Index(i).Interface())
		}
		return clean

	case reflect.Float64:
		float := r.Float()
		if math.IsInf(float, 0) || math.IsNaN(float) {
			return nil
		}
		return float

	default:
		return r.Interface()

	}
}
