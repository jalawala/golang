package main

import "fmt"

type  Salutation string
type  Salu struct {
  name string
  company string
}

const (
  PI = 3.14
  Launguage = "Go"
  A = iota
  B = iota
  C = iota
)

const (
  D = iota
  E = iota
  F = iota
)
const (
  G = iota
  H
  I

)

func Greet(sal Salu)  {
  fmt.Println(sal.name)
  fmt.Println(sal.company)
  fmt.Println(CreateMessage(sal.name, sal.company))
}

func CreateMessage(name, company string) string {
  return name + " " + company
}


func main() {
  //var message string = "Hello World"
  //var message = "Hello World"
  //var mesg Salutation = "hello world"
  message := "Hello World"
  var greeting *string = & message
  *greeting = "hi"
  //var s = Salu{"JP", "AMAZON"}
  //var s = Salu{company: "Amazon", name: "JP"}
  var s = Salu{}
  s.name = "JP"
  s.company = "AWS"
  //var a,b,c int = 1, 2, 3
  //var a,b,c = 1, false, 3
  //a,b,c := 1, false, 3
  //message = "Hello World"
  //fmt.Println("hello world from JP")
  //fmt.Printf(message)
  //fmt.Printf(message)
  //fmt.Println(message,a,b,c, greeting, *greeting)
  //fmt.Println(mesg)
  //fmt.Println(s.name)
  //fmt.Println(s.company)
  //fmt.Println(A, B, C)
  //fmt.Println(D, E, F)
  //fmt.Println(G, H, I)
  //fmt.Println(PI)
  //fmt.Println(Launguage)

  Greet(s)

}
