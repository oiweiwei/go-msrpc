//go:generate go run ./gen/gen.go -o lang.go -pkg lcid -url https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/63d3d639-7fd2-4afb-abbe-0d5b5551eef8
package lcid

import "strings"

// LCID represents a Locale Identifier (LCID) used in Windows to identify a specific language and
// regional settings.
type LCID uint32

var tagToLang = map[string]LCID{}

func init() {
	for lcid, lang := range langToTag {
		tagToLang[lang] = lcid
	}
}

// LangToLCID returns the LCID for a given language tag. If the language tag is not found, it returns false.
// If a sort order is provided, it will be applied to the LCID.
func LangToLCID(lang string, sort ...LCID) (LCID, bool) {
	lcid, ok := tagToLang[strings.ToLower(lang)]

	if ok && len(sort) > 0 {
		lcid |= sort[0]
	}

	return lcid, ok
}
