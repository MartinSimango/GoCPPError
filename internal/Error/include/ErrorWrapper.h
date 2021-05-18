#pragma once
#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <ErrorConstants.h>

void DestroyError(void *err);

char * GetErrorMessage(void *err);

int GetFuncReturnType(void *err);

int GetFuncReturnValue_Int(void *err);

bool GetFuncReturnValue_Bool(void *err);

char * GetFuncReturnValue_String(void *err);

double GetFuncReturnValue_Double(void *err);

void* GetFuncReturnValue_Ptr(void * err, int type);

#ifdef __cplusplus
}  // extern "C"
#endif
