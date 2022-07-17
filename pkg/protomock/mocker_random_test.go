package protomock

import (
	"testing"

	"github.com/nlm/protoc-gen-mock/pkg/pb/testpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestRandomFieldValue(t *testing.T) {
	test := testpb.Test{}
	fields := test.ProtoReflect().Descriptor().Fields()

	var testCases = []struct {
		Name  string
		Value any
	}{
		{"field_double", float64(0)},
		{"field_float", float32(0)},
		{"field_int32", int32(0)},
		{"field_int64", int64(0)},
		{"field_uint32", uint32(0)},
		{"field_uint64", uint64(0)},
		{"field_sint32", int32(0)},
		{"field_sint64", int64(0)},
		{"field_fixed32", uint32(0)},
		{"field_fixed64", uint64(0)},
		{"field_sfixed32", int32(0)},
		{"field_sfixed64", int64(0)},
		{"field_bool", bool(false)},
		{"field_string", string("")},
		{"field_bytes", []byte(nil)},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			field := fields.ByName(protoreflect.Name(testCase.Name))
			assert.NotNil(t, field)
			v := randomFieldMocker(field)
			assert.NotNil(t, v)
			switch zeroValue := testCase.Value.(type) {
			case float32:
				assert.Equal(t, zeroValue, float32(field.Default().Float()))
				assert.NotEqual(t, zeroValue, v)
			case float64:
				assert.Equal(t, zeroValue, float64(field.Default().Float()))
				assert.NotEqual(t, zeroValue, v)
			case int32:
				assert.Equal(t, zeroValue, int32(field.Default().Int()))
				assert.NotEqual(t, zeroValue, v)
			case int64:
				assert.Equal(t, zeroValue, int64(field.Default().Int()))
				assert.NotEqual(t, zeroValue, v)
			case uint32:
				assert.Equal(t, zeroValue, uint32(field.Default().Uint()))
				assert.NotEqual(t, zeroValue, v)
			case uint64:
				assert.Equal(t, zeroValue, uint64(field.Default().Uint()))
				assert.NotEqual(t, zeroValue, v)
			case bool:
				// random bool can be false
				assert.Equal(t, zeroValue, bool(field.Default().Bool()))
			case string:
				assert.Equal(t, zeroValue, string(field.Default().String()))
				assert.NotEqual(t, zeroValue, v)
			case []byte:
				assert.Equal(t, zeroValue, []byte(field.Default().Bytes()))
				assert.NotEqual(t, zeroValue, v)
			}
		})
	}
}
