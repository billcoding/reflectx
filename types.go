package reflectx

import "reflect"

// IsBasicType ?
func IsBasicType(typ reflect.Type) bool {
	return IsInt(typ) || IsFloat(typ) || IsBool(typ) || IsString(typ)
}

// IsBool ?
func IsBool(typ reflect.Type) bool {
	return typ.Kind() == reflect.Bool
}

// IsInt ?
func IsInt(typ reflect.Type) bool {
	return typ.Kind() == reflect.Int8 ||
		typ.Kind() == reflect.Int16 ||
		typ.Kind() == reflect.Int ||
		typ.Kind() == reflect.Int32 ||
		typ.Kind() == reflect.Int64
}

// IsUint ?
func IsUint(typ reflect.Type) bool {
	return typ.Kind() == reflect.Uint8 ||
		typ.Kind() == reflect.Uint16 ||
		typ.Kind() == reflect.Uint ||
		typ.Kind() == reflect.Uint32 ||
		typ.Kind() == reflect.Uint64
}

// IsFloat ?
func IsFloat(typ reflect.Type) bool {
	return typ.Kind() == reflect.Float32 || typ.Kind() == reflect.Float64
}

// IsString ?
func IsString(typ reflect.Type) bool {
	return typ.Kind() == reflect.String
}

// IsStruct ?
func IsStruct(typ reflect.Type) bool {
	return typ.Kind() == reflect.Struct
}

// IsPtr ?
func IsPtr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

// IsSlice ?
func IsSlice(typ reflect.Type) bool {
	return typ.Kind() == reflect.Slice
}

// IsArray ?
func IsArray(typ reflect.Type) bool {
	return typ.Kind() == reflect.Array
}
