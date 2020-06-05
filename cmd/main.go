package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  // a is the byte array for the first file listed
  fileone := os.Args[1]
  a, err := ioutil.ReadFile(fileone)
  if err != nil {
    fmt.Print(err)
  }

  // b is the byte array for the second file listed
  filetwo := os.Args[2]
  b, err := ioutil.ReadFile(filetwo)
  if err != nil {
    fmt.Print(err)
  }

  for i, v := range a {  // get index and value at that index in the first array
    if v != b[i] {
      fmt.Println(i, v, string(v), b[i], string(b[i]))
    }
  }
}
