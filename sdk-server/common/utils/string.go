package utils

import (
	"regexp"
	"strings"
)

// func NewSB() *StringBuilder {
// 	return &StringBuilder{}
// }

// type StringBuilder struct {
// 	build strings.Builder
// }

// func (s *StringBuilder) Append(str string) *StringBuilder {
// 	s.build.WriteString(str)
// 	return s
// }

// func (s *StringBuilder) AppendBytes(bs []byte) *StringBuilder {
// 	s.build.Write(bs)
// 	return s
// }

// func (s *StringBuilder) AppendByte(c byte) *StringBuilder {
// 	s.build.WriteByte(c)
// 	return s
// }

// func (s *StringBuilder) AppendRune(r rune) *StringBuilder {
// 	s.build.WriteRune(r)
// 	return s
// }

// func (s *StringBuilder) Len() int {
// 	return s.build.Len()
// }

// func (s *StringBuilder) String() string {
// 	return s.build.String()
// }

func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

// RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}

func GetFileExtension(filename string) string {
	lastDotIndex := -1
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			lastDotIndex = i
			break
		}
	}
	if lastDotIndex != -1 {
		return filename[lastDotIndex+1:]
	}
	return ""
}
