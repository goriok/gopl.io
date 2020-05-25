package main

import(
  "fmt"
  "os"
)

func main(){
  printArgWithJoin()
  printArgWithoutJoin()
}

func printArgWithJoin(){
  fmt.Println(os.Args[1:])
}

func printArgWithoutJoin(){
  var s string
  sep := "\n"
  for _, arg := range os.Args[1:] {
    s += arg + sep
  }
  fmt.Println(s)
}

