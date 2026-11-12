package lint

import (
	"path"
	"regexp"
	"strings"
)

var coreImagePattern = regexp.MustCompile(`!\[([^\]]*)\]\(`)

func mediaStyleViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	for i, line := range lines {
		if insideFence[i] {
			continue
		}
		for _, match := range coreImagePattern.FindAllStringSubmatchIndex(masked[i], -1) {
			alt := line[match[2]:match[3]]
			start := match[1]
			for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
				start++
			}
			if start < len(line) && line[start] == '<' {
				start++
			}
			end := start
			for end < len(line) && line[end] != ' ' && line[end] != '\t' && line[end] != ')' && line[end] != '>' {
				end++
			}
			source := line[start:end]
			if strings.TrimSpace(alt) == "" {
				violations = append(violations, styleViolation{
					ID: "MEDIA-001", Line: i + 1, Column: coreColumn(line, match[0]),
					Detail: "image " + source + " has empty alt text; provide a meaningful textual description",
				})
				continue
			}
			filename := coreSourceFilename(source)
			if filename != "" && strings.EqualFold(strings.TrimSpace(alt), filename) {
				violations = append(violations, styleViolation{
					ID: "MEDIA-001", Line: i + 1, Column: coreColumn(line, match[0]),
					Detail: "alt text " + alt + " repeats the source filename; describe the image content instead",
				})
			}
		}
	}
	return violations
}

func coreSourceFilename(source string) string {
	if at := strings.IndexAny(source, "?#"); at >= 0 {
		source = source[:at]
	}
	source = strings.TrimSuffix(source, "/")
	if source == "" {
		return ""
	}
	return path.Base(source)
}
