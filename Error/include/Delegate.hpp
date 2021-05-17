#pragma once

//Code below from influenced from:
//https://simoncoenen.com/blog/programming/CPP_Delegates#:~:text=Usage-,Delegate,reference%20count%20of%20a%20shared_ptr.

template<typename RetVal, typename Object, typename ...Args>
struct MemberFunctionPtr
{
    using Type = RetVal(Object::*)(Args...);
};


template<typename RetVal, typename ...Args>
struct FunctionPtr
{
    using Type = RetVal(*)(Args...);
};


template<typename RetVal, typename... Args>
class Delegate
{
public:
    
    virtual void Execute(Args... args) = 0;

};
