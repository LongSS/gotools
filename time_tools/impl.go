package time_tools

import (
	"fmt"
	"strconv"
	"time"
)

type TimeUtilsImpl struct {
}

func (impl TimeUtilsImpl) GetStandardTime() string {
	return impl.TimeToStr(time.Now(), "2006-01-02 15:04:05")
}

// TimeToStr 获取时间戳用time.Now().Unix()，格式化时间用t.Format，解析时间用time.Parse。
func (impl TimeUtilsImpl) TimeToStr(t time.Time, layout string) string {
	str := time.Time.Format(t, layout)
	return str
}
func (impl TimeUtilsImpl) StrToTime(str string, layout string) time.Time {
	ts, _ := time.ParseInLocation(layout, str, time.Local)
	return ts
}
func (impl TimeUtilsImpl) GetNowDate() string {
	return impl.TimeToStr(time.Now(), "0102")
}
func (impl TimeUtilsImpl) GetNowYM() string {
	return impl.TimeToStr(time.Now(), "200601")
}
func (impl TimeUtilsImpl) GetNowYMD() string {
	return impl.TimeToStr(time.Now(), "20060102")
}
func (impl TimeUtilsImpl) GetTodayEnd() time.Time {
	tm := time.Now().AddDate(0, 0, 1)
	dateStr := impl.TimeToStr(tm, "20060102")
	d := impl.StrToTime(dateStr, "20060102")
	return d
}
func (impl TimeUtilsImpl) TodayEndSinceNowStr() string {
	d := impl.GetTodayEnd()
	ttl := d.Sub(time.Now()).Seconds()
	s := strconv.FormatFloat(ttl, 'f', 0, 64)
	return s
}
func (impl TimeUtilsImpl) TodayEndSinceNow() int {
	d := impl.GetTodayEnd()
	ttl := d.Sub(time.Now()).Seconds()
	return int(ttl)
}
func (impl TimeUtilsImpl) GetYesterdayDate() string {
	return impl.TimeToStr(time.Now().AddDate(0, 0, -1), "0102")
}
func (impl TimeUtilsImpl) GetNowTs() int {
	return int(time.Now().Unix())
}
func (impl TimeUtilsImpl) GetNow() int64 {
	return time.Now().Unix()
}
func (impl TimeUtilsImpl) GetNowTsStr() string {
	r := strconv.Itoa(int(time.Now().Unix()))
	return r
}
func (impl TimeUtilsImpl) GetNowMsStr() string {
	ms := time.Now().UnixNano() / 1e3
	r := strconv.FormatInt(ms, 10)
	return r
}

type JSONTime time.Time

func (p JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(p).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
