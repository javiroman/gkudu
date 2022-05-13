.DEFAULT_GOAL := test

test:
	go test ./pkg

demo:
	go run examples/demo.go

doc:
	godoc -http=:6060

clean:
	rm examples/demo 2> /dev/null || true