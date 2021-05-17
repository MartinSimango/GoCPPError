package cerror

import (
	"unsafe"
)

// #cgo CFLAGS:  -I${SRCDIR}/../../FileClient/include -I../../FilePacket/include -I../../Error/include
// #cgo LDFLAGS: -L${SRCDIR}/../../Error/lib -lgoerror
// #include "ErrorWrapper.h"
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

func (ce CErrorImpl) Error() string {
	return *ce.GetErrorMessage()
}

func (ce CErrorImpl) Free() {
	C.DestroyError(ce.Ptr)
}

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
