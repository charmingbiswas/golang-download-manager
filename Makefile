SUBDIR := ./cmd/gdm
BIN := ./bin
EXECUTABLE := gdm

.PHONY: clean makedir

## this command is just for testing, will be removed in prod
run:
	go run $(SUBDIR)/main.go

build: | makedir
	@go build -o $(BIN)/$(EXECUTABLE) $(SUBDIR)

makedir:
	@mkdir -p $(BIN)

clean:
	@rm -rf $(BIN)