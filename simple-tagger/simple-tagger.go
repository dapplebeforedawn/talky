package main

import (
  "encoding/json"
  "bytes"
  "os"
  "io"
  "bufio"
  "fmt"
  "strings"
  "strconv"
  "regexp"
)

type Words map[string][]string

func main() {
  words := make(Words)
  var b bytes.Buffer
  map_file, err := os.Open("../token_map.json")
  if err != nil {panic(err)}

  _, read_err := b.ReadFrom(map_file)
  if read_err != nil {panic(read_err)}

  json_err := json.Unmarshal(b.Bytes(), &words)
  if json_err != nil {panic(json_err)}

  tagger := Tagger{}
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

type Tagger struct {
  Tags []string
}

func (t *Tagger) Tag(lex []string, words map[string] []string) {
  for _, word := range lex {
    parts := words[word]

    if parts == nil {
      parts = words[strings.ToLower(word)]
    }

    if parts == nil {
      parts = []string{"NN"}
    }
    t.Tags = append(t.Tags, parts[0]) // a flat array of string tag names
  }

  t.applyTransformations(lex)
}

func (t *Tagger) applyTransformations(words []string) {
  for i, tag := range t.Tags {

    //  rule 1: DT, {VBD | VBP} --> DT, NN
    t.Tags[i] = t.ruleOne(i, tag, words)

    // rule 2: convert a noun to a number (CD) if "." appears in the word
    //         or a URL if that works
    t.Tags[i] = t.ruleTwo(i, tag, words)

    // rule 3: convert a noun to a past participle if words ends with "ed"
    t.Tags[i] = t.ruleThree(i, tag, words)

    // rule 4: convert any type to adverb if it ends in "ly";
    t.Tags[i] = t.ruleFour(i, tag, words)

    // rule 5: convert a common noun (NN or NNS) to a adjective if it ends with "al"
    t.Tags[i] = t.ruleFive(i, tag, words)

    // rule 6: convert a noun to a verb if the preceding work is "would"
    t.Tags[i] = t.ruleSix(i, tag, words)

    // rule 7: if a word has been categorized as a common noun and it ends with "s",
    //         then set its type to plural common noun (NNS)
    t.Tags[i] = t.ruleSeven(i, tag, words)

    // rule 8: convert a common noun to a present participle verb (i.e., a gerund)
    t.Tags[i] = t.ruleEight(i, tag, words)
  }
}


func (t *Tagger) ruleEight(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "NN") && strings.HasSuffix(words[i], "ing") { transformed = "VBG" }
  return
}

func (t *Tagger) ruleSeven(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if tag == "NN" && strings.HasSuffix(words[i], "s") { transformed = "NNS" }
  return
}

func (t *Tagger) ruleSix(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if i == 0 { return }
  if strings.HasPrefix(tag, "NN") && strings.ToLower(words[i-1]) == "would" { transformed = "VB" }
  return
}

func (t *Tagger) ruleFive(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "NN") && strings.HasSuffix(words[i], "al") { transformed = "JJ" }
  return
}

func (t *Tagger) ruleFour(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if strings.HasSuffix(words[i], "ly") { transformed = "RB" }
  return
}

func (t *Tagger) ruleThree(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "N") && strings.HasPrefix(words[i], "ed") { transformed = "VBN" }
  return
}

func (t *Tagger) ruleTwo(i int, tag string, words []string) (transformed string) {
  transformed = tag
  word := words[i]
  if !strings.HasPrefix(tag, "N") { return }

  _, parse_err := strconv.ParseFloat(tag, 32)
  if parse_err == nil { transformed = "CD" }

  if strings.Contains(word, ".") {
    _, regex_err := regexp.Match("[a-zA-Z]{2}", []byte(word))
    if regex_err == nil { transformed = "URL" }
  }

  return
}

func (t *Tagger) ruleOne(i int, tag string, words []string) (transformed string) {
  transformed = tag
  if i == 0 { return }
  if t.Tags[i-1] == "DT" {
    if contains([]string{ "VBD", "VBP", "VB" }, tag) {
      transformed = "NN"
    }
  }
  return
}

func contains(haystack []string, needle string) bool {
  for _, a := range haystack { if a == needle { return true } }
  return false
}
