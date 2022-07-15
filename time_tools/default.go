package time_tools

import (
	"fmt"
	"strconv"
	"time"
)

func GetStandardTime() string {
	return TimeToStr(time.Now(), "2006-01-02 15:04:05")
}

// TimeToStr 获取时间戳用time.Now().Unix()，格式化时间用t.Format，解析时间用time.Parse。
func TimeToStr(t time.Time, layout string) string {
	str := time.Time.Format(t, layout)
	return str
}
func StrToTime(str string, layout string) time.Time {
	//ts, _ := time.Parse(layout, str)
	ts, _ := time.ParseInLocation(layout, str, time.Local)
	return ts
}
func GetNowDate() string {
	return TimeToStr(time.Now(), "0102")
}
func GetNowYM() string {
	return TimeToStr(time.Now(), "200601")
}
func GetNowYMD() string {
	return TimeToStr(time.Now(), "20060102")
}
func GetTodayEnd() time.Time {
	tm := time.Now().AddDate(0, 0, 1)
	dateStr := TimeToStr(tm, "20060102")
	d := StrToTime(dateStr, "20060102")
	return d
}
func TodayEndSinceNowStr() string {
	d := GetTodayEnd()
	ttl := d.Sub(time.Now()).Seconds()
	s := strconv.FormatFloat(ttl, 'f', 0, 64)
	return s
}
func TodayEndSinceNow() int {
	d := GetTodayEnd()
	ttl := d.Sub(time.Now()).Seconds()
	return int(ttl)
}
func GetYesterdayDate() string {
	return TimeToStr(time.Now().AddDate(0, 0, -1), "0102")
}
func GetNowTs() int {
	return int(time.Now().Unix())
}
func GetNow() int64 {
	return time.Now().Unix()
}
func GetNowTsStr() string {
	r := strconv.Itoa(int(time.Now().Unix()))
	return r
}
func GetNowMsStr() string {
	ms := time.Now().UnixNano() / 1e3
	r := strconv.FormatInt(ms, 10)
	return r
}

type JSONTime time.Time

func (p JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(p).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
