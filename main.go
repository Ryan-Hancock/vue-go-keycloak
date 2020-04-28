package main

import (
	"github.com/ryan-hancock/VueJWT/server"
)

func main() {
	r := server.NewRouter()
	r.Run(":8090")
}
