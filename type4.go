package date2timestamp

import (
	"github.com/mostafah/go-jalali/jalali"
	"strconv"
	"strings"
)

func Type4ToTimestamp(val string) int64 {
	chunked := strings.Split(val, " ")
	h, mm, s := GetTime(chunked[4])
	d, _ := strconv.Atoi(chunked[0])
	y, _ := strconv.Atoi(chunked[2])

	m := 0
	for index := 0; chunked[1] != MONTH_NAMES_FA[index]; index++ {
		m = index + 2
	}
	return jalali.Jtog(y, m, d, h, mm, s).Unix()
}
