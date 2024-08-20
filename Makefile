SUBDIR := ./cmd/gdm
BIN := ./bin
EXECUTABLE := gdm

.PHONY: clean makedir

build: | makedir
	@go build -o $(BIN)/$(EXECUTABLE) $(SUBDIR)

makedir:
	@mkdir -p $(BIN)

clean:
	@rm -rf $(BIN)