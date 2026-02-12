package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 判断是否为连号（三连号及以上）
func IsConsecutiveNumber(number string, length int) (string, bool) {
	if len(number) < length {
		return "", false
	}
	start := 0
	pre := number[start]
	count := 1
	for i := 1; i < len(number); i++ {
		if number[i] == pre {
			count++
			if count >= length {
				return number[start : i+1], true
			}
		} else {
			start = i
			pre = number[i]
			count = 1
		}
	}
	return "", false
}

// 判断是否为顺子号（三位及以上连续数字）
func IsSequentialNumber(number string, length int) (string, bool) {
	if len(number) < length {
		return "", false
	}
	start := 0
	count := 1
	pre, _ := strconv.Atoi(string(number[start]))
	for i := 1; i < len(number); i++ {
		cur, _ := strconv.Atoi(string(number[i]))
		//fmt.Println(cur, pre, count)
		if cur != pre+1 {
			start = i
			count = 1
		} else {
			count++
			if count >= length {
				return number[start : i+1], true
			}
		}
		pre = cur
	}
	return "", false
}

// 判断是否AABB形式的靓号
func IsAABBNumber(number string) (string, bool) {
	if len(number) < 4 {
		return "", false
	}
	for i := 0; i < len(number)-3; i++ {
		if number[i] == number[i+1] && number[i+2] == number[i+3] {
			return number[i : i+4], true
		}
	}
	return "", false
}

// 判断是否为ABAB形式的靓号
func IsABABNumber(number string) (string, bool) {
	if len(number) < 4 {
		return "", false
	}
	for i := 0; i < len(number)-3; i++ {
		if number[i] == number[i+2] && number[i+1] == number[i+3] {
			return number[i : i+4], true
		}
	}
	return "", false
}

// 判断是否生日靓号
func IsBirthdayNumber(number string) (string, bool) {
	// 正则表达式匹配 01-12 月份和 01-31 日期
	re := regexp.MustCompile(`(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])`)
	matchedAll := re.FindAllString(number, -1)
	if len(matchedAll) > 0 {
		fmt.Println(matchedAll)
		for _, date := range matchedAll {
			month, _ := strconv.Atoi(date[0:2])
			day, _ := strconv.Atoi(date[2:])

			// 假设年份为 2000 年，因为 2000 年是闰年，可以覆盖所有情况
			t := time.Date(2000, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			return date, t.Year() == 2000 && int(t.Month()) == month && t.Day() == day
		}

	}
	return "", false
}

func IsLove(number string) (string, bool) {
	if strings.Contains(number, "520") {
		return "520", true
	}
	if strings.Contains(number, "1314") {
		return "1314", true
	}
	if strings.Contains(number, "3344") {
		return "3344", true
	}
	return "", false
}
