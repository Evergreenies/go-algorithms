package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* The stable marriage problem is defined as follows:
*
* Suppose you have N men and N women, and each person has ranked their prospective
* opposite-sex partners in order of preference.
*
* For example, if N = 3, the input could be something like this:
*
* guy_preferences = {
* 	'andrew': ['caroline', 'abigail', 'betty'],
* 	'bill': ['caroline', 'betty', 'abigail'],
* 	'chester': ['betty', 'caroline', 'abigail'],
* }
*
* gal_preferences = {
* 	'abigail': ['andrew', 'bill', 'chester'],
* 	'betty': ['bill', 'andrew', 'chester'],
* 	'caroline': ['bill', 'chester', 'andrew'],
* }
*
* Write an algorithm that pairs the men and women together in such a way that no
* two people of opposite sex would both rather be with each other than with their
* current partners.
* **/

func removeDuplicateMen(men_preferences map[string][]string) []string {
	men := make([]string, 0)

	for man := range men_preferences {
		found := false
		for _, iMan := range men {
			if man == iMan {
				found = true
			}
		}

		if !found {
			men = append(men, man)
		}
	}

	return men
}

func isWomanExists(woman string, matcches map[string]string) bool {
	for key := range matcches {
		if woman == key {
			return true
		}
	}

	return false
}

func removeProcessedMen(men []string, man string) []string {
	var filteredMen []string
	for _, mn := range men {
		if man != mn {
			filteredMen = append(filteredMen, mn)
		}
	}

	return filteredMen
}

func indexOfPartner(partners_preferences []string, partner string) int {
	for index, prtnr := range partners_preferences {
		if prtnr == partner {
			return index
		}
	}

	return -1
}

func stablePartner(men_preferences, women_preferences map[string][]string) map[string]string {
	men := removeDuplicateMen(men_preferences)
	var matches map[string]string
	matches = make(map[string]string)

	for len(men) > 0 {
		for _, man := range men {
			for _, woman := range men_preferences[man] {
				if !isWomanExists(woman, matches) {
					matches[man] = woman
					men = removeProcessedMen(men, man)
					break
				} else {
					currentPartner := matches[woman]
					if indexOfPartner(women_preferences[woman], man) < indexOfPartner(women_preferences[woman], currentPartner) {
						matches[man] = woman
						men = removeProcessedMen(men, man)
						men = append(men, currentPartner)
						break
					}
				}
			}
		}
	}

	return matches
}

func TestStablePartner(t *testing.T) {
	assert := assert.New(t)

	guy_preferences := map[string][]string{
		"andrew":  {"caroline", "abigail", "betty"},
		"bill":    {"caroline", "betty", "abigail"},
		"chester": {"betty", "caroline", "abigail"},
	}
	gal_preferences := map[string][]string{
		"abigail":  {"andrew", "bill", "chester"},
		"betty":    {"bill", "andrew", "chester"},
		"caroline": {"bill", "chester", "andrew"},
	}

	finalPartner := stablePartner(guy_preferences, gal_preferences)
	fmt.Printf("FINAL PARTENER: %#v\n", finalPartner)

	assert.Equal(len(finalPartner), 3, "there must be three partner at last")
}
