package helper

import "time"

var bangkokLocalTime *time.Location

func init() {
	local, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	bangkokLocalTime = local
}

func ParseBangkokTime(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, bangkokLocalTime)
}
