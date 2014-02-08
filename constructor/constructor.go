package constructor

import (
  "dapplebeforedawn/talky/tagger"
  "math/rand"
  "time"
)

func Construct(tokens []string, t_map tagger.TagMap) (ret []string) {
  rand.Seed( time.Now().UnixNano() )
  for _, token := range tokens {
    options, was_found := t_map[token]
    if !was_found { options = []string{"_"} }
    ret = append(ret, any(options))
  }
  return
}

func any(options []string) string {
  index := rand.Intn(len(options))
  return options[index]
}
