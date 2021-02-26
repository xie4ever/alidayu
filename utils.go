package alidayu

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type structTypeValue struct {
	Type  reflect.StructField
	Value reflect.Value
}

type structTagValue struct {
	Tag   string
	Value reflect.Value
}

func structToStringMap(src interface{}, tags ...string) (result map[string]string, err error) {
	var tag string
	if len(tags) == 0 {
		tag = "json"
	} else {
		tag = tags[0]
	}
	srcV, err := valueOf(src)
	if err != nil {
		return
	}
	srcTVs := getNameAndValues(srcV)
	tvs := make([]structTagValue, len(srcTVs))
	for _, sTV := range srcTVs {
		t := sTV.Type.Tag.Get(tag)
		if len(t) == 0 || strings.EqualFold(t, "-") {
			continue
		}
		tvs = append(tvs, structTagValue{Tag: t, Value: sTV.Value})
	}
	result = make(map[string]string, len(tvs))
	for _, tv := range tvs {
		switch tv.Value.Kind() {
		case reflect.Bool:
			result[tv.Tag] = strconv.FormatBool(tv.Value.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[tv.Tag] = strconv.FormatInt(tv.Value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result[tv.Tag] = strconv.FormatUint(tv.Value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			result[tv.Tag] = strconv.FormatFloat(tv.Value.Float(), 'f', 20, 64)
		case reflect.String:
			result[tv.Tag] = tv.Value.String()
		case reflect.Struct:
			switch item := tv.Value.Interface().(type) {
			case time.Time:
				result[tv.Tag] = item.Format(time.RFC3339)
			}
		}
	}
	return
}

func mustPtr(s interface{}) error {
	if s == nil {
		return errors.New("gsf: cannot copy nil pointer")
	}
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("gsf: %v must be reflect.Ptr", v.Type().Name())
	}
	if v.IsNil() {
		return errors.New("gsf: cannot copy nil pointer")
	}
	return nil
}

func valueOf(s interface{}) (reflect.Value, error) {
	err := mustPtr(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(s).Elem(), nil
}

func getNameAndValues(v reflect.Value) []structTypeValue {
	tvs := make([]structTypeValue, v.NumMethod())
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Anonymous {
			vf := v.Field(i)
			if vf.CanSet() && t.Field(i).Type.Kind() == reflect.Struct {
				tvs = append(tvs, getNameAndValues(vf)...)
			}
		} else {
			tvs = append(tvs, structTypeValue{Type: t.Field(i), Value: v.Field(i)})
		}
	}
	return tvs
}
