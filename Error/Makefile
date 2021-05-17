CC=gcc
CPP=g++
IDIR=include ../Error/include ../Response/include
ODIR=obj
LDIR=lib
SDIR=src
CFLAGS=$(patsubst %, -I%, $(IDIR)) -std=c++17
LIBNAME=libgoerror


ERROR_SOURCES=$(wildcard $(SDIR)/*.cpp)
ERROR_OBJECTS=$(patsubst $(SDIR)/%.cpp, $(ODIR)/%.o, $(ERROR_SOURCES))

ERROR_DLIB_SOURCES = $(ERROR_SOURCES)

ERROR_DEPS = $(patsubst %, %/*, $(IDIR)) 

build: build-dir $(ERROR_OBJECTS) build-lib install

build-dir:
	mkdir -p $(ODIR)
	mkdir -p $(LDIR)

$(ERROR_OBJECTS): $(ODIR)/%.o : $(SDIR)/%.cpp $(ERROR_DEPS)
	$(CPP) -c $< -o $@ $(CFLAGS)


build-lib: $(ERROR_DLIB_SOURCES) $(ERROR_DEPS)
	ar -crs $(LDIR)/$(LIBNAME).a $(ERROR_OBJECTS)
	$(CPP) -shared -fPIC -o $(LDIR)/$(LIBNAME).so $(ERROR_DLIB_SOURCES) $(CFLAGS)
	
install:
	rm -f /usr/local/lib/$(LIBNAME).so 
	rm -f /usr/local/include/ErrorWrapper.h
	rm -f /usr/local/include/ErrorConstants.h
	cp $(LDIR)/$(LIBNAME).so /usr/local/lib
	cp include/ErrorWrapper.h /usr/local/lib
	cp include/ErrorConstants.h /usr/local/lib

clean:
	rm -rf $(ODIR)
	rm -rf $(LDIR)