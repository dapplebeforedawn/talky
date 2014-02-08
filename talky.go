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

  tagger := tagger.Tagger{}
  tokens := Lex(os.Stdin)
  tagger.Tag(tokens, words)
  fmt.Println(tagger.Tags)
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
