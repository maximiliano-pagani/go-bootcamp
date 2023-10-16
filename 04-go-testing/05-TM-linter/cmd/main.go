package main

import (
	"linter/cmd/server"
	"linter/prey"
	"linter/shark"
	"linter/simulator"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	srv := server.NewServer(handler, r)
	srv.MapRoutes()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
