package reflectx

import "reflect"

// IsBasicType ?
func IsBasicType(typ reflect.Type) bool {
	return IsIntType(typ) || IsFloatType(typ) || IsBoolType(typ) || IsStringType(typ)
}

// IsBoolType ?
func IsBoolType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Bool
}

// IsIntType ?
func IsIntType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Int8 ||
		typ.Kind() == reflect.Int16 ||
		typ.Kind() == reflect.Int ||
		typ.Kind() == reflect.Int32 ||
		typ.Kind() == reflect.Int64 ||
		typ.Kind() == reflect.Uint8 ||
		typ.Kind() == reflect.Uint16 ||
		typ.Kind() == reflect.Uint ||
		typ.Kind() == reflect.Uint32 ||
		typ.Kind() == reflect.Uint64
}

// IsFloatType ?
func IsFloatType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Float32 || typ.Kind() == reflect.Float64
}

// IsStringType ?
func IsStringType(typ reflect.Type) bool {
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
