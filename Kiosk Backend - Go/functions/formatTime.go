package function

import (
	"fmt"
	"strings"
	"time"
)

var formats = map[int]string{
	1:   "01/02/06",
	2:   "06.01.02",
	3:   "02/01/06",
	4:   "02.01.06",
	5:   "02-01-06",
	6:   "02 Jan 06",
	7:   "Jan 02, 06",
	10:  "01-02-06",
	11:  "06/01/02",
	12:  "060102",
	101: "01/02/2006",
	102: "2006.01.02",
	103: "02/01/2006",
	104: "02.01.2006",
	105: "02-01-2006",
	106: "02 Jan 2006",
	107: "Jan 02, 2006",
	110: "01-02-2006",
	111: "2006/01/02",
	112: "20060102",
	120: "2006-01-02 15:04:05",
	121: "2006-01-02 15:04:05.000",
	126: "2006-01-02T15:04:05.000Z",
}

func FormatTime(t time.Time, formatNumber int) string {
	format, found := formats[formatNumber]
	if !found {
		return "Invalid format number"
	}
	return t.Format(format)
}

func GetTime(t string, formatNumber int) (time.Time, error) {
	layout, found := formats[formatNumber]
	if !found {
		return time.Time{}, fmt.Errorf("invalid format number")
	}
	return time.Parse(layout, t)
}

func AddTimeToTodayDate(timeString string) (time.Time, error) {
	now := time.Now()
	timeString = strings.ReplaceAll(timeString, " ", "")
	timeString = strings.ReplaceAll(timeString, ":", "")
	if len(timeString) == 4 {
		timeString = timeString + "00"
	}
	var layout string = "150405"

	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s", timeString)
	}
	combinedTime := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, now.Location())
	return combinedTime, nil
}

func AddTimeToDate(dateString string, timeString string) (time.Time, error) {
	if len(timeString) == 4 {
		timeString = timeString + "00"
	}
	dateTimeString := dateString + timeString
	var layout string = "20060102150405"
	parsedTime, err := time.Parse(layout, dateTimeString)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %s", timeString)
	}
	return parsedTime, nil
}
