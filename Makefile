

compile: clean
	go build -o build/libminequery.so -buildmode=c-shared

clean:
	rm -rf build && mkdir "build"