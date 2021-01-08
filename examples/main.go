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
	case appenv.Local:
		fmt.Println(env.String()) // local
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
