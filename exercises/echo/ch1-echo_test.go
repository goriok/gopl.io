package main

import(
  "testing"
  "os"
)
func getStrings() []string{
  slice := make([]string, 100000)
  for i := 0; i < 10; i++ {
    slice = append(slice, "")
  }
  return slice
}

func BenchmarkWithJoin(b *testing.B){
  os.Stdout,_ = os.Open(os.DevNull)
  os.Args = getStrings()

  for i := 0; i < b.N; i++ {
    printArgWithJoin()
  }
}

func BenchmarkWithoutJoin(b *testing.B){
  os.Stdout,_ = os.Open(os.DevNull)
  os.Args = getStrings()

  for i := 0; i < b.N; i++ {
    printArgWithoutJoin()
  }
}
