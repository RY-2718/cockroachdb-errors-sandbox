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