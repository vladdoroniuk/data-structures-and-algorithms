# Data Structures and Algorithms in TypeScript

The repository contains my own implementations of various data structures and algorithms (DSA) in [TypeScript](https://github.com/microsoft/TypeScript).

## Overview

Each DSA resides in its own folder with the following structure:

```
dsa-name
├─ dsa-name.ts
├─ dsa-name.test.ts
└─ README.md

kebab case looks pretty decent 🤔
```

Some algorithms may have several implementations in a single `.ts` file, usually these are optimized versions. You can also read a brief description of a DSA in a `README` and run test cases specified in a `.test.ts` file, see more in [Testing](#testing).

## List of implemented DSA

The list is occasionally updated.

### Data Structures

_Currently none_

### Algorithms

- Sorting
  - [Bubble sort](./src/bubble-sort/)
- Searching
  - [Binary search](./src/binary-search/)

## Testing

The repository uses [yarn](https://github.com/yarnpkg/yarn) for package management and [Vitest](https://github.com/vitest-dev/vitest) for testing, so setup is quite fast.

1. `yarn install`
2. `yarn test`

## References

Perhaps I should add links to books, articles and scientific papers here 🤷
