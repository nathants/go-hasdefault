# go-hasdefault

## what

a go linter to check that all switch statements have a default case.

## why

sometimes a missing default is an error.

## install

`go install github.com/nathants/go-hasdefault@latest`

## usage

```bash
>> go-hasdefault $(find test/good/ -name '*.go')

>> go-hasdefault $(find test/bad/ -name '*.go')
test/bad/bad.go:3: switch statement missing default case

```
