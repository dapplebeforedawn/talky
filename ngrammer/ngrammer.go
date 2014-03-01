package ngrammer

import (
  "net/http"
  "strings"
  "io"
  "io/ioutil"
  "encoding/json"
)

const api_id string   = "24d4cec1-08a0-4754-b568-c9d8a9f97754"
const format string   = "format=json"
const endpoint string = "http://web-ngram.research.microsoft.com/rest/lookup.svc"
const corpus string   = "/bing-body/jun09/3"
const statType string = "/cp"

func MostCommon(words [][3]string) [3]string {
  grams     := flatten(words)
  body      := strings.Join(grams, "\n")
  statsJSON := getStats(strings.NewReader(body))
  stats     := []float64{}
  json.Unmarshal(statsJSON, &stats)

  maxIdx := max(stats)
  return words[maxIdx]
}

func getStats(body io.Reader) []byte {
  url     := endpoint+corpus+statType+"?u="+api_id+"&"+format
  content := "application/json"

  resp, _ := http.Post(url, content, body)
  defer resp.Body.Close()

  rBody, _ := ioutil.ReadAll(resp.Body)
  return rBody
}

func max(stats []float64) (maxIdx int) {
  for i, stat := range stats {
    if stat >= stats[maxIdx] { maxIdx = i }
  }
  return
}

func flatten(words [][3]string) (grams []string) {
  for _, gramSet := range words {
    grams = append(grams, strings.Join(gramSet[:], " "))
  }
  return
}
