.PHONY: run

all: run

run:
	cat input_board.txt | go run ./cmd/cli
