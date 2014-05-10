package main
import (
  "net/http"
  "io/ioutil"
)


func GetBytes(url string) ([]byte,error) {

  resp,err :=  http.Get(url)

  if (err != nil) {
    return []byte{},err
  }

  defer resp.Body.Close()
  body := resp.Body

  return ioutil.ReadAll(body)
}
