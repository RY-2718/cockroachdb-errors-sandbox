# cockroarchdb-errors-sandbox

This is a sandbox repository for testing [github.com/cockroachdb/errors](https://github.com/cockroachdb/errors).  
The `cockroachdb/errors` package can be used as a replacement for [github.com/pkg/errors](https://github.com/pkg/errors) and Go's standard `errors` package.  
This repository provides an example to demonstrate how to use `cockroachdb/errors` and how errors are displayed using this library.  
For more information about cockroachdb/errors, please refer to the library's official page.

## requirements

* Go 1.13 or later
* Make

## usage

If you want to see how errors are printed without cockroachdb/errors, run the following command.

```bash
# run HTTP server on localhost:8888
$ make run

# in another terminal:
$ make curl-trace-error
$ make curl-trace-library-error
```

You can see the following output.

```bash
$ make run
go run cmd/main.go
Server is running on http://localhost:8888
2023/08/21 17:47:10 error: this is an error from internalFunc
2023/08/21 17:47:14 error: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
```