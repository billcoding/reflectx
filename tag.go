package reflectx

import (
	"reflect"
	"regexp"
	"strings"
)

// ParseTag parse struct tag
func ParseTag(structPtr, tagPtr interface{}, alias, tag string, recursive bool) (structFields []*reflect.StructField, structFieldValues []*reflect.Value, items []interface{}) {
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		panic("[validator.createFromTag]structPtr of non-pointer type")
	}
	if reflect.TypeOf(structPtr).Elem().Kind() != reflect.Struct {
		panic("[validator.createFromTag]structPtr Elem of non-struct type")
	}
	if reflect.TypeOf(tagPtr).Kind() != reflect.Ptr {
		panic("[validator.createFromTag]tagPtr of non-pointer type")
	}
	if reflect.TypeOf(tagPtr).Elem().Kind() != reflect.Struct {
		panic("[validator.createFromTag]tagPtr Elem of non-struct type")
	}
	structType := reflect.TypeOf(structPtr).Elem()
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
	tagItems := make([]interface{}, 0)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		valueOf := reflect.ValueOf(structPtr)
		if valueOf.IsNil() || valueOf.IsZero() {
			continue
		}
		fieldValue := valueOf.Elem().Field(i)
		tagStr, have := field.Tag.Lookup(tag)
		if !have || tagStr == "" {
			continue
		}
		tagItem := reflect.New(tagType)
		re := regexp.MustCompile(`([a-z]+)\(([^()]+)\)`)
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

		if IsBasicType(field.Type) {
			structFields = append(structFields, &field)
			structFieldValues = append(structFieldValues, &fieldValue)
			tagItems = append(tagItems, tagItem.Interface())
		}

		if !recursive {
			continue
		}

		switch {
		case field.Type.Kind() == reflect.Struct:
			// Struct{}
			fs, vs, ts := ParseTag(fieldValue.Addr().Interface(), tagPtr, alias, tag, recursive)
			structFields = append(structFields, fs...)
			structFieldValues = append(structFieldValues, vs...)
			tagItems = append(tagItems, ts...)
		case field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct:
			// *Struct{}
			fs, vs, ts := ParseTag(fieldValue.Interface(), tagPtr, alias, tag, recursive)
			structFields = append(structFields, fs...)
			structFieldValues = append(structFieldValues, vs...)
			tagItems = append(tagItems, ts...)
		case (field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Array) && field.Type.Elem().Kind() == reflect.Struct:
			// []Struct{}
			valLen := fieldValue.Elem().Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Index(i).Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				tagItems = append(tagItems, ts...)
			}
		case field.Type.Kind() == reflect.Ptr && (field.Type.Elem().Kind() == reflect.Slice || field.Type.Elem().Kind() == reflect.Array) && field.Type.Elem().Elem().Kind() == reflect.Struct:
			// *[]Struct{}
			valLen := fieldValue.Elem().Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Elem().Index(i).Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				tagItems = append(tagItems, ts...)
			}
		case field.Type.Kind() == reflect.Ptr && (field.Type.Elem().Kind() == reflect.Slice || field.Type.Elem().Kind() == reflect.Array) && field.Type.Elem().Elem().Kind() == reflect.Ptr && field.Type.Elem().Elem().Elem().Kind() == reflect.Struct:
			// *[]*Struct{}
			valLen := fieldValue.Elem().Len()
			for i := 0; i < valLen; i++ {
				fs, vs, ts := ParseTag(fieldValue.Elem().Index(i).Elem().Interface(), tagPtr, alias, tag, recursive)
				structFields = append(structFields, fs...)
				structFieldValues = append(structFieldValues, vs...)
				tagItems = append(tagItems, ts...)
			}
		}
	}
	return structFields, structFieldValues, tagItems
}
