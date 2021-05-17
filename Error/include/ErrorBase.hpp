#pragma once
#include <stdio.h>

class ErrorBase {

    protected:
    char *errorMessage;
    int funcReturnType;
    void * retVal;

    virtual void setFunctionReturnType() = 0;

    public:

    ErrorBase(){
        errorMessage = NULL;
        funcReturnType = -1;
    }

    virtual ~ErrorBase(){ }
    
    virtual char * getErrorMessage() const {
        return errorMessage;
    }
   
    template<typename RetVal> RetVal* getFunctionReturnValue(){
        return reinterpret_cast<RetVal *>(retVal);
    }

    virtual int getFunctionReturnType() {
        return funcReturnType;
    }

};
