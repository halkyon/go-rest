# go-rest

## What is this?

A simple HTTP server written in Go using a RESTful style API to interact with it.

It doesn't do much, just demonstrates the layout of an Go HTTP server, and might be a good starting point.

## Usage

Start the server:

    docker-compose up -d

Next up, use `curl` to check we get something back:

    curl http://localhost:8010

This should be an empty JSON response.

Create some data:

    curl http://localhost:8010/result -d '{"some":"json"}'

This should return a URI to the data resource just created.

View that data, replacing `{id}` in the following with the ID from the previous response:

    curl http://localhost:8010/data/{id}

## What now?

There's a lot more to be added to make this a real application! Here's a few examples:

* Access data from a database, or some other pluggable source (e.g. for testing)
* Check for data, output HTTP status not found if missing
* Authentication
* Use some kind of standard specification, e.g. OpenAPI or JSON API
* Validate POST data matches expectations, e.g. correct fields provided
* Tests
