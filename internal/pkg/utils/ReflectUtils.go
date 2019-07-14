package utils

import "reflect"

// ConvToSlice 方法可以将任意类型转换成切片
func ConvToSlice(arg interface{}) (out []interface{}, ok bool) {
	slice, success := IsKindOf(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

// IsKindOf 方法用来判断对象实例类型是否相同
func IsKindOf(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return val, ok
}
