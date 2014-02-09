package main

import (
  "dapplebeforedawn/talky/tagger"
  "dapplebeforedawn/talky/lexer"
  "dapplebeforedawn/talky/constructor"
  "encoding/json"
  "bytes"
  "os"
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
  words  := lexer.Lex(os.Stdin)
  tagger.Tag(words, token_map)
  // fmt.Println(tagger.ToTagMap())

  structure := []string{"PRP", "VBD", "NN", "IN", "NNS", "DT", "NN", "VBG", "PRP$", "NN"}
  // structure := []string{"MD", "DT", "NN", "NN", "JJ", "NN"}
  sentence := constructor.NewConstructor(structure, tagger.ToTagMap())
  fmt.Println(sentence.Construct())
}
