package date2timestamp

import (
	"regexp"
	"strconv"
	"strings"
)

func GetTime(val string) (hour int, minute int, second int) {
	chuncked := strings.Split(val, ":")
	if len(chuncked) == 3 {
		hour, _ = strconv.Atoi(chuncked[0])
		minute, _ = strconv.Atoi(chuncked[1])
		second, _ = strconv.Atoi(chuncked[2])
	}
	if len(chuncked) == 2 {
		hour, _ = strconv.Atoi(chuncked[0])
		minute, _ = strconv.Atoi(chuncked[1])
		second = 0
	}
	return
}

func GetDate(val string) (year int, month int, day int) {
	dash := regexp.MustCompile(`-`)
	slash := regexp.MustCompile(`/`)
	var chunked []string
	switch {
	case dash.MatchString(val):
		chunked = strings.Split(val, "-")
		//if len(chunked[0]) == 4{
		//	year,_ = strconv.Atoi(chunked[0])
		//	month,_ = strconv.Atoi(chunked[1])
		//	day,_ = strconv.Atoi(chunked[2])
		//}
		//if len(chunked[2]) == 4{
		//	year,_ = strconv.Atoi(chunked[2])
		//	month,_ = strconv.Atoi(chunked[1])
		//	day,_ = strconv.Atoi(chunked[0])
		//}
		//return
	case slash.MatchString(val):
		chunked = strings.Split(val, "/")
		//if len(chunked[0]) == 4{
		//	year,_ = strconv.Atoi(chunked[0])
		//	month,_ = strconv.Atoi(chunked[1])
		//	day,_ = strconv.Atoi(chunked[2])
		//}
		//if len(chunked[2]) == 4{
		//	year,_ = strconv.Atoi(chunked[2])
		//	month,_ = strconv.Atoi(chunked[1])
		//	day,_ = strconv.Atoi(chunked[0])
		//}
		//return
	}
	if len(chunked[0]) == 4 {
		year, _ = strconv.Atoi(chunked[0])
		month, _ = strconv.Atoi(chunked[1])
		day, _ = strconv.Atoi(chunked[2])
	}
	if len(chunked[2]) == 4 {
		year, _ = strconv.Atoi(chunked[2])
		month, _ = strconv.Atoi(chunked[1])
		day, _ = strconv.Atoi(chunked[0])
	}
	return
}

func chunkedDate(chunk []string, offset int) (year int, month int, day int) {
	if len(chunk[0]) == 4 {
		year, _ = strconv.Atoi(chunk[0])
		month, _ = strconv.Atoi(chunk[1])
		day, _ = strconv.Atoi(chunk[2])
	}
	if len(chunk[2]) == 4 {
		year, _ = strconv.Atoi(chunk[2])
		month, _ = strconv.Atoi(chunk[1])
		day, _ = strconv.Atoi(chunk[0])
	}
	return
}
