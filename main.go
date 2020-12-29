package main

import (
  "corgi/corgi"
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

type Config struct {
  rt *wails.Runtime
}

func (c *Config) WailsInit(runtime *wails.Runtime) error {
  c.rt = runtime
  corgi.Setup(runtime)
  return nil
}

func main() {
  config := &Config{}
  //corgi.Start()

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
  app.Bind(config)
  app.Run()
}
