# gocpperror (This ReadMe is still not finished)

gocpprror is a go pacakge for easily handling C++ exceptions thrown in C++ classes that are dynamically linked to your Go application. <br />

Note: I plan on extending this to also handle C++ exceptions thrown outside of C++ classes and in regular non class functions too.


## Requirements

To use this go package you will need my cpperror library.

To install clone the repository at https://github.com/MartinSimango/cpperror

1. git clone https://github.com/MartinSimango/cpperror  && cd cpperror
3. sudo make build install

The install command will copy the built library and the libary headers files into /usr/local/lib and /usr/local/include respectively.

## Example


Lets we say we had a c++ class Foo that we wanted to wanted to use in our go program. 

Foo.hpp

```cpp

//Foo.hpp

class Foo {

  int div(int a, int b) {
    if (b < 0) {
        throw new FooException("Cannot divide by 0");
    }
    return a/b;
  }
}
```

If we wanted to use this class in Go using cgo we would need to create a wrapper interface file in C in order to wrap our c++ code.  <br /> 
We'd also need to wrapper a c++ class that implements our code in the C wrapper file.

For example the wrapper interface file would be called Foo.h and the implementation of the functions defined in this file will be in a class called FooWrapper.cpp.

Examples of the two classes can be seen below.  <br />

FooWrapper.h  <br />
```c
//FooWrapper.h


#pragma once

#ifdef __cplusplus
extern "C" {
#endif

void * NewFoo();

void * div(void * foo,int a, int b);

#ifdef __cplusplus
}  // extern "C"
#endif

```
FooWrapper.cpp  <br />
```cpp
//FooWrapper.cpp
#include "Foo.hpp"
#include "FooWrapper.h"
#include <ErrorVoid.hpp> //imported from the https://github.com/MartinSimango/cpperror repo


Foo * asFoo(void * foo) {
   return reinterpret_cast<Foo*>(foo); 
}

void * NewFoo() {
  Foo * foo = new Foo();
  return foo;
}

void * div(void *foo, int a, int b) {
  Error<int, Foo, int, int> * error = new Error<int, Foo,int, int>(&Foo::div, AsFoo(foo));
  error->Execute(a, b);
  return dynamic_cast<ErrorBase*>(error);
}

```

In our Go program we could our C function like this

```go
// #cgo LDFLAGS: -lfoo
// #include <FooWrapper.h>

import "C"

package main

import (
	"fmt"
	"os"

	cerror "github.com/MartinSimango/gocpperror"
)

type Foo struct {
  ptr unsafe.Pointer
}


func main() {
  foo := Foo{}
  foo.ptr = C.NewFoo()
  cerr := cpperror.CPPErrorImpl{}
  cerr.Ptr = C.div(foo.ptr, C.int(a), C.int(b))  // (1) a and b be will integers
  errorMessage := cerr.GetErrorMessage()
  if errorMessage != nil { // (2)
    fmt.Println(cerr) // (3) 
    cerr.Free()
    os.Exit(1)
  }
  
  div_ans := cerr.GetFuncReturnValue().(int) // (4)
  
  cerr.Free() // (5)

  
}

```
(1) Our C++ function that throws an exception will be called.  <br />
(2) We check to see if our C++ function threw an exception. <br />
(3) If the C++ function threw an exception we can just print the error and exit the application. <br />
(4) If the C++ function never threw an exception we can just get the result of the function call. <br />
(5) Free the memory that was allocated within the C library. <br />


