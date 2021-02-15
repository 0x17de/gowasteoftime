package wasteoftime

import (
	"testing"
	"time"
)

func TestParseEmptyFormat(t *testing.T) {
	expected1 := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	time1, err := ParseDateWithFormat("", "")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseDate(t *testing.T) {
	expected1 := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S", "2006-01-02 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseDateWithLayout(t *testing.T) {
	expected1 := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	layout, err := ParseLayout("%Y-%m-%d %H:%M:%S")
	if err != nil {
		t.Fatalf("Failed to parse layout: %v", err)
	}
	time1, err := ParseDate(layout, "2006-01-02 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseOneDigitDate(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 5, 4, 3, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%1m-%1d %1H:%1M:%1S", "2006-3-2 5:4:3")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseDate2DigitYear(t *testing.T) {
	expected1 := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%y-%m-%d %H:%M:%S", "06-01-02 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseFraction(t *testing.T) {
	expected1 := time.Date(2006, 1, 2, 15, 4, 5, 123000000, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S%F", "2006-01-02 15:04:05.123")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339Nano), result1.Format(time.RFC3339Nano))
	}
}

func TestParseUnix(t *testing.T) {
	expected1 := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%N", "1136210645")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	time2, err := ParseDateWithFormat("%N", "1136210645000")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()
	result2 := time2.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339Nano), result1.Format(time.RFC3339Nano))
	}
	if !expected1.Equal(result2) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339Nano), result1.Format(time.RFC3339Nano))
	}
}

func TestParseShortNames(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%b-%d %a %H:%M:%S", "2006-Mar-02 Thu 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseLongNames(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%b-%d %a %H:%M:%S", "2006-March-02 Thursday 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseNamesIgnoreCase(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%b-%d %a %H:%M:%S", "2006-mar-02 thu 15:04:05")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParse12AM(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 0, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S %p", "2006-03-02 12:04:05 AM")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParse12PM(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 12, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S %p", "2006-03-02 12:04:05 PM")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParse3PM(t *testing.T) {
	expected1 := time.Date(2006, 3, 2, 15, 4, 5, 0, time.UTC)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S %p", "2006-03-02 03:04:05 PM")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}

func TestParseTimezone(t *testing.T) {
	est, err := time.LoadLocation("EST")
	if err != nil {
		t.Fatalf("Failed to get timezone EST: %v", err)
	}
	expected1 := time.Date(2006, 3, 2, 15, 4, 5, 0, est)
	time1, err := ParseDateWithFormat("%Y-%m-%d %H:%M:%S %z", "2006-03-02 15:04:05 EST")
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}
	result1 := time1.Time()

	if !expected1.Equal(result1) {
		t.Fatalf("Dates don't match. %s != %s", expected1.Format(time.RFC3339), result1.Format(time.RFC3339))
	}
}
