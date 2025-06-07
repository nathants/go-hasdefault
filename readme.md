# go-hasdefault

## What

A go linter to check that all switch statements have a default case.

## Why

Sometimes a missing default is an error.

## Install

`go install github.com/nathants/go-hasdefault@latest`

## Usage

```bash
>> go-hasdefault $(find test/good/ -name '*.go')

>> go-hasdefault $(find test/bad/ -name '*.go')
test/bad/bad.go:3: switch statement missing default case

```

