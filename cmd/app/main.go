package main

import (
	_ "github.com/lib/pq"
	"posts/internal/app"
)

func main() {
	app.RunServer()
}
