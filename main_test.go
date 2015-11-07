package main

import "testing"

func TestWeHaveEveryBadge(t *testing.T) {
	manyBadges := howManyBadges()
	if manyBadges != quantityOfBadges {
		t.Error("We do not have all of the badges, ", manyBadges)
	}
}
