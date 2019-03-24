# Go Qlueless Assembly API

A Go implementation of a simple API to store and access Kanban related entities and events.

- This project is undertaken with the audible aid of [Avantasia](https://www.avantasia.net)
- This README was structured on a template by [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)

## Why is the project called 'Qlueless'?

I apologise, it's a poor play on words attempting to combine:

1. `Clueless`: One of the projects purposes is to learn and experiment with technologies such as Go, React, and public pipeline tools; I'm moderately clueless about the latter two.
2. `Queueless`: I want to experiment with ways of visualising and emphasising work in progress that is not, in fact, being progressed, i.e. half finished work sitting in queues waiting for someone to finish them. Once visible and being measured I can start to analyze it, and experiment with ways of reducing and avoiding it!

## Getting Started

### Prerequisites

- Go: [https://golang.org/dl/]
- Git: [https://git-scm.com]
- https://github.com/gorilla/mux `go get -u github.com/gorilla/mux`
- https://github.com/stretchr/testify `go get -u github.com/stretchr/testify`
- A decent web browser

### Installing (Linux/Bash)

Navigate to a suitable directory, open a terminal, and copy+paste the following:

```
git clone https://github.com/PaulioRandall/go-qlueless-assembly-api.git
cd go-qlueless-assembly-api/scripts
./new-dev-session
```

In order, this will:

1. Clone the source code repository
2. Navigate to the user `scripts` directory
3. Execute script to build the OpenAPI script
4. Execute script to build the application
5. Execute script to test the application
6. Start the application
7. Execute script to open a tab to the entry endpoint in your browser
8. Execute script to open a tab to an OpenAPI specification viewer in your browser
9. Log server output it's terminated

Many different scripts are available under `/scripts` but those prefixed with an underscore `_` are designed to be called by the other scripts.

### Running unit tests (Linux/Bash)

Open a terminal at the project root:

```
cd /scripts
./build-test
```

If you would like detailed test information:

```
./build
./test-verbose
```

### Running API tests (Linux/Bash)

Coming soon!

### Deployment

This has not been researched yet. Use `./build-test-run` within the `/scripts` folder to run.

## Built With

- [OpenAPI](https://swagger.io/docs/specification/about/)
- [Go](https://golang.org)
- [gorilla/mux](https://github.com/gorilla/mux)
- [testify](https://github.com/stretchr/testify)

## Contributing

I don't think this is applicable, at least not within the foreseeable future.

## Versioning

This projects API [CHANGELOG](https://github.com/PaulioRandall/go-qlueless-assembly-api/blob/master/api/CHANGELOG.md) format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and the API adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Authors

- [Me](https://github.com/PaulioRandall)

## License

This project is licensed under the [MIT License](https://github.com/PaulioRandall/go-qlueless-assembly-api/blob/master/LICENSE).

## Acknowledgments

- Influences
  - 'The Goal' by Eliyhahu M. Goldratt
  - Continuous Integration
  - Continuous Delivery