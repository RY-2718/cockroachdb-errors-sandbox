# cockroarchdb-errors-sandbox

This is a sandbox repository for testing [github.com/cockroachdb/errors](https://github.com/cockroachdb/errors).  
The `github.com/cockroachdb/errors` package can be used as a replacement for [github.com/pkg/errors](https://github.com/pkg/errors) and Go's standard `errors` package.

This repository provides an example to demonstrate how to use `github.com/cockroachdb/errors` and how errors are displayed using this library.

For more information about github.com/cockroachdb/errors, please refer to the library's official page ([GitHub](https://github.com/cockroachdb/errors) or [pkg.go.dev](https://pkg.go.dev/github.com/cockroachdb/errors)).

## Requirements

* Go 1.13 or later
* Make

## Project Structure

This project has the following samples.

* **vanilla**: A sample that does not use github.com/cockroachdb/errors. If you want to see how errors are printed without github.com/cockroachdb/errors, run this sample.
* **simple**: A sample that demonstrates github.com/cockroachdb/errors in a straightforward manner. In most scenarios, this approach should be sufficient to add context to the error.
* **withstack**: A sample that demonstrates how to add a stacktrace to an error without including additional messages.
* **redundant**: A sample that demonstrates the redundant use of github.com/cockroachdb/errors. This approach is not recommended as it leads to the repeated printing of the stacktrace.
* **message**: A sample that demonstrates how to add messages to errors using github.com/cockroachdb/errors appropriately. It shows how to enhance error information without adding a repeated stacktrace.
* **join**: A sample that demonstrates `errors.Join` from github.com/cockroachdb/errors, which is compatible with Go's standard `errors` package.

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
This sample demonstrates basic error handling with github.com/cockroachdb/errors.  

For application-generated errors, the only difference from `vanilla` example in the implementation is that the error is generated using `errors.New()`, provided by github.com/cockroachdb/errors provides, rather than Go's standard `errors` package.  
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

For library-generated errors, the error is wrapped using `errors.Wrap()`, provided by github.com/cockroachdb/errors, at the point where library-generated error occurs.  
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

### withstack

In the `withstack` sample, you can expect to see the following output.  
This sample demonstrates how to add a stacktrace to an error without including additional messages.

For library-generated errors, the error is usually wrapped using `errors.Wrap(err, "message")` at the point where library-generated error occurs.  
However, adding messages to errors using `errors.Wrap()` may not be necessary in some cases.  
In such instances, `errors.WithStack(err)` serves an appropriate alternative for adding a stacktrace without adding messages to the error.  
Using `errors.WithStack(err)` results in fewer lines in the stacktrace compared to `errors.Wrap(err, "message")`, as shown below.

```bash
2023/08/23 11:20:06 error: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/model.callInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/model/errorgenerator.go:29
  | github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/model/errorgenerator.go:21
  | github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/handler.TraceLibraryErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/withstack/pkg/handler/traceerror.go:23
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
Wraps: (2) Get "http://invalid-url"
Wraps: (3) dial tcp
Wraps: (4) lookup invalid-url: no such host
Error types: (1) *withstack.withStack (2) *url.Error (3) *net.OpError (4) *net.DNSError
```

### redundant

In the `redundant` sample, you can expect to see the following output.  
This sample demonstrates the redundant use of github.com/cockroachdb/errors.

In both implementations for the 2 type of errors, the error is wrapped using `errors.Wrap()` each time the error is passed to a different function.  
Multiple calls of `errors.Wrap()` leads to the repeated printing of the stacktrace, because each call of `errors.Wrap()` adds a new stacktrace to the error.  
As shown, the repeated stacktrace is omitted when the error is printed, but unnecessary lines are still included.  
Therefore, calling `errors.Wrap()` multiple times is generally not recommended.

```bash
2023/08/22 15:37:56 error: this is an error from ExternalFunc: this is an error from internalFunc
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.ExternalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:11
  | [...repeated from below...]
Wraps: (2) this is an error from ExternalFunc
Wraps: (3) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.internalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:17
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.ExternalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:10
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/handler.TraceErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/handler/traceerror.go:11
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
Wraps: (4) this is an error from internalFunc
Error types: (1) *withstack.withStack (2) *errutil.withPrefix (3) *withstack.withStack (4) *errutil.leafError

2023/08/22 15:37:56 error: this is an error from WrapCallInvalidHTTPRequest: failed to call invalid HTTP request: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:22
  | [...repeated from below...]
Wraps: (2) this is an error from WrapCallInvalidHTTPRequest
Wraps: (3) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.callInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:29
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/model/errorgenerator.go:21
  | github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/handler.TraceLibraryErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/redundant/pkg/handler/traceerror.go:23
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
Wraps: (4) failed to call invalid HTTP request
Wraps: (5) Get "http://invalid-url"
Wraps: (6) dial tcp
Wraps: (7) lookup invalid-url: no such host
Error types: (1) *withstack.withStack (2) *errutil.withPrefix (3) *withstack.withStack (4) *errutil.withPrefix (5) *url.Error (6) *net.OpError (7) *net.DNSError
```

### message

In the `message` sample, you can expect to see the following output.  
This sample demonstrates how to add messages to errors using github.com/cockroachdb/errors appropriately.

In this sample, `errors.WithMessage()` is used instead of `errors.Wrap()` to add a message to errors.  
`errors.WithMessage()` is useful to add messages to errors without including a stacktrace, in contrast to `errors.Wrap()`.  
When using `errors.WithMessage()`, additional messages are displayed above the stacktrace, as shown below.

Please note that you should use `errors.Wrap()` instead of `errors.WithMessage()` when handling library-generated errors.  
In other words, to display a stacktrace, you must call either `errors.New()` or `errors.Wrap()` exactly once, and then use `errors.WithMessage()` for any subsequent message additions.

```bash
2023/08/22 16:09:39 error: this is an error from ExternalFunc: this is an error from internalFunc
(1) this is an error from ExternalFunc
Wraps: (2) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model.internalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model/errorgenerator.go:17
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model.ExternalFunc
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model/errorgenerator.go:10
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/handler.TraceErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/handler/traceerror.go:11
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
Wraps: (3) this is an error from internalFunc
Error types: (1) *errutil.withPrefix (2) *withstack.withStack (3) *errutil.leafError

2023/08/22 16:09:39 error: this is an error from WrapCallInvalidHTTPRequest: failed to call invalid HTTP request: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
(1) this is an error from WrapCallInvalidHTTPRequest
Wraps: (2) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model.callInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model/errorgenerator.go:29
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/model/errorgenerator.go:21
  | github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/handler.TraceLibraryErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/message/pkg/handler/traceerror.go:23
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
Wraps: (3) failed to call invalid HTTP request
Wraps: (4) Get "http://invalid-url"
Wraps: (5) dial tcp
Wraps: (6) lookup invalid-url: no such host
Error types: (1) *errutil.withPrefix (2) *withstack.withStack (3) *errutil.withPrefix (4) *url.Error (5) *net.OpError (6) *net.DNSError
```

### join

In the `join` sample, you can expect to see the following output.  
This sample demonstrates how to use `errors.Join` from github.com/cockroachdb/errors and how the errors are presented using this library.

In this sample, `errors.Join()` is employed to consolidate multiple errors into a single error at the point where library-generated error occurs.  
The resulting error can be evaluated using `errors.Is(err, target)`, enabling us to determine the type or types of the encompassed errors.

When leveraging `errors.Join()`, the stacktrace diverges at the point of error aggregation.  
The divergence is displayed intuitively, resembling the output of the `tree` command.

```bash
2023/09/22 18:49:35 errors.Is(err, model.ExternalError) = true
2023/09/22 18:49:35 error: Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
(1) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model.callInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model/errorgenerator.go:31
  | github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model.WrapCallInvalidHTTPRequest
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model/errorgenerator.go:23
  | github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/handler.TraceLibraryErrorHandler
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/handler/traceerror.go:26
  | net/http.HandlerFunc.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2136
  | net/http.(*ServeMux).ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2514
  | net/http.serverHandler.ServeHTTP
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/net/http/server.go:2938
  | [...repeated from below...]
Wraps: (2) Get "http://invalid-url": dial tcp: lookup invalid-url: no such host
  | this is external error
└─ Wraps: (3) attached stack trace
  -- stack trace:
  | github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model.init
  |     /Users/yoshitani.ryo.10456/ghq/github.com/RY-2718/cockroachdb-errors-sandbox/join/pkg/model/errorgenerator.go:9
  | runtime.doInit1
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/proc.go:6740
  | runtime.doInit
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/proc.go:6707
  | runtime.main
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/proc.go:249
  | runtime.goexit
  |     /opt/homebrew/Cellar/go/1.21.0/libexec/src/runtime/asm_arm64.s:1197
  └─ Wraps: (4) this is external error
└─ Wraps: (5) Get "http://invalid-url"
  └─ Wraps: (6) dial tcp
    └─ Wraps: (7) lookup invalid-url: no such host
Error types: (1) *withstack.withStack (2) *join.joinError (3) *withstack.withStack (4) *errutil.leafError (5) *url.Error (6) *net.OpError (7) *net.DNSError
```
