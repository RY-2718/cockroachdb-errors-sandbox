.PHONY: run trace-error trace-library-error

run:
	go run cmd/main.go

trace-error:
	curl http://localhost:8888/trace-error

trace-library-error:
	curl http://localhost:8888/trace-library-error
