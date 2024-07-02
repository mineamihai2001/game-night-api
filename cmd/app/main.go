package main

import (
	"fmt"
	"github.com/mineamihai2001/game-night/helpers"
	"github.com/mineamihai2001/game-night/internal/api/router"
)

func main() {
	r := router.Create()

	env := helpers.Env()

	r.Run(fmt.Sprintf(":%d", env.App.Port))
}
