package reflectx

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "[reflectx]", log.LstdFlags)

// SetValue from sourceValue to distValue
func SetValue(sourceValue reflect.Value, distValue reflect.Value) {
	switch distValue.Kind() {
	case reflect.Bool:
		if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
			boolVal, err := strconv.ParseBool(sourceValue.String())
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			} else {
				distValue.SetBool(boolVal)
			}
		} else if sourceValue.Type().Kind() == reflect.Bool {
			distValue.Set(sourceValue)
		}
	case reflect.String:
		if sourceValue.CanInterface() {
			distValue.SetString(fmt.Sprintf("%v", sourceValue.Interface()))
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
			intVal, err := strconv.ParseInt(sourceValue.String(), 10, 64)
			if err != nil {
				logger.Println(fmt.Sprintf("[setValue]%v", err))
			} else {
				distValue.SetInt(intVal)
			}
		} else if sourceValue.Type().Kind() == reflect.Float32 || sourceValue.Type().Kind() == reflect.Float64 {
			distValue.SetInt(int64(sourceValue.Float()))
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
			uintVal, err := strconv.ParseUint(sourceValue.String(), 10, 64)
			if err != nil {
				logger.Println(fmt.Sprintf("[setValue]%v", err))
			} else {
				distValue.SetUint(uintVal)
			}
		} else if sourceValue.Type().Kind() == reflect.Float32 || sourceValue.Type().Kind() == reflect.Float64 {
			distValue.SetUint(uint64(sourceValue.Float()))
		}
	case reflect.Float32, reflect.Float64:
		if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
			floatVal, err := strconv.ParseFloat(sourceValue.String(), 32)
			if err != nil {
				logger.Println(fmt.Sprintf("[setValue]%v", err))
			} else {
				distValue.SetFloat(floatVal)
			}
		} else if sourceValue.Type().Kind() == reflect.Float32 || sourceValue.Type().Kind() == reflect.Float64 {
			distValue.SetFloat(sourceValue.Float())
		}
	case reflect.Slice, reflect.Array:
		if sourceValue.Type().Kind() == reflect.Array || sourceValue.Type().Kind() == reflect.Slice {
			distValue.Set(sourceValue)
		}
	case reflect.Struct:
		switch distValue.Type() {
		case reflect.TypeOf(time.Time{}):
			if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
				distValue.Set(reflect.ValueOf(ParseTime(sourceValue.String())))
			}
		case reflect.TypeOf(time.Second):
			if sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "" {
				distValue.Set(reflect.ValueOf(ParseDuration(sourceValue.String())))
			}
		}
	}
}
