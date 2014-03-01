package constructor

import (
  "dapplebeforedawn/talky/tagger"
  "math/rand"
  "time"
  "bytes"
  "regexp"
  "strings"
)

const NOT_FOUND string = "_"

type Constructor struct {
  Tokens  []string
  TagMap  tagger.TagMap
}

func NewConstructor(tokens []string, t_map tagger.TagMap) *Constructor {
  rand.Seed( time.Now().UnixNano() )
  return &Constructor{
    Tokens: tokens,
    TagMap: t_map,
  }
}

func (c *Constructor) Construct() []string {
  sentance_words := make([]string, 0)
  for _, token := range c.Tokens {
    options, was_found := c.TagMap[token]
    if !was_found { options = []string{NOT_FOUND} }
    sentance_words = append(sentance_words, any(options))
  }
  return c.condition(sentance_words)
}

func (c *Constructor) condition(words []string) []string {
  conditioned_words := words
  conditioned_words = c.fixCapitalization(conditioned_words)
  return conditioned_words
}

func (c *Constructor) fixCapitalization(words []string) []string {
  cleaned := words
  cleaned[0] = string( bytes.Title([]byte(words[0])) )
  for i, word := range words[1:] {
    if !c.isProperNoun(word) {
      cleaned[i+1] = strings.ToLower(word)
    }
  }
  return cleaned
}

func (c *Constructor) anForA(words []string) []string {
  cleaned := words
  for i, word := range words[:len(words)-1] {
    normalized := strings.ToLower(word)
    if normalized == "a" && startsWithVowel(words[i+1]) {
      cleaned[i] = words[i] + "n"
    }
  }
  return cleaned
}

func (c *Constructor) aForAn(words []string) []string {
  cleaned := words
  for i, word := range words[:len(words)-1] {
    normalized := strings.ToLower(word)
    if normalized == "an" && !startsWithVowel(words[i+1]) {
      cleaned[i] = strings.TrimSuffix(words[i], "n")
    }
  }
  return cleaned
}

func (c *Constructor) isProperNoun(word string) bool {
  proper_nouns := append(c.TagMap["NNP"], c.TagMap["NNPS"]...)
  proper_nouns  = append(proper_nouns, "I")
  return contains(proper_nouns, word)
}

func contains(haystack []string, needle string) bool {
  for _, a := range haystack { if a == needle { return true } }
  return false
}

func any(options []string) string {
  index := rand.Intn(len(options))
  return options[index]
}

func startsWithVowel(word string) bool{
  matched, _ := regexp.Match("^[aeiou]", []byte(word))
  return matched
}

