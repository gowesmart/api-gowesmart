package main

import (
	"fmt"

	"github.com/gowesmart/api-gowesmart/app"
	"github.com/gowesmart/api-gowesmart/utils"
)

func main() {
	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", utils.GetEnv("SERVER_PORT", "3000")))
}
