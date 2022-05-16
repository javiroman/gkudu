.DEFAULT_GOAL := demo
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
PROTO_INCLUDE_FOLDERS = $(shell find pkg/proto -type d | tr '\n' ':')
GENERATE_FILES_FOLDER = pkg/proto
FILES = $(shell find pkg/proto -name *.proto)

all : test demo doc clean generate
.PHONY: all

test:
	go test ./pkg

demo:
	go run examples/demo.go

generate:
	mkdir $(GENERATE_FILES_FOLDER) 2> /dev/null || true

	protoc --proto_path=$(PROTO_INCLUDE_FOLDERS) \
           --go_opt=paths=source_relative \
           --go_out=$(GENERATE_FILES_FOLDER)  \
           $(FILES)

doc:
	godoc -http=:6060

clean:
	rm examples/demo 2> /dev/null || true
	rm -f $(shell find pkg/proto -name *.pb.go) 2> /dev/null || true
