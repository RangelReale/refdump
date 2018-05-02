// This library contains helpers to generate textual strings of the "reflect"
// package's values, like "reflect.Value", "reflect.Type", "reflect.Kind".
package refdump

/*
Example

    type S1 struct {
        A1 int
        A2 string
    }

    ta := &S1{
        A1: 10,
        A2: "Value",
    }

    fmt.Printf("%s\n" RefDumpValue(reflect.ValueOf(ta)))

    type XX struct {
    }

    m := make(map[string]*XX)

    fmt.Printf("%s\n", RefDumpValue(reflect.ValueOf(m)))


Output:

    Kind:(Ptr Struct) Name:(*main.S1) [!CanAddr,!CanSet]

    Kind:(Map) Key:{Kind:(String)} Elem:{Kind:(Ptr Struct) Name:(*main.XX)} Len:(0) [!CanAddr,!CanSet]

*/
