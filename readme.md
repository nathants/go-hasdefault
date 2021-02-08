# go-hasdefault

## what

check that all switch statements have a default case.

## why

sometimes a missing default is an error.

## install

`go get github.com/nathants/go-hasdefault`

## usage

```bash
>> cat bad.go
package bad

func bad() {
        switch "" {
        }
}

>> go-hasdefault bad.go
bad.go:3 switch "" {

>> cat good.go
package good

func good() {
        switch "" {
        default:
        }
}

0>> go-hasdefault good.go

```
