# appenv

[![GoDoc](https://godoc.org/github.com/hlts2/appenv?status.svg)](http://godoc.org/github.com/hlts2/appenv)

appenv is simple golang application runtime envrionment library.

## Requirement

Go 1.15

## Installing

```
go get github.com/hlts2/appenv
```

## Possible environment

- test
- development
- staging
- production
- unknown


## Example

```go

package main

import (
	"fmt"
	"os"

	"github.com/hlts2/appenv"
)

func init() {
	os.Setenv("APP_ENV", "production")
}

func main() {
	env := appenv.Env("APP_ENV")

	switch env {
	case appenv.Test:
		fmt.Println(env.String()) // test
	case appenv.Dev:
		fmt.Println(env.String()) // development
	case appenv.Stg:
		fmt.Println(env.String()) // staging
	case appenv.Prod:
		fmt.Println(env.String()) // production
	default:
		fmt.Println(env.String()) // unknown
	}
}
```
