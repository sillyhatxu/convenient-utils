package gotime

import "time"

const (
	SingaporeLocal                  = "Asia/Singapore"
	AsiaSingaporeHour time.Duration = 8
	LayoutTime                      = "2006-01-02 15:04:05"
)

func Local(name string) (*time.Location, error) {
	l, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return nil, err
	}
	return l, nil
}

func ParseTimeInLocation(value string) (time.Time, error) {
	local, err := Local(SingaporeLocal)
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation(LayoutTime, value, local)
}

func FormatLocation(t time.Time) string {
	return t.Format(LayoutTime)
}

func FormatTimestamp(i int64) time.Time {
	return time.Unix(0, i*int64(time.Millisecond))
}

func ToLocationMillisecond(t time.Time) int64 {
	return TimeFormat(t).UnixNano() / int64(time.Millisecond)
}

func TimeFormat(inputTime time.Time) time.Time {
	secondsEastOfUTC := int((AsiaSingaporeHour * time.Hour).Seconds())
	timezoneAsiaSingapore := time.FixedZone(SingaporeLocal, secondsEastOfUTC)
	return time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), inputTime.Hour(), inputTime.Minute(), inputTime.Second(), inputTime.Nanosecond(), timezoneAsiaSingapore)
}
