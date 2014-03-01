package main

import (
  "net/http"
  "strings"
  "io"
  "io/ioutil"
)

func main() {
  words := []string{
    "hundreds of dollars",
    "hundreds in dollars",
  }

  body  := strings.Join(words, "\n")
  stats := post(strings.NewReader(body))
  println(stats)
}

func post(body io.Reader) string {
  url     := "http://web-ngram.researctringh.microsoft.com/rest/lookup.svc/bing-body/jun09/3/jp?u=24d4cec1-08a0-4754-b568-c9d8a9f97754&format=json"
  content := "application/json"

  resp, _ := http.Post(url, content, body)
  defer resp.Body.Close()

  rBody, _ := ioutil.ReadAll(resp.Body)
  return string(rBody)
}
