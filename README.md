# cockroarchdb-errors-sandbox

This is a sandbox repository for testing [github.com/cockroachdb/errors](https://github.com/cockroachdb/errors).  
The `cockroachdb/errors` package can be used as a replacement for [github.com/pkg/errors](https://github.com/pkg/errors) and Go's standard `errors` package.

This repository provides an example to demonstrate how to use `cockroachdb/errors` and how errors are displayed using this library.

For more information about cockroachdb/errors, please refer to the library's official page.

## Requirements

* Go 1.13 or later
* Make

## Project Structure

This project has the following samples.

* **vanilla**: A sample that does not use cockroachdb/errors. If you want to see how errors are printed without cockroachdb/errors, run this sample.
* **simple**: A sample that demonstrates cockroachdb/errors in a straightforward manner. In most scenarios, this approach should be sufficient to add context to the error.

## What is Implemented

Each sample implements a simple HTTP server running on `localhost:8080` with 2 endpoints.  

* /trace-error: This endpoint handles an application-generated error, such as `errors.New("this is an error")`.
* /trace-library-error: This endpoint handles library-generated error, like `_, err := http.Get("http://invalid-url")`.

When these endpoints are called, the server displays the error message associated with the particular endpoint.  
Showcasing these error messages is the main purpose of this repository.

## How to Run

You can run the samples in the following way.

1. **Navigate to the Sample Directory**: Change to the directory of the sample you want to run. For example, for the vanilla sample:

```bash
$ cd vanilla
```

2. **Start the HTTP Server**: Run the HTTP server on localhost:8888.

```bash
$ make run
```

3. **Send a Request to the Server**: Open a new terminal window, navigate to the same sample directory, and send a request to the server.

```bash
# Open a new terminal and navigate to the sample directory
$ cd vanilla

# Send requests to both endpoints
$ make curl-errors

# Send a request to only 'trace-error' endpoint
$ make curl-trace-error

# Send a request to only 'trace-library-error' endpoint
$ make curl-trace-library-error
```

4. You can see the output in the terminal where the server is running.

```bash
2023/08/21 17:47:10 error: this is an error from internalFunc

2023/08/21 17:47:14 error: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
```

## Outputs

### vanilla

In the `vanilla` sample, you can expect to see the following output.  
The server only displays the error message without providing the stack trace.  
This makes it difficult to understand the cause of the error solely from the error message, as it lacks context or additional information.

```bash
2023/08/21 17:47:10 error: this is an error from internalFunc

2023/08/21 17:47:14 error: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
```

### simple

In the `simple` sample, you can expect to see the following output.  
This sample demonstrates basic error handling with cockroachdb/errors.  

For application-generated errors, the only difference from `vanilla` example in the implementation is that the error is generated using `errors.New()`, provided by cockroachdb/errors provides, rather than Go's standard `errors` package.  
With this slight change, the error message is displayed alongside the stack trace, allowing us to understand the context of the cause of the error.

```bash
2023/08/22 11:48:48 error: this is an error from internalFunc
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model.internalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model/errorgenerator.go:17
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model.ExternalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model/errorgenerator.go:10
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/handler.TraceErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/handler/traceerror.go:11
  | net/http.HandlerFunc.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2136
  | net/http.(*ServeMux).ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2514
  | net/http.serverHandler.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2938
  | net/http.(*conn).serve
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2009
  | runtime.goexit
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/asm_arm64.s:1197
Wraps: (2) this is an error from internalFunc
Error types: (1) *withstack.withStack (2) *errutil.leafError
```

For library-generated errors, the error is wrapped using `errors.Wrap()`, provided by cockroachdb/errors, at the point where library-generated error occurs.  
With this change, additional information is added to the error likewise the application-generated error.

```bash
2023/08/22 11:48:48 error: failed to call invalid HTTP request: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model.callInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model/errorgenerator.go:29
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/model/errorgenerator.go:21
  | github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/handler.TraceLibraryErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/simple/pkg/handler/traceerror.go:23
  | net/http.HandlerFunc.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2136
  | net/http.(*ServeMux).ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2514
  | net/http.serverHandler.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2938
  | net/http.(*conn).serve
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2009
  | runtime.goexit
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/asm_arm64.s:1197
Wraps: (2) failed to call invalid HTTP request
Wraps: (3) Get "http://invalid-url"
Wraps: (4) dial tcp
Wraps: (5) lookup invalid-url: no such host
Error types: (1) *withstack.withStack (2) *errutil.withPrefix (3) *url.Error (4) *net.OpError (5) *net.DNSError
```