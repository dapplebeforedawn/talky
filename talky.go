package main

import (
  "dapplebeforedawn/talky/tagger"
  "encoding/json"
  "bytes"
  "os"
  "io"
  "bufio"
  "fmt"
)

type Tokens map[string][]string

func main() {
  token_map := make(Tokens)
  var b bytes.Buffer
  map_file, err := os.Open("./token_map.json")
  if err != nil {panic(err)}

  _, read_err := b.ReadFrom(map_file)
  if read_err != nil {panic(read_err)}

  json_err := json.Unmarshal(b.Bytes(), &token_map)
  if json_err != nil {panic(json_err)}

  tagger := tagger.Tagger{}
  words := Lex(os.Stdin)
  tagger.Tag(words, token_map)
  fmt.Println(*tagger.ToTagMap())
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
