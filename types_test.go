package date2timestamp

import (
	"github.com/mavihq/persian"
	"testing"
)

func Test_Type2(t *testing.T) {
	timestamp := Type2ToTimestamp(persian.ToEnglishDigits("۱۳۹۹/۰۹/۰۲ - ۱۵:۴۰"))
	if timestamp > 0 {
		t.Log("success", timestamp)
	} else {
		t.Failed()
	}
}

func BenchmarkType2ToTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Type2ToTimestamp(persian.ToEnglishDigits("۱۳۹۹/۰۹/۰۲ - ۱۵:۴۰"))
	}
}

func Test_Type4(t *testing.T) {
	timestamp := Type4ToTimestamp(persian.ToEnglishDigits("۰۳ آذر ۱۳۹۹ - ۱۴:۰۵"))
	if timestamp > 0 {
		t.Log("success", timestamp)
	} else {
		t.Error("")
	}
}

func BenchmarkType4ToTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Type4ToTimestamp(persian.ToEnglishDigits("۰۳ آذر ۱۳۹۹ - ۱۴:۰۵"))
	}
}

func Test_Type5(t *testing.T) {
	timestamp := Type5ToTimestamp(persian.ToEnglishDigits("شنبه ۱۷ آبان ۱۳۹۹ ساعت ۲۰:۱۴"))
	if timestamp > 0 {
		t.Log("success", timestamp)
	} else {
		t.Error("")
	}
}

func BenchmarkType5ToTimestamp(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Type5ToTimestamp(persian.ToEnglishDigits("شنبه ۱۷ آبان ۱۳۹۹ ساعت ۲۰:۱۴"))
	}
}
