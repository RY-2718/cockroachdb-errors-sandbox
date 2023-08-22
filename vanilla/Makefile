.PHONY: run curl-errors curl-trace-error curl-trace-library-error

run:
	go run cmd/main.go

curl-errors: curl-trace-error curl-trace-library-error

curl-trace-error:
	curl http://localhost:8888/trace-error

curl-trace-library-error:
	curl http://localhost:8888/trace-library-error
