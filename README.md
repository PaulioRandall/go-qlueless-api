# Qlueless Assembly Line API

This is a Go implementation of an API to access Kanban style lists with a manufacturing theme. 

- This project is undertaken with the audible aid of [Avantasia](https://www.avantasia.net)
- This README was structured on a template by [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)

## Why is the project called 'Qlueless'?

I apologise, it's a poor play on words attempting to combine:

1. `Clueless`: One of the projects purposes is to learn and experiment with technologies such as Go, React, and public pipeline tools which I'm moderately clueless about.
2. `Queueless`: I want to experiment with ways of visualising and emphasising work in progress that is not, in fact, being progressed, i.e. half finished work sitting in queues waiting for someone to finish them. Once visible and measurable I can start to reduce it and Get Things Done!

## Getting Started

### Prerequisites

- Go: [https://golang.org/dl/]
- https://github.com/gorilla/mux `got get -u github.com/gorilla/mux`

### Installing (Linux/Bash)

Within a terminal:

```
mkdir -p "${GOPATH}/src/github.com/PaulioRandall"
cd "${GOPATH}/src/github.com/PaulioRandall"
git clone https://github.com/PaulioRandall/qlueless-assembly-line-api.git
```

From there move to the `scripts` directory for a range of activities including building, running, and opening documentation within your browser:

```
cd qlueless-assembly-line-api/scripts
```

### Running the tests

Don't have any yet :|

### Deployment

Hang on I haven't even got any tests yet.

## Built With

- [Go](https://golang.org)
- [OpenAPI](https://swagger.io/docs/specification/about/)

## Contributing

I don't think this is applicable, at least not within the foreseeable future.

## Versioning

This projects [CHANGELOG](https://github.com/PaulioRandall/qlueless-assembly-line-api/blob/master/api/CHANGELOG.md) format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Authors

- [Me](https://github.com/PaulioRandall)

## License

This project is licensed under the [MIT License](https://github.com/PaulioRandall/qlueless-assembly-line-api/blob/master/LICENSE).

## Acknowledgments

- Designs influences
  - 'The Goal' (by Eliyhahu M. Goldratt)
  - Continuous Integration
  - Continuous Delivery