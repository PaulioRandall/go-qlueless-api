# Go Qlueless API

A Go implementation of a simple API to store and access Kanban related entities and events.

- This project is undertaken with the audible aid of [Avantasia](https://www.avantasia.net) and [Dream Theater](http://dreamtheater.net)
- This README was structured on a template by [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)

## Why is the project called 'Qlueless'?

I apologise, it's a poor play on words attempting to combine:

1. `Clueless`: One of the projects purposes is to learn and experiment with technologies such as Go, React, and public pipeline tools; I'm moderately clueless about the latter two.
2. `Queueless`: I want to experiment with ways of visualising and emphasising work in progress that is not, in fact, being progressed, i.e. half finished work sitting in queues waiting for someone to finish them. Once visible and being measured I can start to analyze it, and experiment with ways of reducing and avoiding it!

## Getting Started

### Prerequisites

- Go: [https://golang.org/dl/]
- Git: [https://git-scm.com]
- An internet connection
- A decent web browser

### Running

Navigate to a suitable directory, open a terminal, and copy+paste the following:

```
git clone https://github.com/PaulioRandall/go-qlueless-api.git
cd go-qlueless-api/scripts
./build-test-api-run.go
```

In order, this will:

1. Clone the source code repository
2. Navigate to the user `scripts` directory
3. Execute a build of the OpenAPI specification
4. Execute a build of the application
5. Execute unit tests within the application
6. Execute black box API tests on the application
7. Starts the application

### Running unit tests

Open a terminal at the project root:

```
cd /scripts
./build-test.go
```

### Running API tests

Open a terminal at the project root:

```
cd /scripts
./build-test-api.go
```

### Deployment 

> Coming soon! See **Running** in the meantime.

## Built With

- [OpenAPI](https://swagger.io/docs/specification/about/)
- [Go](https://golang.org)
- [testify](https://github.com/stretchr/testify)
- [mapstructure](https://github.com/mitchellh/mapstructure)

## Contributing

> Not applicable.

## Versioning

This projects API [CHANGELOG](https://github.com/PaulioRandall/go-qlueless-api/blob/master/api/CHANGELOG.md) format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and the API adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Authors

- [Me](https://github.com/PaulioRandall)

## License

This project is licensed under the [MIT License](https://github.com/PaulioRandall/go-qlueless-api/blob/master/LICENSE).

## Acknowledgments

- Influences
  - 'The Goal' by Eliyhahu M. Goldratt
  - Continuous Integration
  - Continuous Delivery