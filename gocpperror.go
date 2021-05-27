package gocpperror

import (
	"unsafe"
)

// #cgo LDFLAGS: -lgoerror
// #include <cpperror/ErrorWrapper.h>
import "C"

type CPPError interface {
	error
	GetFuncReturnType() int
	GetFuncReturnValue() interface{}
	GetFuncReturnStructValue(CStructTypeId uint32) unsafe.Pointer
	Free()
}

type CPPErrorImpl struct {
	ptr unsafe.Pointer
}

//check the CPPError is implemented
var _ CPPError = &CPPErrorImpl{}

//NewCPPErrorImpl is a constructor
func NewCPPErrorImpl(ptr unsafe.Pointer) *CPPErrorImpl {
	return &CPPErrorImpl{
		ptr: ptr,
	}
}

//Error returns the error message
func (cppe CPPErrorImpl) Error() string {
	return *cppe.GetErrorMessage()
}

//GetErrorMessage returns a pointer to the error message
func (cppe CPPErrorImpl) GetErrorMessage() *string {
	errorMessage := C.GetErrorMessage(cppe.ptr)
	if errorMessage == nil {
		return nil
	}
	errorMessageString := C.GoString(errorMessage)
	return &errorMessageString
}

//GetFuncReturnType returns the type id of the cppe's delegated function
func (cppe CPPErrorImpl) GetFuncReturnType() int {
	return int(C.GetFuncReturnType(cppe.ptr))
}

//GetFuncReturnValue returns the value of the cppe's delegated function
func (cppe CPPErrorImpl) GetFuncReturnValue() interface{} {
	switch cppe.GetFuncReturnType() {
	case INT_TYPE:
		return int(C.GetFuncReturnValue_Int(cppe.ptr))
	case BOOL_TYPE:
		return bool(C.GetFuncReturnValue_Bool(cppe.ptr))
	case STRING_TYPE:
		return C.GoString(C.GetFuncReturnValue_String(cppe.ptr))
	case DOUBLE_TYPE:
		return float64(C.GetFuncReturnValue_Double(cppe.ptr))
	}
	return nil
}

//GetFuncReturnStructValue returns the value of the cppe's delgated function. The return type will be
//the type that maps to the Struct with id CStructTypeId
func (cppe CPPErrorImpl) GetFuncReturnStructValue(CStructTypeId uint32) unsafe.Pointer {
	return C.GetFuncReturnValue_Struct(cppe.ptr, CStructTypeId)
}

//Free deallocates the memory allocated to cppe.Ptr.
func (cppe CPPErrorImpl) Free() {
	C.DestroyError(cppe.ptr)
}
