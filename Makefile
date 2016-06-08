BIN=giles

$(BIN): *.go
	go build
	
run:
	./giles
