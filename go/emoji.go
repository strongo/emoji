package emoji

import (
		"github.com/pkg/errors"
)

type Emoji struct {
	Code    string
	Codes   []string
	Unicode string
}

type Slice = []Emoji

type Map = map[string]Emoji

var All Map

func init() {
	All = make(Map, len(Food)+len(Flags))
	if err := Add(Flags...); err != nil {
		panic(err)
	}
	if err := Add(Food...); err != nil {
		panic(err)
	}
}

var ErrDuplicateCode = errors.New("duplicate code")

func Add(emoji ...Emoji) error {
	errDuplicate := func(code string) error{
		return errors.WithMessage(ErrDuplicateCode, code)
	}
	for _, e := range emoji {
		if e.Code != "" {
			if _, ok := All[e.Code]; ok {
				return errDuplicate(e.Code)
			}
			All[e.Code] = e
		}
		for _, code := range e.Codes {
			if _, ok := All[e.Code]; ok {
				return errDuplicate(code)
			}
			All[code] = e
		}
	}
	return nil
}
