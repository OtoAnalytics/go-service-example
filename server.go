package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  // Automatically serves static files from "public" directory. Add assets too:
  m.Use(martini.Static("assets"))

  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Get("/locations", func() string {
    return
    })

  m.Run()
}

