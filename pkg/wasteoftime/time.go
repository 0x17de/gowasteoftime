package wasteoftime

import (
	"container/list"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	g_ParseTimePatternMatcher = regexp.MustCompile("%(1?[^%])|([^%]+)")
	g_Months                  = []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	g_MonthsAbbrev = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	g_Days = []string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}
	g_DaysAbbrev = []string{
		"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun",
	}
)

type TTimeData struct {
	tz     *time.Location
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
	nsec   int
}

func newTimeData() *TTimeData {
	return &TTimeData{
		tz:    time.UTC,
		year:  1970,
		month: 1,
		day:   1,
	}
}
func (td *TTimeData) Time() time.Time {
	return time.Date(td.year, time.Month(td.month), td.day, td.hour, td.minute, td.second, td.nsec, td.tz)
}

type TTimeLayout struct {
	dateMatcher *regexp.Regexp
	parsers     *list.List
}

func ParseLayout(format string) (layout *TTimeLayout, err error) {
	var restr strings.Builder
	parsers := list.New()
	layout = &TTimeLayout{
		parsers: parsers,
	}
	restr.WriteString("(?i)") // ignore case

	for _, item := range g_ParseTimePatternMatcher.FindAllStringSubmatch(format, -1) {
		if item[1] != "" {
			switch item[1][0] {
			case '1': // one digit numbers
				switch item[1][1] {
				case 'm':
					restr.WriteString("([0-9]?[0-9])")
					parsers.PushBack(func(match string, data *TTimeData) (err error) {
						data.month, err = strconv.Atoi(match)
						return
					})
				case 'd':
					restr.WriteString("([0-9]?[0-9])")
					parsers.PushBack(func(match string, data *TTimeData) (err error) {
						data.day, err = strconv.Atoi(match)
						return
					})
				case 'H':
					restr.WriteString("([0-9]?[0-9])")
					parsers.PushBack(func(match string, data *TTimeData) (err error) {
						data.hour, err = strconv.Atoi(match)
						return
					})
				case 'M':
					restr.WriteString("([0-9]?[0-9])")
					parsers.PushBack(func(match string, data *TTimeData) (err error) {
						data.minute, err = strconv.Atoi(match)
						return
					})
				case 'S':
					restr.WriteString("([0-9]?[0-9])")
					parsers.PushBack(func(match string, data *TTimeData) (err error) {
						data.second, err = strconv.Atoi(match)
						return
					})
				}
			case 'y':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					data.year, err = strconv.Atoi(match)
					if data.year >= 70 { // gt 70 -> 1970-1999
						data.year += 1900
					} else {
						data.year += 2000
					}
					return
				})
			case 'Y':
				restr.WriteString("([0-9]{4})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.year, err = strconv.Atoi(match); return })
			case 'm':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.month, err = strconv.Atoi(match); return })
			case 'd':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.day, err = strconv.Atoi(match); return })
			case 'H':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.hour, err = strconv.Atoi(match); return })
			case 'M':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.minute, err = strconv.Atoi(match); return })
			case 'S':
				restr.WriteString("([0-9]{2})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) { data.second, err = strconv.Atoi(match); return })
			case 'F':
				restr.WriteString("(?:.([0-9]{0,9}))?")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					fmt.Printf("NSEC: %s\n", match)
					if match != "" {
						nsec, err := strconv.Atoi(match)
						if err != nil {
							return err
						}
						if len(match) < 9 {
							nsec *= int(math.Pow(float64(10), float64(9-len(match))))
						}
						data.nsec = nsec
					}
					return nil
				})
			case 'p':
				restr.WriteString("(AM|PM)")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					if strings.EqualFold(match, "PM") {
						if data.hour < 12 {
							data.hour = (data.hour + 12) % 24
						}
					} else {
						if data.hour == 12 {
							data.hour = 0
						}
					}
					return nil
				})
			case 'z':
				restr.WriteString("([A-Z]{3})")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					data.tz, err = time.LoadLocation(match)
					return
				})
			case 'a':
				restr.WriteString("(?:")
				restr.WriteString(strings.Join(g_DaysAbbrev, "|"))
				restr.WriteString("|")
				restr.WriteString(strings.Join(g_Days, "|"))
				restr.WriteString(")")
				// will not modify any data
			case 'b':
				restr.WriteString("(")
				restr.WriteString(strings.Join(g_MonthsAbbrev, "|"))
				restr.WriteString("|")
				restr.WriteString(strings.Join(g_Months, "|"))
				restr.WriteString(")")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					for i := 0; i < len(g_MonthsAbbrev); i++ {
						if strings.EqualFold(match, g_MonthsAbbrev[i]) {
							data.month = i + 1
							return nil
						}
					}
					for i := 0; i < len(g_Months); i++ {
						if strings.EqualFold(match, g_Months[i]) {
							data.month = i + 1
							return nil
						}
					}
					return nil
				})
			case 'N': // unix timestamp
				restr.WriteString("([0-9]{10}(?:[0-9]{3})?)")
				parsers.PushBack(func(match string, data *TTimeData) (err error) {
					var unix int64
					var t time.Time
					unix, err = strconv.ParseInt(match, 10, 64)
					if len(match) == 13 { // including millisecs
						t = time.Unix(unix/1000, 0)
					} else {
						t = time.Unix(unix, 0)
					}
					data.year = t.Year()
					data.month = int(t.Month())
					data.day = t.Day()
					data.hour = t.Hour()
					data.minute = t.Minute()
					data.second = t.Second()
					return
				})
			default:
				return nil, fmt.Errorf("Format specifier %%%c is invalid", item[1][0])
			}
		} else {
			restr.WriteString(regexp.QuoteMeta(item[2]))
		}
	}
	layout.dateMatcher, err = regexp.Compile(restr.String())
	return layout, err
}

func ParseDateWithFormat(format string, date string) (result *TTimeData, err error) {
	layout, err := ParseLayout(format)
	if err != nil {
		return nil, fmt.Errorf("Could not parse date format: %v", err)
	}

	match := layout.dateMatcher.FindStringSubmatch(date)
	if match == nil {
		return nil, fmt.Errorf("Could not parse date %s with format %s\n", date, format)
	}

	result = newTimeData()
	parser := layout.parsers.Front()
	for i := 1; i <= layout.parsers.Len(); i++ {
		if parser.Value != nil {
			fn, _ := parser.Value.(func(string, *TTimeData) error)
			err = fn(match[i], result)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse date: %v", err)
			}
		}
		parser = parser.Next()
	}
	return result, nil
}

func ParseDate(layout *TTimeLayout, date string) (result *TTimeData, err error) {
	match := layout.dateMatcher.FindStringSubmatch(date)
	if match == nil {
		return nil, fmt.Errorf("Could not parse date")
	}

	result = newTimeData()
	parser := layout.parsers.Front()
	for i := 1; i <= layout.parsers.Len(); i++ {
		if parser.Value != nil {
			fn, _ := parser.Value.(func(string, *TTimeData) error)
			err = fn(match[i], result)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse date: %v", err)
			}
		}
		parser = parser.Next()
	}
	return result, nil
}
