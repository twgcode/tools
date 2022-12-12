/**
@Author: wei-G
@Email: 17600113577@163.com
@Date: 2022/12/12 09:54
@Description: 获取基础数据类型指针的值
*/

package simplestruct

func PtrBoolDefaultValue(a *bool, defaultValue bool) bool {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrBooValue(a *bool) bool {
	return PtrBoolDefaultValue(a, false)
}

func PtrIntDefaultValue(a *int, defaultValue int) int {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrIntValue(a *int) int {
	return PtrIntDefaultValue(a, 0)
}

func PtrInt8DefaultValue(a *int8, defaultValue int8) int8 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrInt8Value(a *int8) int8 {
	return PtrInt8DefaultValue(a, 0)
}

func PtrInt16DefaultValue(a *int16, defaultValue int16) int16 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrInt16Value(a *int16) int16 {
	return PtrInt16DefaultValue(a, 0)
}

func PtrInt32DefaultValue(a *int32, defaultValue int32) int32 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrInt32Value(a *int32) int32 {
	return PtrInt32DefaultValue(a, 0)
}

func PtrInt64DefaultValue(a *int64, defaultValue int64) int64 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrInt64Value(a *int64) int64 {
	return PtrInt64DefaultValue(a, 0)
}

func PtrUintDefaultValue(a *uint, defaultValue uint) uint {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrUintValue(a *uint) uint {
	return PtrUintDefaultValue(a, 0)
}

func PtrUint8DefaultValue(a *uint8, defaultValue uint8) uint8 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrUint8Value(a *uint8) uint8 {
	return PtrUint8DefaultValue(a, 0)
}

func PtrUint16DefaultValue(a *uint16, defaultValue uint16) uint16 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrUint16Value(a *uint16) uint16 {
	return PtrUint16DefaultValue(a, 0)
}

func PtrUint32DefaultValue(a *uint32, defaultValue uint32) uint32 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrUint32Value(a *uint32) uint32 {
	return PtrUint32DefaultValue(a, 0)
}

func PtrUint64DefaultValue(a *uint64, defaultValue uint64) uint64 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrUint64Value(a *uint64) uint64 {
	return PtrUint64DefaultValue(a, 0)
}

func PtrFloat32DefaultValue(a *float32, defaultValue float32) float32 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrFloat32Value(a *float32) float32 {
	return PtrFloat32DefaultValue(a, 0)
}

func PtrFloat64DefaultValue(a *float64, defaultValue float64) float64 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrFloat64Value(a *float64) float64 {
	return PtrFloat64DefaultValue(a, 0)
}

func PtrComplex64DefaultValue(a *complex64, defaultValue complex64) complex64 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrComplex64Value(a *complex64) complex64 {
	return PtrComplex64DefaultValue(a, 0)
}

func PtrComplex128DefaultValue(a *complex128, defaultValue complex128) complex128 {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrComplex128Value(a *complex128) complex128 {
	return PtrComplex128DefaultValue(a, 0)
}

func PtrStringDefaultValue(a *string, defaultValue string) string {
	if a == nil {
		return defaultValue
	}
	return *a
}
func PtrStringValue(a *string) string {
	return PtrStringDefaultValue(a, "")
}

/*
Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
*/
