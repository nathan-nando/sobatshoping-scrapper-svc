package main

import (
	"fmt"
	"github.com/nathan-nando/sobatshoping-scrapper-svc/app"
)

func main() {
	cfg := app.New()
	fmt.Println(cfg.Env)
}
