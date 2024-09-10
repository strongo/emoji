package emojis

import (
	"testing"
)

func TestInit(t *testing.T) {
	if len(All) == 0 {
		t.Fatal("All is empty")
	}

	validateList := func(category string, emojis []string) {
		for _, emoji := range emojis {
			if info, ok := All[emoji]; !ok {
				t.Errorf("%s: emoji %v not found in All", category, emoji)
			} else if emoji == "" {
				t.Errorf("%s: emoji %s has no description", category, emoji)
			} else {
				if info.Description == "" {
					t.Errorf("%s: emoji %s has no description", category, emoji)
				}
				if info.Category == "" {
					t.Errorf("%s: emoji %s has no category", category, emoji)
				}
			}
		}
	}
	if len(ByCategory) == 0 {
		t.Fatal("ByCategory is empty")
	}

	for category, emojis := range ByCategory {
		validateList(category, emojis)
	}
}
