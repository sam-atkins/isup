# isup

![Tests](https://github.com/sam-atkins/isup/workflows/Tests/badge.svg)

This is a reworking of [is-up-cli](https://github.com/sindresorhus/is-up-cli) written in Golang. It uses [isitup.org](https://isitup.org/) under the hood.

- [isup](#isup)
  - [Install](#install)
  - [Usage](#usage)
  - [Tests](#tests)

## Install

Prerequisites:

- [Golang](https://golang.org/dl/)

To install:

- git clone
- from the newly git cloned project directory, run `go build`
- the binary should now be installed in `~/go/bin/`

## Usage

```bash
$ isup github.com
✅  github.com is up
```

## Tests

```bash
go test

# with coverage
go test -cover
```
