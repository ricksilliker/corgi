package main

import (
  "corgi/corgi"
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func basic() string {
  return "Hello World!"
}

func main() {

  corgi.Start()

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "corgi",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(basic)
  app.Run()
}
