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

type Printer func(string) ()

func Greet3(sal Salu, do Printer, isFormal bool) {
  mesg1, mesg2 := CreateMultiMessage3(sal.name, sal.company, "param3")
  fmt.Println(isFormal)
  if isFormal {
    do(mesg1)
  }
  do(mesg2)

}


func Greet2(sal Salu, do func(string)) {
  mesg1, mesg2 := CreateMultiMessage3(sal.name, sal.company, "param3")
  do(mesg1)
  do(mesg2)

}

func Greet(sal Salu)  {
  fmt.Println(sal.name)
  fmt.Println(sal.company)
  fmt.Println(CreateMessage(sal.name, sal.company))

  _, mesg2 := CreateMultiMessage1(sal.name, sal.company)
  //fmt.Println(mesg1)
  fmt.Println(mesg2)

  mesg1, mesg2 := CreateMultiMessage2(sal.name, sal.company)
  fmt.Println(mesg1)
  fmt.Println(mesg2)

  mesg1, mesg3 := CreateMultiMessage3(sal.name, sal.company, "param3")
  fmt.Println(mesg1)
  fmt.Println(mesg3)
}

func CreateMultiMessage1(name, company string) (string, string) {
  return name + " CreateMultiMessage1 " + company, company + " CreateMultiMessage1 " + name

}


func CreateMultiMessage2(name, company string) (mesg1 string, mesg2 string) {
  mesg1 = name + " CreateMultiMessage2 " + company
  mesg2 = company + " CreateMultiMessage2 " + name
  return

}

func CreateMultiMessage3(name string, company ...string) (mesg1 string, mesg2 string) {
  fmt.Println(len(company))
  mesg1 = name + " CreateMultiMessage3 " + company[0]
  mesg2 = company[1] + " CreateMultiMessage3 " + name
  return

}


func CreateMessage(name, company string) string {
  return name + " " + company
}

func Print(s string)  {
  fmt.Print(s)
}

func PrintLine(s string)  {
  fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {

  return  func(s string) {
    fmt.Println(s + custom)
  }

}

func main() {
  //var message string = "Hello World"
  //var message = "Hello World"
  //var mesg Salutation = "hello world"
  message := "Hello World"
  var greeting *string = & message
  *greeting = "hi"
  //var s = Salu{"param1", "param2"}
  //var s = Salu{company: "param2", name: "param1"}
  var s = Salu{}
  s.name = "param1"
  s.company = "param2"
  //var a,b,c int = 1, 2, 3
  //var a,b,c = 1, false, 3
  //a,b,c := 1, false, 3
  //message = "Hello World"
  //fmt.Println("hello world from param1")
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

  //Greet(s)
  //Greet2(s, Print)
  //Greet2(s, PrintLine)
  //Greet3(s, PrintLine)
  Greet3(s, CreatePrintFunction("!!!"), true)

}
