package emojis

type Info struct {
	Description string
	Category    string
	Tags   []string
}

//var All map[string]Info

//var ErrDuplicateCode = errors.New("duplicate code")

//func Add(emoji ...Emoji) error {
//	errDuplicate := func(code string) error{
//		return errors.WithMessage(ErrDuplicateCode, code)
//	}
//	for _, e := range emoji {
//		if e.Code != "" {
//			if _, ok := All[e.Code]; ok {
//				return errDuplicate(e.Code)
//			}
//			All[e.Code] = e
//		}
//		for _, code := range e.Codes {
//			if _, ok := All[e.Code]; ok {
//				return errDuplicate(code)
//			}
//			All[code] = e
//		}
//	}
//	return nil
//}
