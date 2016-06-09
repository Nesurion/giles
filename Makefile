BIN=giles

$(BIN): *.go
	go build
	
run: $(BIN)
	./giles
