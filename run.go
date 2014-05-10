package main


import (
  "os"
  "fmt"
)

func main() {


  urls := os.Args[1:]

  for _,url := range urls {
    err := DoFile(url)
    if (err != nil) {
      fmt.Println("error on",url)
      fmt.Println(err.Error())
    }
  }

}
