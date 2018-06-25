package emoji

import (
	"testing"
	"github.com/pkg/errors"
)

func TestInit(t *testing.T) {
	if len(All) == 0 {
		t.Fatal("All is empty")
	}
}

func TestAdd(t *testing.T) {
	if err := Add(Emoji{Code: "code1"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := Add(Emoji{Code: "code1"}); errors.Cause(err) != ErrDuplicateCode {
		t.Fatalf("Expected ErrDuplicateCode, got: %v", err)
	}
}
