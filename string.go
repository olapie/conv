package conv

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// ToString converts i to string
// i can be string, integer types, bool, []byte or any types which implement fmt.Stringer
func ToString(i any) (string, error) {
	i = IndirectToStringerOrError(i)
	if i == nil {
		return "", errNilValue
	}
	switch v := i.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	case error:
		return v.Error(), nil
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool, reflect.Float32, reflect.Float64:
		return fmt.Sprint(v.Interface()), nil
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return string(v.Bytes()), nil
		}
	}
	return "", fmt.Errorf("cannot convert %#v of type %T to string", i, i)
}

func MustString(i any) string {
	v, err := ToString(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToStringSlice(i any) ([]string, error) {
	i = Indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]string); ok {
		return l, nil
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice, reflect.Array:
		num := v.Len()
		res := make([]string, num)
		var err error
		for j := 0; j < num; j++ {
			res[j], err = ToString(v.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert element at index %d: %w", i, err)
			}
		}
		return res, nil
	default:
		if s, err := ToString(i); err == nil {
			return strings.Fields(s), nil
		}
		return nil, fmt.Errorf("cannot convert %#v of type %T to []string", i, i)
	}
}

func MustStringSlice(i any) []string {
	v, err := ToStringSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func JoinAsString(a any, sep string) (string, error) {
	l, err := ToStringSlice(a)
	if err != nil {
		return "", err
	}
	return strings.Join(l, sep), nil
}

func TrimSpace[T ~string](s T) T {
	return T(strings.TrimSpace(string(s)))
}
