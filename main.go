package main

import (
	"fmt"

	"github.com/gowesmart/api-gowesmart/app"
	"github.com/gowesmart/api-gowesmart/helper"
)

func main() {
	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", helper.GetEnv("SERVER_PORT", "3000")))
}
