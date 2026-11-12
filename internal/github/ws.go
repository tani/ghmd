package github

// IsJSWhitespace matches JavaScript's \s and the renderer's math scanner.
func IsJSWhitespace(cp int32) bool {
	return (cp >= 9 && cp <= 13) || cp == 0x20 || cp == 0xa0 || cp == 0x1680 ||
		(cp >= 0x2000 && cp <= 0x200a) || cp == 0x2028 || cp == 0x2029 || cp == 0x202f ||
		cp == 0x205f || cp == 0x3000 || cp == 0xfeff
}
