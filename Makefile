
.PHONY: build clean

build:	libcharsetdetect/build/libcharsetdetect.a
	@echo
	@echo "Building Go package"
	@echo "-------------------"
	go build -v -x

libcharsetdetect/build/libcharsetdetect.a: libcharsetdetect/Makefile
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
