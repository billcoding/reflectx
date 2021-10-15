package reflectx

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type (
	_model struct {
		S string
		B bool
		I int64
		U uint64
		F float64
	}
	_testCase struct {
		expect interface{}
		src    interface{}
		ptr    *_model
	}
)

func TestValueString(t *testing.T) {
	cases := make([]*_testCase, 0)
	// test for String->String
	{
		cases = append(cases, &_testCase{"apple", "apple", &_model{}})
		cases = append(cases, &_testCase{"orange", "orange", &_model{}})
		cases = append(cases, &_testCase{"banana", "banana", &_model{}})
	}
	// test for Bool->String
	{
		cases = append(cases, &_testCase{"true", true, &_model{}})
		cases = append(cases, &_testCase{"false", false, &_model{}})
	}
	// test for Int->String
	{
		cases = append(cases, &_testCase{"100", 100, &_model{}})
		cases = append(cases, &_testCase{"100", int8(100), &_model{}})
		cases = append(cases, &_testCase{"100", int16(100), &_model{}})
		cases = append(cases, &_testCase{"100", int32(100), &_model{}})
		cases = append(cases, &_testCase{"100", int64(100), &_model{}})
	}
	// test for Uint->String
	{
		cases = append(cases, &_testCase{"100", uint(100), &_model{}})
		cases = append(cases, &_testCase{"100", uint8(100), &_model{}})
		cases = append(cases, &_testCase{"100", uint16(100), &_model{}})
		cases = append(cases, &_testCase{"100", uint32(100), &_model{}})
		cases = append(cases, &_testCase{"100", uint64(100), &_model{}})
	}
	// test for Float->String
	{
		cases = append(cases, &_testCase{"100", 100, &_model{}})
		cases = append(cases, &_testCase{"100.25", 100.25, &_model{}})
		cases = append(cases, &_testCase{"100", float32(100), &_model{}})
		cases = append(cases, &_testCase{"100.25", float32(100.25), &_model{}})
	}
	// test for Array->String
	{
		cases = append(cases, &_testCase{"[]", [0]string{}, &_model{}})
	}
	// test for Slice->String
	{
		cases = append(cases, &_testCase{"[]", []string{}, &_model{}})
	}
	// test for map->String
	{
		cases = append(cases, &_testCase{"map[apple:100]", map[string]string{"apple": "100"}, &_model{}})
	}
	// test for Struct->String
	{
		cases = append(cases, &_testCase{"{}", struct{}{}, &_model{}})
		cases = append(cases, &_testCase{"{100}", struct{ id int }{100}, &_model{}})
		cases = append(cases, &_testCase{"{100}", struct{ apple string }{"100"}, &_model{}})
	}
	for _, c := range cases {
		SetValue(reflect.ValueOf(c.src), reflect.ValueOf(c.ptr).Elem().FieldByName("S"))
		require.Equal(t, c.expect, c.ptr.S)
	}
}

func TestValueBool(t *testing.T) {
	cases := make([]*_testCase, 0)
	// test for String->Bool
	{
		cases = append(cases, &_testCase{true, "true", &_model{}})
		cases = append(cases, &_testCase{false, "false", &_model{}})
	}
	// test for Bool->Bool
	{
		cases = append(cases, &_testCase{true, true, &_model{}})
		cases = append(cases, &_testCase{false, false, &_model{}})
	}
	for _, c := range cases {
		SetValue(reflect.ValueOf(c.src), reflect.ValueOf(c.ptr).Elem().FieldByName("B"))
		require.Equal(t, c.expect, c.ptr.B)
	}
}

func TestValueInt(t *testing.T) {
	expect := int64(100)
	cases := make([]*_testCase, 0)
	// test for String->Int
	{
		cases = append(cases, &_testCase{expect, "100", &_model{}})
		cases = append(cases, &_testCase{expect, "100.0", &_model{}})
		cases = append(cases, &_testCase{expect, "100.25", &_model{}})
	}
	// test for Uint->Int
	{
		cases = append(cases, &_testCase{expect, uint(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint8(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint16(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint32(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint64(100), &_model{}})
	}
	// test for Float->Int
	{
		cases = append(cases, &_testCase{expect, 100.0, &_model{}})
		cases = append(cases, &_testCase{expect, 100.25, &_model{}})
		cases = append(cases, &_testCase{expect, float32(100), &_model{}})
		cases = append(cases, &_testCase{expect, float64(100), &_model{}})
	}
	// test for Int->Int
	{
		cases = append(cases, &_testCase{expect, 100, &_model{}})
	}
	for _, c := range cases {
		SetValue(reflect.ValueOf(c.src), reflect.ValueOf(c.ptr).Elem().FieldByName("I"))
		require.Equal(t, c.expect, c.ptr.I)
	}
}

func TestValueUint(t *testing.T) {
	expect := uint64(100)
	cases := make([]*_testCase, 0)
	// test for String->Uint
	{
		cases = append(cases, &_testCase{expect, "100", &_model{}})
		cases = append(cases, &_testCase{expect, "100.0", &_model{}})
		cases = append(cases, &_testCase{expect, "100.25", &_model{}})
	}
	// test for Int->Uint
	{
		cases = append(cases, &_testCase{expect, 100, &_model{}})
		cases = append(cases, &_testCase{expect, int8(100), &_model{}})
		cases = append(cases, &_testCase{expect, int16(100), &_model{}})
		cases = append(cases, &_testCase{expect, int32(100), &_model{}})
		cases = append(cases, &_testCase{expect, int64(100), &_model{}})
	}
	// test for Float->Uint
	{
		cases = append(cases, &_testCase{expect, 100.0, &_model{}})
		cases = append(cases, &_testCase{expect, 100.25, &_model{}})
		cases = append(cases, &_testCase{expect, float32(100), &_model{}})
		cases = append(cases, &_testCase{expect, float64(100), &_model{}})
	}
	// test for Uint->Uint
	{
		cases = append(cases, &_testCase{expect, uint(100), &_model{}})
	}
	for _, c := range cases {
		SetValue(reflect.ValueOf(c.src), reflect.ValueOf(c.ptr).Elem().FieldByName("U"))
		require.Equal(t, c.expect, c.ptr.U)
	}
}

func TestValueFloat(t *testing.T) {
	expect := float64(100)
	cases := make([]*_testCase, 0)
	// test for String->Float
	{
		cases = append(cases, &_testCase{expect, "100", &_model{}})
		cases = append(cases, &_testCase{expect, "100.0", &_model{}})
		cases = append(cases, &_testCase{expect, "100.00", &_model{}})
	}
	// test for Int->Float
	{
		cases = append(cases, &_testCase{expect, 100, &_model{}})
		cases = append(cases, &_testCase{expect, int8(100), &_model{}})
		cases = append(cases, &_testCase{expect, int16(100), &_model{}})
		cases = append(cases, &_testCase{expect, int32(100), &_model{}})
		cases = append(cases, &_testCase{expect, int64(100), &_model{}})
	}
	// test for Uint->Float
	{
		cases = append(cases, &_testCase{expect, uint(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint8(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint16(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint32(100), &_model{}})
		cases = append(cases, &_testCase{expect, uint64(100), &_model{}})
	}
	// test for Float->Float
	{
		cases = append(cases, &_testCase{expect, 100.0, &_model{}})
		cases = append(cases, &_testCase{expect, 100.00, &_model{}})
		cases = append(cases, &_testCase{expect, float32(100), &_model{}})
		cases = append(cases, &_testCase{expect, float64(100), &_model{}})
	}
	for _, c := range cases {
		SetValue(reflect.ValueOf(c.src), reflect.ValueOf(c.ptr).Elem().FieldByName("F"))
		require.Equal(t, c.expect, c.ptr.F)
	}
}
