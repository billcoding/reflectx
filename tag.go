package reflectx

import (
	"reflect"
	"regexp"
	"strings"
)

func ParseTag(structPtr, tagPtr interface{}, alias, tag string, recursive bool) ([]*reflect.StructField, []*reflect.Value, []interface{}) {
	// FIXED: when alias contains `_`, not match!!!
	return ParseTagWithRe(structPtr, tagPtr, alias, tag, recursive, `([a-zA-Z0-9_]+)\(([^()]+)\)`)
}

func ParseTagWithRe(structPtr, tagPtr interface{}, alias, tag string, recursive bool, re string) ([]*reflect.StructField, []*reflect.Value, []interface{}) {
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		panic("structPtr of non-pointer type")
	}
	if reflect.TypeOf(structPtr).Elem().Kind() != reflect.Struct {
		panic("structPtr Elem of non-struct type")
	}
	if reflect.TypeOf(tagPtr).Kind() != reflect.Ptr {
		panic("tagPtr of non-pointer type")
	}
	if reflect.TypeOf(tagPtr).Elem().Kind() != reflect.Struct {
		panic("tagPtr Elem of non-struct type")
	}
	structType := reflect.TypeOf(structPtr).Elem()
	structValue := reflect.ValueOf(structPtr)
	tagType := reflect.TypeOf(tagPtr).Elem()
	aliasMap := make(map[string]string, 0)
	for i := 0; i < tagType.NumField(); i++ {
		field := tagType.Field(i)
		aliasStr := strings.TrimSpace(field.Tag.Get(alias))
		if aliasStr != "" {
			aliasMap[aliasStr] = field.Name
		} else {
			aliasMap[field.Name] = field.Name
		}
	}
	structFields := make([]*reflect.StructField, 0)
	structFieldValues := make([]*reflect.Value, 0)
	items := make([]interface{}, 0)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		var fieldValue reflect.Value
		if structValue.IsNil() {
			fieldValue = reflect.New(field.Type).Elem()
		} else {
			fieldValue = structValue.Elem().Field(i)
		}
		tagStr, have := field.Tag.Lookup(tag)
		if !have || tagStr == "" {
			continue
		}
		tagItem := reflect.New(tagType)
		re := regexp.MustCompile(re)
		tagMatches := re.FindAllStringSubmatch(tagStr, -1)
		if len(tagMatches) <= 0 {
			continue
		}
		for _, matches := range tagMatches {
			if len(matches) < 3 {
				continue
			}
			_ = matches[0]
			name := strings.TrimSpace(matches[1])
			val := strings.TrimSpace(matches[2])
			fieldName, have := aliasMap[name]
			if have {
				SetValue(reflect.ValueOf(val), tagItem.Elem().FieldByName(fieldName))
			}
		}
		structFields = append(structFields, &field)
		structFieldValues = append(structFieldValues, &fieldValue)
		items = append(items, tagItem.Interface())

		if !recursive {
			continue
		}

		switch {
		case IsStruct(field.Type):
			// Struct{}
			fs, vs, ts := ParseTag(fieldValue.Addr().Interface(), tagPtr, alias, tag, recursive)
			structFields = append(structFields, fs...)
			structFieldValues = append(structFieldValues, vs...)
			items = append(items, ts...)
		case IsPtr(field.Type) && IsStruct(field.Type.Elem()):
			// *Struct{}
			fs, vs, ts := ParseTag(fieldValue.Interface(), tagPtr, alias, tag, recursive)
			structFields = append(structFields, fs...)
			structFieldValues = append(structFieldValues, vs...)
			items = append(items, ts...)
		case (IsSlice(field.Type) || IsArray(field.Type)) && IsStruct(field.Type.Elem()):
			// []Struct{}
			valLen := fieldValue.Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Index(i).Addr().Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				items = append(items, ts...)
			}
		case (IsSlice(field.Type) || IsArray(field.Type)) && IsPtr(field.Type.Elem()) && IsStruct(field.Type.Elem().Elem()):
			// []*Struct{}
			valLen := fieldValue.Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Index(i).Elem().Addr().Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				items = append(items, ts...)
			}
		case IsPtr(field.Type) && (IsSlice(field.Type.Elem()) || IsArray(field.Type.Elem())) && IsStruct(field.Type.Elem().Elem()):
			// *[]Struct{}
			valLen := fieldValue.Elem().Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Elem().Index(i).Addr().Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				items = append(items, ts...)
			}
		case IsPtr(field.Type) && (IsSlice(field.Type.Elem()) || IsArray(field.Type.Elem())) && IsPtr(field.Type.Elem().Elem()) && IsStruct(field.Type.Elem().Elem().Elem()):
			// *[]*Struct{}
			valLen := fieldValue.Elem().Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Elem().Index(i).Elem().Addr().Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				items = append(items, ts...)
			}
		}
	}
	return structFields, structFieldValues, items
}
