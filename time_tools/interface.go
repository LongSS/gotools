package time_tools

import "time"

type TimUtilsInterface interface {
	// GetStandardTime 获取标准格式时间字符串
	GetStandardTime() string
	// TimeToStr 获取时间戳用time.Now().Unix()，格式化时间用t.Format， 解析时间用time.Parse。
	TimeToStr(t time.Time, layout string) string
	// StrToTime 格式话字符串
	StrToTime(str string, layout string) time.Time
	// GetNowDate 获取当前时间是几月几号
	GetNowDate() string
	// GetNowYM 获取当前时间的年月
	GetNowYM() string
	// GetNowYMD 获取当前时间的年月日
	GetNowYMD() string
	// GetTodayEnd 获取今天结束时间点time
	GetTodayEnd() time.Time
	// TodayEndSinceNowStr 获取现在到今天结束时间点的秒数字符串
	TodayEndSinceNowStr() string
	// TodayEndSinceNow 获取现在到今天结束时间的秒数
	TodayEndSinceNow() int
	// GetYesterdayDate 获取昨天的日期
	GetYesterdayDate() string
	// GetNowTs 获取当前时间戳
	GetNowTs() int
	// GetNow 获取当前时间戳
	GetNow() int64
	// GetNowTsStr 获取当前时间戳的字符串
	GetNowTsStr() string
	// GetNowMsStr 获取当前时间戳字符串
	GetNowMsStr() string
}
