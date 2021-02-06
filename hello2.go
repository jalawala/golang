package main

import (
  "fmt"
  greeting1 "golang/greet"
  "time"
)

func RenameToFrog(r greeting1.Renamable) {
  r.Rename("Frog")
}


func main() {

  var s = greeting1.Salu{}

  var sl []int
  sl = make ([]int, 3)
  sl[0] = 1
  sl[1] = 2
  sl[2] = 3
  //sl[4] = 4

  sl1 := []int { 1, 2, 3}
  sl1[1] = 1

  slice := greeting1.Salutations{
    { "k1", "v1"},
    {"k2", "v2"},
    {"k3", "v3"},
  }
  s.Name = "k1"
  //s.Name = "1234567890"
  s.Company = "v1"

  //slice[0].Rename("K11")
  //RenameToFrog(&slice[0])
  //fmt.Fprintf(&slice[0], "the count is %d", 10)


  //slice = slice[1:2]
  //slice = append (slice, greeting1.Salu{"k4", "v4"})

  //slice = append (slice, slice...)
  //slice = append(slice[:1], slice[2:]...)
  //greeting1.Greet3(s, greeting1.CreatePrintFunction("!!!"), true, 5)
  //greeting1.TypeSwitchFunc("dd")
  //greeting1.Greet4(slice, greeting1.CreatePrintFunction("!!!"), true, 5)

  done := make (chan bool, 2)

  go func() {
    slice.Greet5(greeting1.CreatePrintFunction("!!!"), true, 5)
    done <- true
    time.Sleep(100 * time.Millisecond)
    done <- true
    fmt.Println("done")
  }()

  slice.Greet5(greeting1.CreatePrintFunction("???"), true, 5)
  //time.Sleep(100 * time.Microsecond)
  <- done
  //for {
  //  time.Sleep(100 * time.Millisecond)
  //}

  c := make (chan greeting1.Salu)
  c2 := make (chan greeting1.Salu)

  //go slice.ChannelGreeter(c)
  //for s:= range c {
  //  fmt.Println(s.Name)
  //}
  go slice.ChannelGreeter(c)
  go slice.ChannelGreeter(c2)

  for {

    select {
      case s, ok := <- c:
        if ok {
          fmt.Println(s.Name, ":1")
        } else {
          return
        }
    case s, ok := <- c2:
      if ok {
        fmt.Println(s.Name, ":2")
      } else {
        return
      }
    default:
      fmt.Println("Waiting")
    }


  }


}
