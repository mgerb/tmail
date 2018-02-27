linux:
	go build -o ./build/tmail-linux ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./build/tmail-mac ./main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./build/tmail-windows.exe ./main.go

clean:
	rm -rf ./build

all: linux mac windows
