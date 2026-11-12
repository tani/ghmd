package github

import "strings"

func unicodeInRanges(cp rune, ranges []unicodeRange) bool {
	lo, hi := 0, len(ranges)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if cp < ranges[mid].first {
			hi = mid
		} else if cp > ranges[mid].last {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func unicodeCased(cp rune) bool {
	return unicodeInRanges(cp, unicodeCasedRanges)
}

func unicodeIgnorable(cp rune) bool {
	return unicodeInRanges(cp, unicodeIgnorableRanges)
}

// unicodeLowerDecode matches cmark_utf8proc_iterate. Malformed input consumes
// one source byte in unicodeLower, as it does in the C implementation.
func unicodeLowerDecode(s string, pos int) (rune, int, bool) {
	first := s[pos]
	width := 0
	switch {
	case first <= 0x7f:
		width = 1
	case first >= 0xc2 && first <= 0xdf:
		width = 2
	case first >= 0xe0 && first <= 0xef:
		width = 3
	case first >= 0xf0 && first <= 0xf7:
		width = 4
	default:
		return 0, 0, false
	}
	if width > len(s)-pos {
		return 0, 0, false
	}
	for i := 1; i < width; i++ {
		if s[pos+i]&0xc0 != 0x80 {
			return 0, 0, false
		}
	}
	var cp rune
	switch width {
	case 1:
		cp = rune(first)
	case 2:
		cp = rune(first&0x1f)<<6 | rune(s[pos+1]&0x3f)
	case 3:
		cp = rune(first&0x0f)<<12 | rune(s[pos+1]&0x3f)<<6 | rune(s[pos+2]&0x3f)
		if cp < 0x800 || cp >= 0xd800 && cp < 0xe000 {
			return 0, 0, false
		}
	case 4:
		cp = rune(first&0x07)<<18 | rune(s[pos+1]&0x3f)<<12 |
			rune(s[pos+2]&0x3f)<<6 | rune(s[pos+3]&0x3f)
		if cp < 0x10000 || cp >= 0x110000 {
			return 0, 0, false
		}
	}
	return cp, width, true
}

func unicodeFollowedByCased(s string, pos int) bool {
	for pos < len(s) {
		cp, width, ok := unicodeLowerDecode(s, pos)
		if !ok {
			return false
		}
		if !unicodeIgnorable(cp) {
			return unicodeCased(cp)
		}
		pos += width
	}
	return false
}

// unicodeLower applies Unicode 17 default lowercase mappings, including the
// context-sensitive final sigma rule, to one UTF-8 string.
func unicodeLower(value string) string {
	var out strings.Builder
	out.Grow(len(value))
	precedingCased := false
	for pos := 0; pos < len(value); {
		cp, width, ok := unicodeLowerDecode(value, pos)
		if !ok {
			out.WriteString("\xef\xbf\xbd")
			precedingCased = false
			pos++
			continue
		}
		if cp == 0x3a3 && precedingCased && !unicodeFollowedByCased(value, pos+width) {
			out.WriteString("\xcf\x82")
		} else {
			lo, hi := 0, len(unicodeLowerMappings)
			for lo < hi {
				mid := lo + (hi-lo)/2
				if uint32(cp) > unicodeLowerMappings[mid].cp {
					lo = mid + 1
				} else {
					hi = mid
				}
			}
			if lo < len(unicodeLowerMappings) && unicodeLowerMappings[lo].cp == uint32(cp) {
				offset := int(unicodeLowerMappings[lo].offset)
				length := int(unicodeLowerMappings[lo].length)
				out.WriteString(unicodeLowerBytes[offset : offset+length])
			} else {
				out.WriteString(value[pos : pos+width])
			}
		}
		if !unicodeIgnorable(cp) {
			precedingCased = unicodeCased(cp)
		}
		pos += width
	}
	return out.String()
}
