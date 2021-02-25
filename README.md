# Semver

[![Build Status](https://github.com/kinbiko/semver/workflows/Go/badge.svg)](https://github.com/kinbiko/semver/actions)
[![Coverage Status](https://coveralls.io/repos/github/kinbiko/semver/badge.svg?branch=master)](https://coveralls.io/github/kinbiko/semver?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kinbiko/semver)](https://goreportcard.com/report/github.com/kinbiko/semver)
[![Coverage Status](https://coveralls.io/repos/github/kinbiko/semver/badge.svg)](https://coveralls.io/github/kinbiko/semver)
[![Latest version](https://img.shields.io/github/tag/kinbiko/semver.svg?label=latest%20version&style=flat)](https://github.com/kinbiko/semver/releases)
[![Go Documentation](http://img.shields.io/badge/godoc-documentation-blue.svg?style=flat)](http://pkg.go.dev/github.com/kinbiko/semver)
[![License](https://img.shields.io/github/license/kinbiko/semver.svg?style=flat)](https://github.com/kinbiko/semver/blob/master/LICENSE)

## `Upversion` Usage

This package exposes a function `Upversion` for identifying the next semantic version
based on a current version, given an increment (major, minor, patch).

Supports the `v` prefix; if your version begins with a `v`, then the resulting
version will also begin with a `v`.

```go
version, err := semver.Upversion("minor", "v3.2.6")
if err != nil {
    return err
}
fmt.Println(version) // v3.3.0
```

### Binary

Assuming you have `$GOPATH/bin` on your path, you can install the binary via:

```console
$ go get -u github.com/kinbiko/semver/cmd/upversion
```

Then you can get the next version for the current repository with:

```console
$ upversion minor # Alternatively: `upversion major` or `upversion patch`
v0.2.0
```

## Docs

You can find the [GoDocs for this package here](https://pkg.go.dev/github.com/kinbiko/semver).

## Contributing

Contributions are welcome. Please discuss feature requests in an issue before opening a PR.
