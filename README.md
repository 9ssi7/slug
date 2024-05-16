# Slug

[![Go Reference](https://pkg.go.dev/badge/github.com/9ssi7/slug.svg)](https://pkg.go.dev/github.com/9ssi7/slug) [![Go Report Card](https://goreportcard.com/badge/github.com/9ssi7/slug)](https://goreportcard.com/report/github.com/9ssi7/slug) [![Apache License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Slug is a simple, lightweight, and fast slug generator for Go.

## Installation

```bash
go get github.com/9ssi7/slug
```

## Usage

```go
packslug main

import (
    "fmt"

    "github.com/9ssi7/slug"
)

func main() {
    fmt.Println(slug.New("Hello World!"))
    // Output: hello-world

    // check if a string is a valid slug
    res := slug.Is("hello-world") // true
}
```

## Documentation

Documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/9ssi7/slug).

## Contributing

Contributions are always welcome!

## License

[MIT](LICENSE)
