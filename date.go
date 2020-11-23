package date2timestamp

import (
	"fmt"
	"github.com/mavihq/persian"
	"regexp"
	"time"
)

//func ToTimestamp(val string) int64 {
//	val = persian.ToEnglishDigits(persian.FixArabic(val))
//	test := strings.Split(val, " ")
//	res := []string{}
//	for _, i := range test {
//		t := true
//		if i == ":" {
//			t = false
//		}
//		if (i == "ساعت") || (i == "-") {
//			t = false
//		}
//		//if WEEKDAY_NAMES_FA[i] {
//		//	t = false
//		//}
//		if t {
//			res = append(res, i)
//		}
//	}
//		for index,i := range MONTH_NAMES_FA {
//			for j := 0; j < len(res); j++ {
//				if i == res[j]{
//					res[j] = fmt.Sprint(index)
//				}
//			}
//		}
//		var hour,minute,second int
//		for _, part := range res {
//			TwoDigitClock, _ := regexp.Compile(`\d{2}:\d{2}`)
//			if TwoDigitClock.MatchString(part) {
//				chuncked := strings.Split(part,":")
//				hour,_ = strconv.Atoi(chuncked[0])
//				minute,_ = strconv.Atoi(chuncked[1])
//				second = 0
//			}
//			ThreeDigitClock, _ := regexp.Compile(`\d{2}:\d{2}:\d{2}`)
//			if ThreeDigitClock.MatchString(part){
//				chuncked := strings.Split(part,":")
//				hour,_ = strconv.Atoi(chuncked[0])
//				minute,_ = strconv.Atoi(chuncked[1])
//				second,_ = strconv.Atoi(chuncked[2])
//			}
//		}
//		var date []string
//		for _,j := range res {
//			matched, err := regexp.MatchString("/", j)
//			if err != nil {
//				log.Printf(err.Error())
//			}
//			if matched{
//				date = strings.Split(j, "/")
//			}
//		}
//		var year, month, day int
//		if len(date) == 3 {
//			if len(date[0]) == 4 {
//				year,_ = strconv.Atoi(date[0])
//				month,_ = strconv.Atoi(date[1])
//				day,_ = strconv.Atoi(date[2])
//			}
//			if len(date[2]) == 4 {
//				year,_ = strconv.Atoi(date[2])
//				month,_ = strconv.Atoi(date[1])
//				day,_ = strconv.Atoi(date[0])
//			}
//		}
//		tt :=jalali.Jtog(year, month, day, hour, minute, second)
//		fmt.Println(tt.Unix())
//	return 0
//}

//ToTimestamp convert date in string form to timestamp
func ToTimestamp(val string) int64 {
	months := `(فروردین|اردیبهشت|خرداد|تیر|مرداد|شهریور|مهر|آبان|آذر|دی|بهمن|اسفند)`
	days := `(شنبه|یکشنبه|دوشنبه|سه‌شنبه|سه شنبه|چهارشنبه|پنج‌شنبه|پنچ شنبه)`
	dateType1 := `(?i)^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`
	dateType2 := `(?i)^\d{4}/\d{2}/\d{2}\s-\s\d{2}:\d{2}`
	dateType3 := `(?i)^\d{4}-\d{2}-\d{2}`
	dateType4 := fmt.Sprintf(`(?im)^\d{2}\s%s\s\d{4}\s-\s\d{2}:\d{2}`, months)
	dateType5 := fmt.Sprintf(`(?im)^%s\s\d{2}\s%s\s\d{4}\sساعت\s\d{2}:\d{2}`, days, months)
	dateType6 := `(?im)^\d{2}/\d{2}/\d{4}\s\d{1,2}:\d{2}:\d{2}\s(AM|PM)`

	val = persian.ToEnglishDigits(persian.FixArabic(val))
	regex1 := regexp.MustCompile(dateType1)
	regex3 := regexp.MustCompile(dateType3)
	regex2 := regexp.MustCompile(dateType2)
	//regex4 := regexp.MustCompile(`(?im)^[\p{N}]{4}/[\p{N}]{2}/[\p{N}]{2}\s-\s[\p{N}]{2}:[\p{N}]{2}`)
	regex4 := regexp.MustCompile(dateType4)
	regex5 := regexp.MustCompile(dateType5)
	regex6 := regexp.MustCompile(dateType6)

	switch {
	case regex1.MatchString(val):
		tt, _ := time.Parse(time.RFC3339, val)
		return tt.Unix()
	case regex2.MatchString(val):
		return Type2ToTimestamp(val)
	case regex3.MatchString(val):
		tt, _ := time.Parse(`2006-01-02`, val)
		return tt.Unix()
	case regex4.MatchString(val):
		return Type4ToTimestamp(val)
	case regex5.MatchString(val):
		return Type5ToTimestamp(val)
	case regex6.MatchString(val):
		tt, _ := time.Parse(`01/02/2006 3:04:05 PM`, val)
		return tt.Unix()
	}
	return 0
}
