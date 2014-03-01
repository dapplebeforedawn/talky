package smoother

import (
  "dapplebeforedawn/talky/tagger"
  "dapplebeforedawn/talky/ngrammer"
)

type Smoother struct {
  Pos         string
  sandwiches  []sandwich
  replaceOpts []replaceOpt
  bestFits    []bestFit
}

type sandwich struct {
  position int      // index in the word array the target was found
  words    [3]string // 3 strings, the middle one is targted for replace
}

type replaceOpt struct {
  sandwich *sandwich
  options   [][3]string // all the permutations
}

type bestFit struct {
  sandwich *sandwich
  bestFit  [3]string
}

func Smooth(pos string, tags []tagger.TagPair) []string {
  smoother := &Smoother{ Pos: pos }
  smoother.buildSandwiches(tags)
  smoother.buildReplaceOptions()
  smoother.buildBestFits()
  return smoother.smooth(tags)
}

func (s *Smoother)smooth(tags []tagger.TagPair) (smoothed []string) {
  for _, fit := range s.bestFits {
    tags[fit.sandwich.position].Word = fit.bestFit[1]
  }
  for _, tag := range tags {
    smoothed = append(smoothed, tag.Word)
  }
  return
}

func (s *Smoother)buildBestFits() {
  for _, opt := range s.replaceOpts {
    bestFit := bestFit{
      sandwich: opt.sandwich,
      bestFit: ngrammer.MostCommon(opt.options),
    }
    s.bestFits = append(s.bestFits, bestFit)
  }
}

func (s *Smoother)buildSandwiches(tags []tagger.TagPair) {
  for i, pair := range tags {
    if i == 0 || i+1 == len(tags) { continue } // skip the first and last
    if pair.Tag == s.Pos {
      sandwich := sandwich{
        position: i,
        words:    [3]string{tags[i-1].Word, tags[i].Word, tags[i+1].Word},
      }
      s.sandwiches = append(s.sandwiches, sandwich)
    }
  }
}

func (s *Smoother)buildReplaceOptions() {
  for _, sandwich := range s.sandwiches {
    option := replaceOpt{
      sandwich: &sandwich,
      options: buildOpts(sandwich.words[0], sandwich.words[2]),
    }
    s.replaceOpts = append(s.replaceOpts, option)
  }
}

func buildOpts(left string, right string) (ngrams [][3]string) {
  for _, prep := range Prepositions {
    ngram := [3]string{left, prep, right}
    ngrams = append(ngrams, ngram)
  }
  return
}

