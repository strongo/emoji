package emoji

const (

	FlagRUCode = "ru"
	FlagRUUnicode = "🇷🇺"

	FlagUSCode = "us"
	FlagUSUnicode = "🇺🇸"

	FlagUnitedKingdomCode = "uk"
	FlagUnitedKingdomUnicode = "🇬🇧"

)

var Flags = Slice {
	Emoji{
		Code: FlagRUCode,
		Codes: []string{"russian_federation"},
		Unicode: FlagRUUnicode,
	},
	Emoji{
		Code: FlagUSCode,
		Codes: []string{"uss"},
		Unicode: FlagUSUnicode,
	},
	Emoji{
		Code: FlagUnitedKingdomCode,
		Codes: []string{"united_kingdom"},
		Unicode: FlagUnitedKingdomUnicode,
	},
}