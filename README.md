# Reverse Proxy

Go is a statically typed, compiled programming language design by Rob Pike, Robert Griesemer and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing and CSP-style concurrency.

Go is used to build a variety of software, but we'll use it here for:

- Networking software

Go is also know for its simplicity, efficiency and scalability. A reverse proxy has many applications (uses) we just start by implementing a load balancer.

## Features

- Load balancing

## Installation

To install the **reverse-proxy**, you will need to have Go installed. You can then install the dependencies by running the following command:

`go get .`

## Documentation

For more documentation, please see the following:

- [README file](README.md)

## Examples

The folder **examples** contains three one file servers that could be use to test the capabilities of load balancing of the reverse-proxy. They are leave with the main package to easy its use.

```go
// Collecting a list of backend servers.
endpoints := []string{
    // By not specifying an IP address before the colon, the server will listen
    // on every IP address associated with your computer
    ":8081",
    ":8082",
    ":8083",
}

// Create a Load Balancer
loadBalancer := &http.ReverseProxy{
    BackendService: http.BackendService{Backend: endpoints},
}

// Run the Load Balancer
loadBalancer.Run()
```

## Contributing

If you would like to contribute to the library, please fork the repository and submit a pull request.

## License

The library is licensed under the MIT License.

```

```
