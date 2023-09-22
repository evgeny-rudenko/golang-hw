package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string) []string {
	regExpr := regexp.MustCompile(`[-a-zA-Zа-яА-Я,]+`)
	if text == "" {
		return make([]string, 0)
	}
	text = strings.ToLower(text)
	wordCount := make(map[string]int)
	words := regExpr.FindAllString(text, -1)
	for _, word := range words {
		if word == "-" {
			continue
		}
		wordCount[word]++
	}
	type wordFrequency struct {
		word      string
		frequency int
	}
	freqList := make([]wordFrequency, 0, len(wordCount))
	for word, count := range wordCount {
		freqList = append(freqList, wordFrequency{word, count})
	}

	sort.Slice(freqList, func(i, j int) bool {
		if freqList[i].frequency == freqList[j].frequency {
			return freqList[i].word < freqList[j].word
		}
		return freqList[i].frequency > freqList[j].frequency
	})
	top10 := make([]string, 0, 10)
	for i := 0; i < len(freqList) && i < 10; i++ {
		top10 = append(top10, freqList[i].word)
	}
	return top10
}
