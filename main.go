package main

import (
	"github.com/go-macaron/pongo2"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Static("public", macaron.StaticOptions{
		SkipLogging: true,
	}))
	m.Use(pongo2.Pongoer())
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "jeremy"
		ctx.HTML(200, "login")
	})
	m.Run()
}
