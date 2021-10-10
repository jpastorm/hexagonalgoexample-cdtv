package api

import (
	"github.com/jpastorm/hexagonalgoexample-cdtv/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}