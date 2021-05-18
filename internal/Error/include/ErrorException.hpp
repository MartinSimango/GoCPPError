#pragma once
#include <iostream>
#include <exception>


namespace error
 {

    class ErrorException : public std::exception {
        
        public:
            ErrorException() : std::exception(){}
            
            virtual char * getErrorMessage(char * error) = 0;
        
    };

}