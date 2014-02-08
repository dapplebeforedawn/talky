package tagger

import (
  "strings"
  "strconv"
  "regexp"
)


// after tagging, Words and Tags will be the
//same length
type Tagger struct {
  Tags  []string
  Words []string
}

type TagMap map[string][]string

func (t *Tagger) ToTagMap() *TagMap {
  tag_map := make(TagMap)
  for i, tag := range t.Tags {
    words, was_found := tag_map[tag]
    if was_found {
      tag_map[tag] = append(words, t.Words[i])
    } else {
      tag_map[tag] = []string{ t.Words[i] }
    }
  }
  return &tag_map
}

func (t *Tagger) Tag(lex []string, token_map map[string][]string) {
  t.Words = lex
  for _, word := range t.Words {
    parts := token_map[word]

    if parts == nil {
      parts = token_map[strings.ToLower(word)]
    }

    if parts == nil {
      parts = []string{"NN"}
    }
    t.Tags = append(t.Tags, parts[0]) // a flat array of string tag names
  }

  t.applyTransformations()
}

func (t *Tagger) applyTransformations() {
  for i, tag := range t.Tags {

    //  rule 1: DT, {VBD | VBP} --> DT, NN
    t.Tags[i] = t.ruleOne(i, tag)

    // rule 2: convert a noun to a number (CD) if "." appears in the word
    //         or a URL if that works
    t.Tags[i] = t.ruleTwo(i, tag)

    // rule 3: convert a noun to a past participle if words ends with "ed"
    t.Tags[i] = t.ruleThree(i, tag)

    // rule 4: convert any type to adverb if it ends in "ly";
    t.Tags[i] = t.ruleFour(i, tag)

    // rule 5: convert a common noun (NN or NNS) to a adjective if it ends with "al"
    t.Tags[i] = t.ruleFive(i, tag)

    // rule 6: convert a noun to a verb if the preceding work is "would"
    t.Tags[i] = t.ruleSix(i, tag)

    // rule 7: if a word has been categorized as a common noun and it ends with "s",
    //         then set its type to plural common noun (NNS)
    t.Tags[i] = t.ruleSeven(i, tag)

    // rule 8: convert a common noun to a present participle verb (i.e., a gerund)
    t.Tags[i] = t.ruleEight(i, tag)
  }
}


func (t *Tagger) ruleEight(i int, tag string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "NN") && strings.HasSuffix(t.Words[i], "ing") { transformed = "VBG" }
  return
}

func (t *Tagger) ruleSeven(i int, tag string) (transformed string) {
  transformed = tag
  if tag == "NN" && strings.HasSuffix(t.Words[i], "s") { transformed = "NNS" }
  return
}

func (t *Tagger) ruleSix(i int, tag string) (transformed string) {
  transformed = tag
  if i == 0 { return }
  if strings.HasPrefix(tag, "NN") && strings.ToLower(t.Words[i-1]) == "would" { transformed = "VB" }
  return
}

func (t *Tagger) ruleFive(i int, tag string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "NN") && strings.HasSuffix(t.Words[i], "al") { transformed = "JJ" }
  return
}

func (t *Tagger) ruleFour(i int, tag string) (transformed string) {
  transformed = tag
  if strings.HasSuffix(t.Words[i], "ly") { transformed = "RB" }
  return
}

func (t *Tagger) ruleThree(i int, tag string) (transformed string) {
  transformed = tag
  if strings.HasPrefix(tag, "N") && strings.HasPrefix(t.Words[i], "ed") { transformed = "VBN" }
  return
}

func (t *Tagger) ruleTwo(i int, tag string) (transformed string) {
  transformed = tag
  word := t.Words[i]
  if !strings.HasPrefix(tag, "N") { return }

  _, parse_err := strconv.ParseFloat(tag, 32)
  if parse_err == nil { transformed = "CD" }

  if strings.Contains(word, ".") {
    _, regex_err := regexp.Match("[a-zA-Z]{2}", []byte(word))
    if regex_err == nil { transformed = "URL" }
  }

  return
}

func (t *Tagger) ruleOne(i int, tag string) (transformed string) {
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
