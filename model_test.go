package main

import (
	"testing"
	"time"
)

func TestReview(t *testing.T) {
	review := Review{1, 100, "Tim", "Lawrence", 7, "This is a great beer", time.Now()}

	if review.Beer_id != 55 {
		t.Error("Expected 55 got", review.Beer_id)
	}

}
