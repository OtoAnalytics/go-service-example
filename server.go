package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	// Martini automatically serves static files from "public" directory. As an example of adding an assets directory too:
	m.Use(martini.Static("assets"))

	// Example of a root-level Get endpoint.
	m.Get("/", func() string {
		return "Hello world!"
	})

	// Example of a REST endpoint for a "locations" resource with GET for a specific entry.
	m.Get("/locations/:id", func(params martini.Params) string {
		return "Hello location " + params["id"]
	})

	// With the routes defined, start the server.
	m.Run()
}
