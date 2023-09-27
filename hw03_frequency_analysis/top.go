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

	var words []word

	for val, count := range counter {
		words = append(words, word{val, count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}

		return words[i].value < words[j].value
	})

	if len(words) < top {
		top = len(words)
	}

	for i := 0; i < top; i++ {
		result = append(result, words[i].value)
	}

	return result
}
