package main

import (
  "dapplebeforedawn/talky/tagger"
  "dapplebeforedawn/talky/lexer"
  "dapplebeforedawn/talky/smoother"
  "dapplebeforedawn/talky/constructor"
  "dapplebeforedawn/talky/options"
  "encoding/json"
  "bytes"
  "os"
  "strings"
  "fmt"
)

type Tokens map[string][]string

func main() {
  opts := opts.Options{}
  opts.Parse()

  token_map := make(Tokens)
  var b bytes.Buffer
  map_file, err := os.Open("./token_map.json")
  if err != nil {panic(err)}

  _, read_err := b.ReadFrom(map_file)
  if read_err != nil {panic(read_err)}

  json_err := json.Unmarshal(b.Bytes(), &token_map)
  if json_err != nil {panic(json_err)}

  iTagger := tagger.Tagger{}
  words  := lexer.Lex(os.Stdin)
  iTagger.Tag(words, token_map)

  structure         := opts.Blueprint
  sentenceConstruct := constructor.NewConstructor(structure, iTagger.ToTagMap())

  sentence  := sentenceConstruct.Construct()
  sTagger   := tagger.Tagger{}
  sTagger.Tag(sentence, token_map)
  smooth    := smoother.Smooth("IN", sTagger.ToTagPairs())

  fmt.Println(strings.Join(smooth, " "))
}
