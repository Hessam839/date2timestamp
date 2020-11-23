package date2timestamp

import (
	"github.com/mostafah/go-jalali/jalali"
	"strings"
)

func Type2ToTimestamp(val string) int64 {
	chuncked := strings.Split(val, " ")
	y, m, d := GetDate(chuncked[0])
	h, mm, s := GetTime(chuncked[2])
	return jalali.Jtog(y, m, d, h, mm, s).Unix()
}
