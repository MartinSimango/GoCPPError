package cerror

import (
	"unsafe"
)

// #cgo LDFLAGS: -L${SRCDIR}/../../internal/Error/lib  -lgoerror
// #cgo CFLAGS: -I${SRCDIR}/../../internal/Error/include
// #include <ErrorWrapper.h>
import "C"

type CError interface {
	error
	Free()
	GetFuncReturnType() int
	GetFuncReturnValue() interface{}
	GetFuncReturnPtrValue(funcReturnType C.int) unsafe.Pointer
}

type CErrorImpl struct {
	Ptr unsafe.Pointer
}

//check the CError is implemented
var _ CError = &CErrorImpl{}

//Error returns the error message
func (ce CErrorImpl) Error() string {
	return *ce.GetErrorMessage()
}

//Free frees the memory allocated to ce.Ptr. This method will soon be removed once smart pointers are used within C++
// Error library
func (ce CErrorImpl) Free() {
	C.DestroyError(ce.Ptr)
}

//GetErrorMessage returns a pointer to the error message
func (ce CErrorImpl) GetErrorMessage() *string {
	errorMessage := C.GetErrorMessage(ce.Ptr)
	if errorMessage == nil {
		return nil
	}
	errorMessageString := C.GoString(errorMessage)
	return &errorMessageString
}

func (ce CErrorImpl) GetFuncReturnType() int {
	return int(C.GetFuncReturnType(ce.Ptr))
}

func (ce CErrorImpl) GetFuncReturnValue() interface{} {
	switch ce.GetFuncReturnType() {
	case int(C.INT_TYPE):
		return int(C.GetFuncReturnValue_Int(ce.Ptr))
	case int(C.BOOL_TYPE):
		return bool(C.GetFuncReturnValue_Bool(ce.Ptr))
	case int(C.STRING_TYPE):
		return C.GoString(C.GetFuncReturnValue_String(ce.Ptr))
	case int(C.DOUBLE_TYPE):
		return float64(C.GetFuncReturnValue_Double(ce.Ptr))
	}
	return nil
}

func (ce CErrorImpl) GetFuncReturnPtrValue(funcReturnType C.int) unsafe.Pointer {
	return C.GetFuncReturnValue_Ptr(ce.Ptr, funcReturnType)

}
