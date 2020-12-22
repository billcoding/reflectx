package reflectx

import (
	"reflect"
	"regexp"
	"strings"
)

//CreateFromTag
func CreateFromTag(structPtr, distPtr interface{}, alias, tag string) []*reflect.StructField {
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		panic("[CreateFromTag]structPtr of non-pointer type")
	}
	if reflect.TypeOf(structPtr).Elem().Kind() != reflect.Struct {
		panic("[CreateFromTag]structPtr Elem of non-struct type")
	}
	if reflect.TypeOf(distPtr).Kind() != reflect.Ptr {
		panic("[CreateFromTag]distPtr of non-ptr type")
	}
	if reflect.TypeOf(distPtr).Elem().Kind() != reflect.Slice {
		panic("[CreateFromTag]distPtr Elem of non-slice type")
	}
	if reflect.TypeOf(distPtr).Elem().Elem().Kind() != reflect.Ptr {
		panic("[CreateFromTag]distPtr Elem Elem of non-ptr type")
	}
	if reflect.TypeOf(distPtr).Elem().Elem().Elem().Kind() != reflect.Struct {
		panic("[CreateFromTag]distPtr of non-struct type")
	}
	distType := reflect.TypeOf(distPtr).Elem().Elem().Elem()
	structType := reflect.TypeOf(structPtr).Elem()
	aliasMap := make(map[string]string, 0)
	for i := 0; i < distType.NumField(); i++ {
		field := distType.Field(i)
		alias := strings.TrimSpace(field.Tag.Get(alias))
		if alias != "" {
			aliasMap[alias] = field.Name
		} else {
			aliasMap[field.Name] = field.Name
		}
	}
	distValues := reflect.ValueOf(distPtr).Elem()
	fields := make([]*reflect.StructField, 0)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		validateTag, have := field.Tag.Lookup(tag)
		if !have || validateTag == "" {
			continue
		}
		distValue := reflect.New(distType)
		re := regexp.MustCompile(`([a-z]+)\(([^()]+)\)`)
		tagMatchs := re.FindAllStringSubmatch(validateTag, -1)
		if len(tagMatchs) <= 0 {
			continue
		}
		for _, matchs := range tagMatchs {
			if len(matchs) < 3 {
				continue
			}
			_ = matchs[0] //fully match
			vname := strings.TrimSpace(matchs[1])
			vval := strings.TrimSpace(matchs[2])
			fieldName, have := aliasMap[vname]
			if have {
				SetValue(reflect.ValueOf(vval), distValue.Elem().FieldByName(fieldName))
			}
		}
		fields = append(fields, &field)
		distValues.Set(reflect.Append(distValues, distValue))
	}
	return fields
}
