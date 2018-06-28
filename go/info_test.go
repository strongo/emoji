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
				t.Errorf(category + ":emoji %v not found in All")
			} else if emoji == "" {
				t.Errorf(category + ":emoji %v has no description")
			} else {
				if info.Description == "" {
					t.Errorf(category + ":emoji %v has no description")
				}
				if info.Category == "" {
					t.Errorf(category + ":emoji %v has no category")
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
