
.PHONY: test build install clean

test:	build
	@echo
	@echo "Testing Go package"
	@echo "------------------"
	go test -v

build:	libcharsetdetect/build/libcharsetdetect.dylib
	@echo
	@echo "Building Go package"
	@echo "-------------------"
	go build -v -x

install:	build
	@echo
	@echo "Installing Go package"
	@echo "---------------------"
	go install -x

libcharsetdetect/build/libcharsetdetect.dylib: libcharsetdetect/Makefile
	@echo
	@echo "Building libcharsetdetect"
	@echo "-------------------------"
	cd libcharsetdetect && make

libcharsetdetect/Makefile:
	@echo
	@echo "Configuring libcharsetdetect"
	@echo "----------------------------"
	cd libcharsetdetect && ./configure

clean:
	@echo
	@echo "Cleaning up"
	@echo "-----------"
	go clean -x
	cd libcharsetdetect && test -f Makefile && make clean && rm Makefile CMakeCache.txt
