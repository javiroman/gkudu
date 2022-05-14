.DEFAULT_GOAL := demo
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
PROTO_DIRS = "pkg/proto"
GENERATE_FILES_FOLDER = pkg/proto

all : test demo doc clean generate
.PHONY: all

test:
	go test ./pkg

demo:
	go run examples/demo.go

generate:
	mkdir $(GENERATE_FILES_FOLDER) 2> /dev/null || true
	protoc --proto_path=$(PROTO_DIRS) --go_opt=module=${PACKAGE} --go_out=$(GENERATE_FILES_FOLDER) pkg/proto/*.proto

doc:
	godoc -http=:6060

clean:
	rm examples/demo 2> /dev/null || true
	rm -fr $(GENERATE_FILES_FOLDER)/*.pb.go 2> /dev/null || true
