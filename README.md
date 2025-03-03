# Data Structures and Algorithms in Go

The repository contains my own implementations of various data structures and algorithms (DSA) in [Go](https://go.dev).

## Overview

Each DSA resides in its own folder with the following structure:

```
dsa_name
├─ dsa_name.go
├─ dsa_name_test.go
└─ README.md
```

Some algorithms may have several implementations in a single `.go` file, usually these are optimized versions. You can also read a brief description of a DSA in a `README` and run test cases specified in a `_test.go` file, see more in [Testing](#testing).

## List of implemented DSA

The list is occasionally updated.

### Data Structures

_Currently none_

### Algorithms

- Sorting
  - [Bubble sort](./bubble_sort/)
- Searching
  - [Binary search](./binary_search/)

## Testing

Just run `go test -v ./...`
