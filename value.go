package reflectx

import (
	"fmt"
	"github.com/billcoding/calls"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "[reflectx]", log.LstdFlags)

//SetValue
func SetValue(sourceValue reflect.Value, distValue reflect.Value) {
	switch distValue.Kind() {
	case reflect.Bool:
		calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
			vboolval, err := strconv.ParseBool(sourceValue.String())
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			} else {
				distValue.SetBool(vboolval)
			}
		})
		calls.True(sourceValue.Type().Kind() == reflect.Bool, func() {
			distValue.Set(sourceValue)
		})
	case reflect.String:
		calls.True(sourceValue.CanInterface(), func() {
			distValue.SetString(fmt.Sprintf("%v", sourceValue.Interface()))
		})
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
			vintval, err := strconv.ParseInt(sourceValue.String(), 10, 64)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			} else {
				distValue.SetInt(vintval)
			}
		})
		calls.True(sourceValue.Type().Kind() == reflect.Float32 ||
			sourceValue.Type().Kind() == reflect.Float64, func() {
			distValue.SetInt(int64(sourceValue.Float()))
		})
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
			vuintval, err := strconv.ParseUint(sourceValue.String(), 10, 64)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			} else {
				distValue.SetUint(vuintval)
			}
		})
		calls.True(sourceValue.Type().Kind() == reflect.Float32 ||
			sourceValue.Type().Kind() == reflect.Float64, func() {
			distValue.SetUint(uint64(sourceValue.Float()))
		})
	case reflect.Float32, reflect.Float64:
		calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
			vfloatval, err := strconv.ParseFloat(sourceValue.String(), 32)
			if err != nil {
				logger.Println(fmt.Sprintf("[SetValue]%v", err))
			} else {
				distValue.SetFloat(vfloatval)
			}
		})
		calls.True(sourceValue.Type().Kind() == reflect.Float32 ||
			sourceValue.Type().Kind() == reflect.Float64, func() {
			distValue.SetFloat(sourceValue.Float())
		})
	case reflect.Slice, reflect.Array:
		calls.True(sourceValue.Type().Kind() == reflect.Array ||
			sourceValue.Type().Kind() == reflect.Slice, func() {
			distValue.Set(sourceValue)
		})
	case reflect.Struct:
		switch distValue.Type() {
		case reflect.TypeOf(time.Time{}):
			calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
				distValue.Set(reflect.ValueOf(ParseTime(sourceValue.String())))
			})
		case reflect.TypeOf(time.Second):
			calls.True(sourceValue.Type().Kind() == reflect.String && sourceValue.String() != "", func() {
				distValue.Set(reflect.ValueOf(ParseDuration(sourceValue.String())))
			})
		}
	}
}
