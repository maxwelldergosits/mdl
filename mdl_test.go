package main

import (
  "fmt"
  "testing"
  "net/http"
)



/*
**************************************************
This section tests the GetBytes function (kind of important
in a file downloader eh?)
*************************************************
*/
var test_bytes []byte = []byte{0xde,0xad,0xbe,0xef}

func init() {
  testServer()
}

func testServer() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write(test_bytes)
  })

  go http.ListenAndServe(":8080", nil)
}

func TestGetBytesWithGoodURL(t * testing.T) {

  testURL := "http://127.0.0.1:8080/"

  recieved_bytes,err := GetBytes(testURL)

  if (err != nil) {
    fmt.Println("There was an unexpected error")
    t.Fail()
    return
  }

  for i,v := range recieved_bytes {
    if (v != test_bytes[i]) {
      fmt.Println("Didn't receive the correct bytes")
      t.Fail()
      return
    }
  }
}

func TestGetBytesWithBadURL(t * testing.T) {

  failURL := "not a url"

  recieved_bytes,err := GetBytes(failURL)

  if (err == nil) {
    fmt.Printf("There was no error")
    t.Fail()
    return
  }

  if(len(recieved_bytes) != 0) {
    fmt.Printf("I Received some bytes when I wasn't supposed to")
    t.Fail()
    return
  }
}

/*
*****************************************************************

Tests some directory/file functions

*****************************************************************
*/
func TestGetExt(t * testing.T) {
  url := "/path/to/this/hello.txt"
  if (GetExt(url) != ".txt") {
    fmt.Println("Wrong extension")
    t.Fail()
  }
}

func TestGetFilename(t * testing.T) {
  url := "/path/to/this/hello.txt"
  fname,_ := GetFileName(url)
  if (fname != "hello.txt") {
    fmt.Println("Wrong filename")
    t.Fail()
  }
}

