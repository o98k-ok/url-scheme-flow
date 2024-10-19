all: 
	rm -rf output
	mkdir -p output
	make build
	cp -r icons output/
	cp config/* output/
	cp scripts/* output/

build:
	GOOS=darwin GOARCH=arm64 go build -o output/one-punch_darwin_arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o output/one-punch_darwin_amd64 main.go
	makefat output/one-punch output/one-punch_*
	rm -rf output/one-punch_*

clean:
	rm -rf output
