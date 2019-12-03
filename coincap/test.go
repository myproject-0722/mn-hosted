package main

import (
  "fmt"
//  "os"
//  "path/filepath"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://api.coincap.io/v2/assets"
  method := "GET"

  client := &http.Client {
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  //defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
