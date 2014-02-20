# Talky

Some experimental code to generate MF DOOM style hip-hop lyrics.   They parse, and flow, but you have _no idea_ what he's talking about half the time.

## TODO

I'm pretty happy with the sentences that are generated from phrasal templates.  I think some more improvements can be made by a trigram phase to refine which prepositions are chosen.  e.g. It may generate "hundreds on dollars" or hundreds in dollars" with the same phrasal template, but both of those sound weird in a sentence, an some basic n-gram analysis shows that "hundreds of dollars" is a far more common trigram.

My suspicion is that the flow of language comes much more from the preposition, conjunctions and other glue words; and far less from the nouns and verbs conveying the subject and actions.
