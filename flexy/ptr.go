package flexy

import (
	"reflect"
	"unsafe"
)

type flexValue struct {
	// typ holds the type of the value represented by a Value.
	typ unsafe.Pointer

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer
}

type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func ToPointerValue(v interface{}) reflect.Value {
	return reflect.NewAt(reflect.TypeOf(v), (*emptyInterface)(unsafe.Pointer(&v)).word)
}

func FastPack(v reflect.Value) (i interface{}) {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	vv := (*flexValue)(unsafe.Pointer(&v))

	e.word = vv.ptr
	e.typ = vv.typ

	return i
}
