package curl

import "testing"

func TestGetArtTagsInRandomOrder(t *testing.T) {
	tags := GetArtTagsInRandomOrder()
	counter := 0
	for _, t := range TAGS {
		if contains(tags, t) {
			counter += 1
		}
	}
	if len(TAGS) != counter {
		t.Errorf("Expected slice of randomized tags to contain %v elements, but it contained %v.", len(TAGS), counter)
	}
}

func contains(collection []string, item string) bool {
	for _, ci := range collection {
		if ci == item {
			return true
		}
	}
	return false
}
