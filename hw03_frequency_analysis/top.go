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

	if len(counter) < top {
		top = len(counter)
	}

	words := make([]word, top)

	for val, count := range counter {
		words = append(words, word{val, count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}

		return words[i].value < words[j].value
	})

	for i := 0; i < top; i++ {
		result = append(result, words[i].value)
	}

	return result
}
