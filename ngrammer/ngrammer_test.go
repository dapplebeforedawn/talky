package ngrammer

import "testing"
import "fmt"

func TestMostCommonFirst(t *testing.T) {
  commonWord   := []string{"hundreds", "of", "dollars"}
  uncommonWord := []string{"hundreds", "in", "dollars"}
  word := MostCommon( [][]string{commonWord, uncommonWord} )
  if fmt.Sprintf("%v", word) == fmt.Sprintf("%v", uncommonWord) {
    t.Error("For", "common words",
      "Expected", commonWord,
      "Got",      word,
    )
  }
}

func TestMostCommonLast(t *testing.T) {
  commonWord   := []string{"hundreds", "of", "dollars"}
  uncommonWord := []string{"hundreds", "in", "dollars"}
  word := MostCommon( [][]string{uncommonWord, commonWord} )
  if fmt.Sprintf("%v", word) == fmt.Sprintf("%v", uncommonWord) {
    t.Error("For", "common words",
      "Expected", commonWord,
      "Got",      word,
    )
  }
}

