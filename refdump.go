package refdump

import (
	"fmt"
	"reflect"
	"strings"
)

func RefDumpType(typ reflect.Type) string {
	_, ret := RefDumpTypeGet(typ)
	return ret
}

func RefDumpTypeGet(typ reflect.Type) (reflect.Type, string) {
	kind := ""
	ptr := ""
	vt := typ
	for vt.Kind() == reflect.Ptr {
		kind += "Ptr "
		ptr += "*"
		vt = vt.Elem()
	}
	kind += RefDumpKind(vt.Kind())
	ret := fmt.Sprintf("Kind:(%s)", kind)

	if vt.PkgPath() != "" {
		ret += fmt.Sprintf(" Name:(%s%s.%s)", ptr, vt.PkgPath(), vt.Name())
	}

	// map
	if vt.Kind() == reflect.Map {
		ret += fmt.Sprintf(" Key:{%s}", RefDumpType(vt.Key()))
	}

	// array / map
	if vt.Kind() == reflect.Array || vt.Kind() == reflect.Slice || vt.Kind() == reflect.Map {
		ret += fmt.Sprintf(" Elem:{%s}", RefDumpType(vt.Elem()))
	}

	return vt, ret
}

func RefDumpValue(value reflect.Value) string {
	_, ret := RefDumpTypeGet(value.Type())

	v := value
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// len
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		ret += fmt.Sprintf(" Len:(%d)", v.Len())
	}

	// value
	if sv, isvalue := RefDumpValueString(value); isvalue {
		ret += fmt.Sprintf(" Value:(%s)", sv)
	}

	// flags
	flags := make([]string, 0)
	if !value.CanAddr() {
		flags = append(flags, "!CanAddr")
	}
	if !value.CanInterface() {
		flags = append(flags, "!CanInterface")
	}
	if !value.CanSet() {
		flags = append(flags, "!CanSet")
	}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if value.IsNil() {
			flags = append(flags, "IsNil")
		}
	}
	if !value.IsValid() {
		flags = append(flags, "!IsValid")
	}
	if len(flags) > 0 {
		ret += fmt.Sprintf(" [%s]", strings.Join(flags, ","))
	}

	return ret
}

func RefDumpKind(kind reflect.Kind) string {
	switch kind {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Bool:
		return "Bool"
	case reflect.Int:
		return "Int"
	case reflect.Int8:
		return "Int8"
	case reflect.Int16:
		return "Int16"
	case reflect.Int32:
		return "Int32"
	case reflect.Int64:
		return "Int64"
	case reflect.Uint:
		return "Uint"
	case reflect.Uint8:
		return "Uint8"
	case reflect.Uint16:
		return "Uint16"
	case reflect.Uint32:
		return "Uint32"
	case reflect.Uint64:
		return "Uint64"
	case reflect.Uintptr:
		return "Uintptr"
	case reflect.Float32:
		return "Float32"
	case reflect.Float64:
		return "Float64"
	case reflect.Complex64:
		return "Complex64"
	case reflect.Complex128:
		return "Complex128"
	case reflect.Array:
		return "Array"
	case reflect.Chan:
		return "Chan"
	case reflect.Func:
		return "Func"
	case reflect.Interface:
		return "Interface"
	case reflect.Map:
		return "Map"
	case reflect.Ptr:
		return "Ptr"
	case reflect.Slice:
		return "Slice"
	case reflect.String:
		return "String"
	case reflect.Struct:
		return "Struct"
	case reflect.UnsafePointer:
		return "UnsafePointer"
	default:
		return fmt.Sprintf("Unknown[%d]", kind)
	}
}

func RefDumpValueString(value reflect.Value) (result string, isvalue bool) {
	switch value.Kind() {
	case reflect.Invalid:
		return "<INVALID>", false
	case reflect.Bool:
		if value.Bool() {
			return "TRUE", true
		} else {
			return "FALSE", true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", value.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", value.Uint()), true
	case reflect.Uintptr:
		return "<UINTPTR>", false
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", value.Float()), true
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%f", value.Complex()), true
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
		return "", false
	case reflect.String:
		return fmt.Sprintf("%q", value.String()), true
	}
	return "", false
}
