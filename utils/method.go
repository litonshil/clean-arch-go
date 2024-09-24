package utils

import "reflect"

func IsEmpty(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func InArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
