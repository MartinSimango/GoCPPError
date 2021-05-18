#pragma once

#include <type_traits>
#include <ErrorConstants.h>
#include <string.h>
#include <iostream>
#include <Delegate.hpp>
#include <ErrorBase.hpp>
#include <ErrorException.hpp>
#include <ResponseWrapper.h>

//TODO add support for const functions

//Assumes clas Object method that is being called throws Exceptions dervied from the Exception class
template<typename RetVal, typename Object, typename... Args>
class Error: public Delegate<RetVal, Args...>, public ErrorBase {

    private:

        using MemberFunctionPtrType = typename MemberFunctionPtr<RetVal, Object, Args...>::Type; 
        using FunctionPtrType = typename FunctionPtr<RetVal, Args...>::Type;
        Object * object;
        MemberFunctionPtrType memberFunctionPtr;
        FunctionPtrType* functionPtr;
    
        void setFunctionReturnType() override {
            if (std::is_same<RetVal, int>::value) {
                funcReturnType = INT_TYPE;
            }

            else if (std::is_same<RetVal, bool>::value) {
                funcReturnType = BOOL_TYPE;
            }

            else if (std::is_same<RetVal, char *>::value || std::is_same<RetVal, std::string>::value) {
                funcReturnType = STRING_TYPE;
            }

            else if (std::is_same<RetVal, double>::value) {
                funcReturnType = DOUBLE_TYPE;
            }
            else if (std::is_same<RetVal, CreateResponseStruct>::value) {
                funcReturnType = CREATE_RESPONSE_TYPE;
            }
            else {
                funcReturnType = PTR_TYPE;
            }
        }


    public:

    explicit Error (MemberFunctionPtrType delegateFunc, Object * obj) : ErrorBase(), memberFunctionPtr(delegateFunc), object(obj) {
        retVal = (RetVal*) NULL;
        setFunctionReturnType();
    }


    ~Error() {
        delete errorMessage;
        delete reinterpret_cast<RetVal*>(retVal);
        errorMessage = NULL;
    }
   

    virtual void Execute(Args... args) override
    {       
        RetVal *tempVal = reinterpret_cast<RetVal*>(retVal); //ensures that if exception is thrown retVal retains it's defaul value

         try {
            retVal = new RetVal((object->*memberFunctionPtr)(std::forward<Args>(args)...));
            delete tempVal;
                
         }
         catch(error::ErrorException *e){
            errorMessage = new char[ERROR_MAX_LENGTH];
            e->getErrorMessage(errorMessage);
            retVal = tempVal;
            delete e;
         }
         

    }

      
    RetVal getFunctionReturnValue(){ //overrides default implementation in ErrorBase
        // todo throw exception is null
        return *reinterpret_cast<RetVal*>(retVal);
    }
    

};

