default: all

build: bin
	go build -o bin/cli main.go

cross-build:
	GOOS=linux GOARCH=amd64 go build -o bin/cli-linux main.go

bin:
	@make -p bin


sync:

