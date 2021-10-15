package reflectx

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestIsBasicType(t *testing.T) {
	require.Equal(t, true, IsBasicType(reflect.TypeOf(true)))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(false)))

	require.Equal(t, true, IsBasicType(reflect.TypeOf(100)))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(int8(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(int16(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(int32(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(int64(100))))

	require.Equal(t, true, IsBasicType(reflect.TypeOf(uint(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(uint8(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(uint16(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(uint32(100))))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(uint64(100))))

	require.Equal(t, true, IsBasicType(reflect.TypeOf(100.25)))
	require.Equal(t, true, IsBasicType(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsBasicType(reflect.TypeOf("apple")))
	require.Equal(t, false, IsBasicType(reflect.TypeOf("orange")))

	require.Equal(t, false, IsBasicType(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsBasicType(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsBasicType(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsBasicType(reflect.TypeOf([]string{})))

}

func TestIsBool(t *testing.T) {
	require.Equal(t, true, IsBool(reflect.TypeOf(true)))
	require.Equal(t, true, IsBool(reflect.TypeOf(false)))

	require.Equal(t, false, IsBool(reflect.TypeOf(100)))
	require.Equal(t, false, IsBool(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsBool(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsBool(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsBool(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsBool(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsBool(reflect.TypeOf("apple")))
	require.Equal(t, false, IsBool(reflect.TypeOf("orange")))

	require.Equal(t, false, IsBool(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsBool(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsBool(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsBool(reflect.TypeOf([]string{})))
}

func TestIsInt(t *testing.T) {
	require.Equal(t, false, IsInt(reflect.TypeOf(true)))
	require.Equal(t, false, IsInt(reflect.TypeOf(false)))

	require.Equal(t, true, IsInt(reflect.TypeOf(100)))
	require.Equal(t, true, IsInt(reflect.TypeOf(int8(100))))
	require.Equal(t, true, IsInt(reflect.TypeOf(int16(100))))
	require.Equal(t, true, IsInt(reflect.TypeOf(int32(100))))
	require.Equal(t, true, IsInt(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsInt(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsInt(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsInt(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsInt(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsInt(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsInt(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsInt(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsInt(reflect.TypeOf("apple")))
	require.Equal(t, false, IsInt(reflect.TypeOf("orange")))

	require.Equal(t, false, IsInt(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsInt(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsInt(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsInt(reflect.TypeOf([]string{})))
}

func TestIsUint(t *testing.T) {
	require.Equal(t, false, IsUint(reflect.TypeOf(true)))
	require.Equal(t, false, IsUint(reflect.TypeOf(false)))

	require.Equal(t, false, IsUint(reflect.TypeOf(100)))
	require.Equal(t, false, IsUint(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsUint(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsUint(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsUint(reflect.TypeOf(int64(100))))

	require.Equal(t, true, IsUint(reflect.TypeOf(uint(100))))
	require.Equal(t, true, IsUint(reflect.TypeOf(uint8(100))))
	require.Equal(t, true, IsUint(reflect.TypeOf(uint16(100))))
	require.Equal(t, true, IsUint(reflect.TypeOf(uint32(100))))
	require.Equal(t, true, IsUint(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsUint(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsUint(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsUint(reflect.TypeOf("apple")))
	require.Equal(t, false, IsUint(reflect.TypeOf("orange")))

	require.Equal(t, false, IsUint(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsUint(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsUint(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsUint(reflect.TypeOf([]string{})))
}

func TestIsFloat(t *testing.T) {
	require.Equal(t, false, IsFloat(reflect.TypeOf(true)))
	require.Equal(t, false, IsFloat(reflect.TypeOf(false)))

	require.Equal(t, false, IsFloat(reflect.TypeOf(100)))
	require.Equal(t, false, IsFloat(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsFloat(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsFloat(reflect.TypeOf(uint64(100))))

	require.Equal(t, true, IsFloat(reflect.TypeOf(100.25)))
	require.Equal(t, true, IsFloat(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsFloat(reflect.TypeOf("apple")))
	require.Equal(t, false, IsFloat(reflect.TypeOf("orange")))

	require.Equal(t, false, IsFloat(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsFloat(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsFloat(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsFloat(reflect.TypeOf([]string{})))
}

func TestIsString(t *testing.T) {
	require.Equal(t, false, IsString(reflect.TypeOf(true)))
	require.Equal(t, false, IsString(reflect.TypeOf(false)))

	require.Equal(t, false, IsString(reflect.TypeOf(100)))
	require.Equal(t, false, IsString(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsString(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsString(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsString(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsString(reflect.TypeOf(float32(100.25))))

	require.Equal(t, true, IsString(reflect.TypeOf("apple")))
	require.Equal(t, true, IsString(reflect.TypeOf("orange")))

	require.Equal(t, false, IsString(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsString(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsString(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsString(reflect.TypeOf([]string{})))
}

func TestIsStruct(t *testing.T) {
	require.Equal(t, false, IsStruct(reflect.TypeOf(true)))
	require.Equal(t, false, IsStruct(reflect.TypeOf(false)))

	require.Equal(t, false, IsStruct(reflect.TypeOf(100)))
	require.Equal(t, false, IsStruct(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsStruct(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsStruct(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsStruct(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsStruct(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsStruct(reflect.TypeOf("apple")))
	require.Equal(t, false, IsStruct(reflect.TypeOf("orange")))

	require.Equal(t, true, IsStruct(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsStruct(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsStruct(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsStruct(reflect.TypeOf([]string{})))
}

func TestIsPtr(t *testing.T) {
	require.Equal(t, false, IsPtr(reflect.TypeOf(true)))
	require.Equal(t, false, IsPtr(reflect.TypeOf(false)))

	require.Equal(t, false, IsPtr(reflect.TypeOf(100)))
	require.Equal(t, false, IsPtr(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsPtr(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsPtr(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsPtr(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsPtr(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsPtr(reflect.TypeOf("apple")))
	require.Equal(t, false, IsPtr(reflect.TypeOf("orange")))

	require.Equal(t, false, IsPtr(reflect.TypeOf(struct{}{})))
	require.Equal(t, true, IsPtr(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsPtr(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsPtr(reflect.TypeOf([]string{})))
}

func TestIsArray(t *testing.T) {
	require.Equal(t, false, IsArray(reflect.TypeOf(true)))
	require.Equal(t, false, IsArray(reflect.TypeOf(false)))

	require.Equal(t, false, IsArray(reflect.TypeOf(100)))
	require.Equal(t, false, IsArray(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsArray(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsArray(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsArray(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsArray(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsArray(reflect.TypeOf("apple")))
	require.Equal(t, false, IsArray(reflect.TypeOf("orange")))

	require.Equal(t, false, IsArray(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsArray(reflect.TypeOf(&struct{}{})))
	require.Equal(t, true, IsArray(reflect.TypeOf([0]string{})))
	require.Equal(t, false, IsArray(reflect.TypeOf([]string{})))
}

func TestIsSlice(t *testing.T) {
	require.Equal(t, false, IsSlice(reflect.TypeOf(true)))
	require.Equal(t, false, IsSlice(reflect.TypeOf(false)))

	require.Equal(t, false, IsSlice(reflect.TypeOf(100)))
	require.Equal(t, false, IsSlice(reflect.TypeOf(int8(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(int16(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(int32(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(int64(100))))

	require.Equal(t, false, IsSlice(reflect.TypeOf(uint(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(uint8(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(uint16(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(uint32(100))))
	require.Equal(t, false, IsSlice(reflect.TypeOf(uint64(100))))

	require.Equal(t, false, IsSlice(reflect.TypeOf(100.25)))
	require.Equal(t, false, IsSlice(reflect.TypeOf(float32(100.25))))

	require.Equal(t, false, IsSlice(reflect.TypeOf("apple")))
	require.Equal(t, false, IsSlice(reflect.TypeOf("orange")))

	require.Equal(t, false, IsSlice(reflect.TypeOf(struct{}{})))
	require.Equal(t, false, IsSlice(reflect.TypeOf(&struct{}{})))
	require.Equal(t, false, IsSlice(reflect.TypeOf([0]string{})))
	require.Equal(t, true, IsSlice(reflect.TypeOf([]string{})))
}
