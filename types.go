package reflectx

import "reflect"

func IsBasicType(typ reflect.Type) bool {
	return IsBool(typ) || IsInt(typ) || IsUint(typ) || IsFloat(typ)
}

func IsBool(typ reflect.Type) bool {
	return typ.Kind() == reflect.Bool
}

func IsInt(typ reflect.Type) bool {
	return typ.Kind() == reflect.Int8 ||
		typ.Kind() == reflect.Int16 ||
		typ.Kind() == reflect.Int ||
		typ.Kind() == reflect.Int32 ||
		typ.Kind() == reflect.Int64
}

func IsUint(typ reflect.Type) bool {
	return typ.Kind() == reflect.Uint8 ||
		typ.Kind() == reflect.Uint16 ||
		typ.Kind() == reflect.Uint ||
		typ.Kind() == reflect.Uint32 ||
		typ.Kind() == reflect.Uint64
}

func IsFloat(typ reflect.Type) bool {
	return typ.Kind() == reflect.Float32 || typ.Kind() == reflect.Float64
}

func IsString(typ reflect.Type) bool {
	return typ.Kind() == reflect.String
}

func IsStruct(typ reflect.Type) bool {
	return typ.Kind() == reflect.Struct
}

func IsPtr(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr
}

func IsSlice(typ reflect.Type) bool {
	return typ.Kind() == reflect.Slice
}

func IsArray(typ reflect.Type) bool {
	return typ.Kind() == reflect.Array
}
