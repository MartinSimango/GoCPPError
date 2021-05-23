package cpperror

import (
	"unsafe"
)

// #cgo LDFLAGS: -lgoerror
// #include <ErrorWrapper.h>
import "C"

type CPPError interface {
	error
	Free()
	GetFuncReturnType() int
	GetFuncReturnValue() interface{}
	GetFuncReturnPtrValue(funcReturnType int) unsafe.Pointer
}

type CPPErrorImpl struct {
	Ptr unsafe.Pointer
}

//check the CPPError is implemented
var _ CPPError = &CPPErrorImpl{}

//Error returns the error message
func (cppe CPPErrorImpl) Error() string {
	return *cppe.GetErrorMessage()
}

//Free frees the memory allocated to cppe.Ptr. This method will soon be removed once smart pointers are used within C++
// Error library
func (cppe CPPErrorImpl) Free() {
	C.DestroyError(cppe.Ptr)
}

//GetErrorMessage returns a pointer to the error message
func (cppe CPPErrorImpl) GetErrorMessage() *string {
	errorMessage := C.GetErrorMessage(cppe.Ptr)
	if errorMessage == nil {
		return nil
	}
	errorMessageString := C.GoString(errorMessage)
	return &errorMessageString
}

//GetFuncReturnType returns the type id of the cppe's delegated function
func (cppe CPPErrorImpl) GetFuncReturnType() int {
	return int(C.GetFuncReturnType(cppe.Ptr))
}

//GetFuncReturnValue returns the value of the cppe's delegated function
func (cppe CPPErrorImpl) GetFuncReturnValue() interface{} {
	switch cppe.GetFuncReturnType() {
	case INT_TYPE:
		return int(C.GetFuncReturnValue_Int(cppe.Ptr))
	case BOOL_TYPE:
		return bool(C.GetFuncReturnValue_Bool(cppe.Ptr))
	case STRING_TYPE:
		return C.GoString(C.GetFuncReturnValue_String(cppe.Ptr))
	case DOUBLE_TYPE:
		return float64(C.GetFuncReturnValue_Double(cppe.Ptr))
	}
	return nil
}

func (cppe CPPErrorImpl) GetFuncReturnStructValue(CStructTypeId int) unsafe.Pointer {
	return C.GetFuncReturnValue_Struct(cppe.Ptr, C.int(CStructTypeId))

}
