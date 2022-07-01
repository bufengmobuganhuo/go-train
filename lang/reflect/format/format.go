package format

import (
	"reflect"
	"strconv"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// 格式化一个值，且不分析它的内部结构
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	// 没有任何值
	case reflect.Invalid:
		return "invalue"
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
		// 忽略浮点数的处理
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		// 返回具体类型和指针地址
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}
