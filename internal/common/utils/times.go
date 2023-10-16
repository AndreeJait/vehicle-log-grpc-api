package utils

import (
	"strconv"
	"strings"
	"time"
)

const DefaultFormatDate = "2006-01-02"
const DefaultFormatTime = "15:04"

func NowJkrt() time.Time {
	tzJakarta, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now()
	timeNowJkrt, _ := time.ParseInLocation(time.RFC3339, timeNow.Format(time.RFC3339), tzJakarta)
	return timeNowJkrt
}

func ConvertToTime(date, timeStr string) time.Time {
	tzJakarta, _ := time.LoadLocation("Asia/Jakarta")
	timeDate, _ := time.Parse(DefaultFormatDate, date)
	year, month, day := timeDate.Date()
	hourSplit := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(hourSplit[0])
	minute, _ := strconv.Atoi(hourSplit[1])
	timeDate = time.Date(year, month, day, hour, minute, 0, 0, tzJakarta)
	return timeDate
}

func GetEndEndOfDateTime(endOfDate string) time.Duration {
	tzJakarta, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now()
	timeNowJkrt, err := time.ParseInLocation(time.RFC3339, timeNow.Format(time.RFC3339), tzJakarta)
	if err != nil {
		panic(err)
	}
	year, month, day := timeNowJkrt.Date()
	splitHour := strings.Split(endOfDate, ":")
	hour, _ := strconv.Atoi(splitHour[0])
	minute, _ := strconv.Atoi(splitHour[1])
	timeNowJkrtNowEnd := time.Date(year, month, day, hour, minute, 0, 0, tzJakarta)
	if timeNowJkrt.After(timeNowJkrtNowEnd) {
		timeNowJkrtTemp := timeNowJkrt.Add(24 * time.Hour)
		year, month, day = timeNowJkrtTemp.Date()
		timeNowJkrtNowEnd = time.Date(year, month, day, hour, minute, 0, 0, tzJakarta)
	}
	gap := timeNowJkrtNowEnd.Sub(timeNowJkrt)
	return gap
}
