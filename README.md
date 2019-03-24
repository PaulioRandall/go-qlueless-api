# Go Qlueless Assembly API

This is a Go implementation of an API to access Kanban style lists with a manufacturing theme. 

- This project is undertaken with the audible aid of [Avantasia](https://www.avantasia.net)
- This README was structured on a template by [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)

## Why is the project called 'Qlueless'?

I apologise, it's a poor play on words attempting to combine:

1. `Clueless`: One of the projects purposes is to learn and experiment with technologies such as Go, React, and public pipeline tools; I'm moderately clueless about the latter two.
2. `Queueless`: I want to experiment with ways of visualising and emphasising work in progress that is not, in fact, being progressed, i.e. half finished work sitting in queues waiting for someone to finish them. Once visible and being measured I can start to analyze it, and experiment with ways of reducing and avoiding it!

## Getting Started

### Prerequisites

- Go: [https://golang.org/dl/]
- https://github.com/gorilla/mux `go get -u github.com/gorilla/mux`
- https://github.com/stretchr/testify `go get -u github.com/stretchr/testify`

### Installing (Linux/Bash)

Within a terminal:

```
mkdir -p "${GOPATH}/src/github.com/PaulioRandall"
cd "${GOPATH}/src/github.com/PaulioRandall"
git clone https://github.com/PaulioRandall/go-qlueless-assembly-api.git
```

From there move to the `scripts` directory for a range of activities including building, testing, running, and opening documentation within your browser:

```
cd go-qlueless-assembly-api/scripts
```

### Running unit tests (Linux/Bash)

Within a terminal:

```
cd "${GOPATH}/src/github.com/PaulioRandall/go-qlueless-assembly-api/scripts"
./test-all
```

Alternative you can run the tests with verbose `-v` flag enabled using:

```
./test-all-verbose
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

This projects [CHANGELOG](https://github.com/PaulioRandall/go-qlueless-assembly-api/blob/master/api/CHANGELOG.md) format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Authors

- [Me](https://github.com/PaulioRandall)

## License

This project is licensed under the [MIT License](https://github.com/PaulioRandall/go-qlueless-assembly-api/blob/master/LICENSE).

## Acknowledgments

- Influences
  - 'The Goal' by Eliyhahu M. Goldratt
  - Continuous Integration
  - Continuous Delivery