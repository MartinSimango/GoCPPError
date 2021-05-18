#pragma once

#include <type_traits>
#include <string.h>
#include <iostream>
#include <Error.hpp>

//TODO partial template decrease code size by using inheritance maybe

template<typename Object, typename... Args>
class Error<void, Object, Args...> : public Delegate<void, Args...>, public ErrorBase {

    protected:

        using MemberFunctionPtrType = typename MemberFunctionPtr<void, Object, Args...>::Type; 
        using FunctionPtrType = typename FunctionPtr<void, Args...>::Type;
        Object * object;
        MemberFunctionPtrType memberFunctionPtr;
        FunctionPtrType* functionPtr;
    
        void setFunctionReturnType() override {
            funcReturnType = VOID_TYPE;
        }


    public:

        explicit Error (MemberFunctionPtrType delegateFunc, Object * obj) : ErrorBase(), memberFunctionPtr(delegateFunc), object(obj) {
            setFunctionReturnType(); // used for ErrorWrapper
        }


        ~Error() {
            delete errorMessage;
            errorMessage = NULL;
        }
    

        virtual void Execute(Args... args) override
        {   
            try {
                (object->*memberFunctionPtr)(std::forward<Args>(args)...);
            }  
            catch(error::ErrorException *e){
                errorMessage = new char[ERROR_MAX_LENGTH];
                e->getErrorMessage(errorMessage);
                delete e;
            }
        }

        void getFunctionReturnValue(){
            //do Nothing
        }
    
};


