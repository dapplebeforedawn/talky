package ngrammer

import "testing"

func TestMostCommonFirst(t *testing.T) {
  commonWord   := "hundreds of dollars"
  uncommonWord := "hundreds in dollars"
  word := MostCommon( []string{commonWord, uncommonWord} )
  if word == uncommonWord {
    t.Error("For", "common words",
      "Expected", commonWord,
      "Got",      word,
    )
  }
}

func TestMostCommonLast(t *testing.T) {
  commonWord   := "hundreds of dollars"
  uncommonWord := "hundreds in dollars"
  word := MostCommon( []string{uncommonWord, commonWord} )
  if word == uncommonWord {
    t.Error("For", "common words",
      "Expected", commonWord,
      "Got",      word,
    )
  }
}

