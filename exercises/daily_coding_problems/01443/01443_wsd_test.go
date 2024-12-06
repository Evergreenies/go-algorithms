/*
Word sense disambiguation is the problem of determining which sense a word takes on in a
particular setting, if that word has multiple meanings. For example, in the sentence
"I went to get money from the bank", bank probably means the place where people deposit
money, not the land beside a river or lake.

Suppose you are given a list of meanings for several words, formatted like so:

	{
	    "word_1": ["meaning one", "meaning two", ...],
	    ...
	    "word_n": ["meaning one", "meaning two", ...]
	}

Given a sentence, most of whose words are contained in the meaning list above, create an
algorithm that determines the likely sense of each possibly ambiguous word.
*
*/
package dailycodingproblems

import (
	"strings"
	"testing"
)

type WSD struct{}

// WordMeanings is a dictionary mapping words to their possible meanings
var WordMeanings = map[string][]string{
	"bank": {"financial institution", "side of a river"},
	"bat":  {"flying mammal", "sports equipment"},
}

func (WSD) extractContext(words []string, index int) []string {
	context := []string{}
	if index > 0 {
		context = append(context, words[index-1])
	}

	if index < len(words)-1 {
		context = append(context, words[index+1])
	}

	return context
}

func (WSD) scoreMeaning(meaning string, context []string) int {
	score := 0
	for _, word := range context {
		if strings.Contains(meaning, word) {
			score++
		}
	}

	return score
}

func (w WSD) disambiguation(sentence string) map[string]string {
	words := strings.Fields(sentence)
	results := make(map[string]string)

	for index, word := range words {
		if meanings, found := WordMeanings[word]; found {
			context := w.extractContext(words, index)

			bestScore, bestMeaning := 0, ""
			for _, meaning := range meanings {
				score := w.scoreMeaning(meaning, context)

				if score > bestScore {
					bestScore = score
					bestMeaning = meaning
				}
			}

			results[word] = bestMeaning
		}
	}

	return results
}

func TestWSD(t *testing.T) {
	wsd := WSD{}
	sentence := "I went to get money from the bank"
	disambiguated := wsd.disambiguation(sentence)
	for word, meaning := range disambiguated {
		t.Logf("Word: %s, Likely Meaning: %s\n", word, meaning)
	}
}
