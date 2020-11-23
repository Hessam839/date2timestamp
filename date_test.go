package date2timestamp

import (
	"fmt"
	"testing"
)

//func Test_ToTimestamp(t *testing.T)  {
//	ToTimestamp("۱۳۹۹/۰۹/۰۲ - ۱۵:۴۰")
//	ToTimestamp("2020-11-07")
//	ToTimestamp("شنبه ۱۷ آبان ۱۳۹۹ ساعت ۲۰:۱۴")
//}

func Test_ToTimestamp(t *testing.T) {
	fmt.Println(ToTimestamp("۱۳۹۹/۰۹/۰۲ - ۱۵:۴۰"))
	fmt.Println(ToTimestamp("2020-11-07"))
	fmt.Println(ToTimestamp("2020-11-23T10:49:00Z"))
	fmt.Println(ToTimestamp("شنبه ۱۷ آبان ۱۳۹۹ ساعت ۲۰:۱۴"))
	fmt.Println(ToTimestamp("۰۳ آذر ۱۳۹۹ - ۱۴:۰۵"))
	fmt.Println(ToTimestamp("11/23/2020 2:10:29 PM"))
}

func benchmarkToTimestamp(val string, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		ToTimestamp(val)
	}
}
func BenchmarkToTimestamp1(b *testing.B) { benchmarkToTimestamp("۱۳۹۹/۰۹/۰۲ - ۱۵:۴۰", b) }
func BenchmarkToTimestamp2(b *testing.B) { benchmarkToTimestamp("2020-11-07", b) }
func BenchmarkToTimestamp3(b *testing.B) { benchmarkToTimestamp("2020-11-23T10:49:00Z", b) }
func BenchmarkToTimestamp4(b *testing.B) {
	benchmarkToTimestamp("شنبه ۱۷ آبان ۱۳۹۹ ساعت ۲۰:۱۴", b)
}
func BenchmarkToTimestamp5(b *testing.B) { benchmarkToTimestamp("۰۳ آذر ۱۳۹۹ - ۱۴:۰۵", b) }
func BenchmarkToTimestamp6(b *testing.B) { benchmarkToTimestamp("11/23/2020 2:10:29 PM", b) }
