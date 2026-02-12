package utils

import (
	"time"
)

func GetZodiacSign(t time.Time) string {

	month, day := t.Month(), t.Day()

	switch {
	case (month == time.March && day >= 21) || (month == time.April && day <= 19):
		return "白羊座"
	case (month == time.April && day >= 20) || (month == time.May && day <= 20):
		return "金牛座"
	case (month == time.May && day >= 21) || (month == time.June && day <= 20):
		return "双子座"
	case (month == time.June && day >= 21) || (month == time.July && day <= 22):
		return "巨蟹座"
	case (month == time.July && day >= 23) || (month == time.August && day <= 22):
		return "狮子座"
	case (month == time.August && day >= 23) || (month == time.September && day <= 22):
		return "处女座"
	case (month == time.September && day >= 23) || (month == time.October && day <= 22):
		return "天秤座"
	case (month == time.October && day >= 23) || (month == time.November && day <= 21):
		return "天蝎座"
	case (month == time.November && day >= 22) || (month == time.December && day <= 21):
		return "射手座"
	case (month == time.December && day >= 22) || (month == time.January && day <= 19):
		return "摩羯座"
	case (month == time.January && day >= 20) || (month == time.February && day <= 18):
		return "水瓶座"
	case (month == time.February && day >= 19) || (month == time.March && day <= 20):
		return "双鱼座"
	default:
		return ""
	}
}

// 白羊座：Aries，发音：英 [ˈeəriːz]，美 [ˈeriːz]，象征公羊，展现出活力、进取心和开拓精神。
// 金牛座：Taurus，发音：英 [ˈtɔːrəs]，美 [ˈtɔːrəs]，代表公牛，寓意着稳健、坚毅与执着。
// 双子座：Gemini，发音：英 [ˈdʒemɪnaɪ]，美 [ˈdʒemənaɪ]，源于拉丁语，本意是双胞胎，体现出机智灵活、善于交流。
// 巨蟹座：Cancer，发音：英 [ˈkænsə(r)]，美 [ˈkænsər]，本意是螃蟹，象征着温柔、敏感，以及强烈的自我保护意识。
// 狮子座：Leo，发音：英 [ˈliːəʊ]，美 [ˈliːoʊ]，象征狮子，代表着自信、领导力与高贵。
// 处女座：Virgo，发音：英 [ˈvɜːɡəʊ]，美 [ˈvɜːrɡoʊ]，象征纯洁、谦逊和细致，注重细节、追求完美。
// 天秤座：Libra，发音：英 [ˈliːbrə]，美 [ˈlaɪbrə]，本意是天平，代表平衡、公正，善于协调和权衡。
// 天蝎座：Scorpio，发音：英 [ˈskɔːpiəʊ]，美 [ˈskɔːrpiəʊ]，象征蝎子，代表着神秘、敏锐，意志力极为强大。
// 射手座：Sagittarius，发音：英 [ˌsædʒɪˈteəriəs]，美 [ˌsædʒɪˈteriəs]，象征弓箭手，代表着乐观、自由和探索精神。
// 摩羯座：Capricorn，发音：英 [ˈkæprɪkɔːn]，美 [ˈkæprɪkɔːrn]，象征山羊，寓意着坚韧、自律与责任感。
// 水瓶座：Aquarius，发音：英 [əˈkweəriəs]，美 [əˈkweriəs]，象征倒水者，代表革新、独立和人道主义。
// 双鱼座：Pisces，发音：英 [ˈpaɪsiːz]，美 [ˈpaɪsiːz]，象征两条鱼，体现出浪漫、富有同情心和想象力。

var zMap = map[string]string{
	"白羊座":         "Aries",
	"金牛座":         "Taurus",
	"双子座":         "Gemini",
	"巨蟹座":         "Cancer",
	"狮子座":         "Leo",
	"处女座":         "Virgo",
	"天秤座":         "Libra",
	"天蝎座":         "Scorpio",
	"射手座":         "Sagittarius",
	"摩羯座":         "Capricorn",
	"水瓶座":         "Aquarius",
	"双鱼座":         "Pisces",
	"Aries":       "白羊座",
	"Taurus":      "金牛座",
	"Gemini":      "双子座",
	"Cancer":      "巨蟹座",
	"Leo":         "狮子座",
	"Virgo":       "处女座",
	"Libra":       "天秤座",
	"Scorpio":     "天蝎座",
	"Sagittarius": "射手座",
	"Capricorn":   "摩羯座",
	"Aquarius":    "水瓶座",
	"Pisces":      "双鱼座",
}

func GetZodiacSignLangChange(cn string) string {
	return zMap[cn]
}
