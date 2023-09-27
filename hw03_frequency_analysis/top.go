package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type word struct {
	value string
	count int
}

var r = regexp.MustCompile(`\S+`)

func Top10(text string) (result []string) {
	matches := r.FindAllString(text, -1)
	counter := map[string]int{}
	top := 10

	for _, val := range matches {
		counter[val]++
	}

	var wordsMap []word

	for val, count := range counter {
		wordsMap = append(wordsMap, word{val, count})
	}

	sort.Slice(wordsMap, func(i, j int) bool {
		if wordsMap[i].count != wordsMap[j].count {
			return wordsMap[i].count > wordsMap[j].count
		}

		return wordsMap[i].value < wordsMap[j].value
	})

	if len(wordsMap) < top {
		top = len(wordsMap)
	}

	for i := 0; i < top; i++ {
		result = append(result, wordsMap[i].value)
	}

	return result
}
