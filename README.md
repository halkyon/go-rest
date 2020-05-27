# go-rest-server

## What is this?

A simple HTTP server written in Go using a RESTful style API to interact with it.

It doesn't do much, just demonstrates the layout of an Go HTTP server, and might be a good starting point.

## Features

* Uses [Gorilla web toolkit](https://www.gorillatoolkit.org) for routing, and [middleware](https://pkg.go.dev/github.com/gorilla/handlers) for logging requests.
* `POST` body is limited using `http.MaxBytesReader`, so the server isn't overloaded from malicious requests.
* Configurable options for address, port, timeouts, max body size, and [performance debugging](https://golang.org/pkg/net/http/pprof/).

## Usage

### Starting the server

Using [`docker-compose`](https://docs.docker.com/compose/):

    docker-compose up -d

Alternatively, you can run `make build` which places binaries for Linux, macOS, and Windows in the `build` directory.

Run the resulting binaries directly with `-h` to see command line options!

### Example API usage

View data by `GET`ing the index route `/`:

    curl http://localhost:8000

Create some data by `POST`ing to the `resource` route:

    curl http://localhost:8000/resource -d '{"Name":"Bob"}' -H "Content-Type: application/json"

View data by `GET`ing the `resource/{id}` route:

    curl http://localhost:8000/resource/1

## What now?

There's a lot more to be added to make this a real application! Here's a few examples:

* Access data from a database, or some other pluggable source (e.g. for testing)
* Check for data, output HTTP status not found if missing
* Add support for DELETE requests
* Authentication
* Use some kind of standard specification, e.g. [OpenAPI](https://swagger.io/specification/) or [JSON API](https://jsonapi.org/)
* Validate POST data matches expectations, e.g. correct fields provided
* Tests
