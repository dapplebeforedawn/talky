package lexer

import (
  "bufio"
  "regexp"
  "io"
)

type Lexer struct {}

func Lex(r io.Reader) []string {
  ret := make([]string, 0)
  scan := bufio.NewScanner(r)
  scan.Split(bufio.ScanWords)
  for scan.Scan() {
    ret = append(ret, scan.Text())
  }
  return removeBlanks(removePunctuation(ret))
}

func removePunctuation(words []string) []string {
  cleaned := words
  reg, _  := regexp.Compile("[\"(){}<>,;:]|\\.[\")}]?$")
  for i, word := range words {
    cleaned[i] = reg.ReplaceAllString(word, "")
  }
  return cleaned
}

func removeBlanks(words []string) []string {
  var cleaned []string
  for _, word := range words {
    if len(word) != 0 {
      cleaned = append(cleaned, word)
    }
  }
  return cleaned
}
