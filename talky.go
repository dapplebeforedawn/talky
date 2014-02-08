package main

import (
  "encoding/json"
  "bytes"
  "os"
  "io"
  "bufio"
  "fmt"
)

type Words map[string][]string

func main() {
  words := make(Words)
  var b bytes.Buffer
  map_file, err := os.Open("./token_map.json")
  if err != nil {panic(err)}

  _, read_err := b.ReadFrom(map_file)
  if read_err != nil {panic(read_err)}

  json_err := json.Unmarshal(b.Bytes(), &words)
  if json_err != nil {panic(json_err)}

  // for k, v := range f {
  //   fmt.Println(k)
  //   fmt.Println(v[0])
  // }

  tokens := Lex(os.Stdin)
  tagged := Tag(tokens, words)
  fmt.Println(tagged)
}

func Lex(r io.Reader) []string {
  ret := make([]string, 0)
  scan := bufio.NewScanner(r)
  scan.Split(bufio.ScanWords)
  for scan.Scan() {
    ret = append(ret, scan.Text())
  }
  return ret
}

func Tag(lex []string, words map[string] []string) [][]string {
  ret := make([][]string, 0)
  for _, word := range lex {
    parts := words[word]
    if parts == nil { parts = make([]string, 0) }
    ret = append(ret, parts)
  }
  return ret
}
