package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })


  m.Get("/woot", func() string {
  		panic("you sucker")
  		return "haha"
  	})

  m.Run()




}


