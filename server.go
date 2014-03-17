package main

import "runtime"
import "strings"
import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    s := []string{"Hello world! Go version:", runtime.Version()}
    return strings.Join(s, " ")
  })
  m.Run()
}
