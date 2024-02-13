package main

import (
	"fmt"
	"strconv"
)

func validateDateTime(dateTime string) bool {
	// Check length
	if len(dateTime) != 20 && len(dateTime) != 25 {
		return false
	}

	// Check separators
	if dateTime[4] != '-' || dateTime[7] != '-' || dateTime[10] != 'T' ||
		dateTime[13] != ':' || dateTime[16] != ':' {
		return false
	}

	// Check year, month, and day
	year, err := strconv.Atoi(dateTime[0:4])
	if err != nil || year < 0 {
		return false
	}

	month, err := strconv.Atoi(dateTime[5:7])
	if err != nil || month < 1 || month > 12 {
		return false
	}

	day, err := strconv.Atoi(dateTime[8:10])
	if err != nil || day < 1 || day > 31 {
		return false
	}

	// Check hour, minute, and second
	hour, err := strconv.Atoi(dateTime[11:13])
	if err != nil || hour < 0 || hour > 23 {
		return false
	}

	minute, err := strconv.Atoi(dateTime[14:16])
	if err != nil || minute < 0 || minute > 59 {
		return false
	}

	second, err := strconv.Atoi(dateTime[17:19])
	if err != nil || second < 0 || second > 59 {
		return false
	}

	// Check time zone designator
	if len(dateTime) == 25 && dateTime[19] != 'Z' {
		if dateTime[19] != '+' && dateTime[19] != '-' {
			return false
		}

		// Check offset hours and minutes
		offsetHours, err := strconv.Atoi(dateTime[20:22])
		if err != nil || offsetHours < 0 || offsetHours > 14 {
			return false
		}

		if dateTime[22] != ':' {
			return false
		}

		offsetMinutes, err := strconv.Atoi(dateTime[23:25])
		if err != nil || offsetMinutes < 0 || offsetMinutes > 59 {
			return false
		}
	}

	return true
}

func normalizeDateTime(dateTime string) (string, error) {
	// Parse the input date-time string
	if len(dateTime) == 20 {
		return fmt.Sprint(dateTime), nil
	} else if len(dateTime) != 25 {
		return "", fmt.Errorf("invalid date-time format: %s", dateTime)
	}

	year, err := strconv.Atoi(dateTime[0:4])
	if err != nil {
		return "", fmt.Errorf("error converting year: %v", err)
	}

	month, err := strconv.Atoi(dateTime[5:7])
	if err != nil {
		return "", fmt.Errorf("error converting month: %v", err)
	}

	day, err := strconv.Atoi(dateTime[8:10])
	if err != nil {
		return "", fmt.Errorf("error converting day: %v", err)
	}

	hour, err := strconv.Atoi(dateTime[11:13])
	if err != nil {
		return "", fmt.Errorf("error converting hour: %v", err)
	}

	minute, err := strconv.Atoi(dateTime[14:16])
	if err != nil {
		return "", fmt.Errorf("error converting minute: %v", err)
	}

	second, err := strconv.Atoi(dateTime[17:19])
	if err != nil {
		return "", fmt.Errorf("error converting second: %v", err)
	}

	operation := string(dateTime[19])
	operationHour, err := strconv.Atoi(dateTime[20:22])
	if err != nil {
		return "", fmt.Errorf("error converting operation hour: %v", err)
	}

	operationMinute, err := strconv.Atoi(dateTime[23:25])
	if err != nil {
		return "", fmt.Errorf("error converting operation minute: %v", err)
	}

	// Perform addition or subtraction based on timezone
	if operation == "+" {
		hour += operationHour
		minute += operationMinute
	} else if operation == "-" {
		hour -= operationHour
		minute -= operationMinute
	}

	if minute < 0 {
		minute += 60
		hour--
	} else if minute >= 60 {
		minute -= 60
		hour++
	}

	if hour < 0 {
		hour += 24
		day--
	} else if hour >= 24 {
		hour -= 24
		day++
	}

	if day < 1 {
		day += 31
		month--
	} else if day > 31 {
		day -= 31
		month++
	}

	if month < 1 {
		month += 12
		year--
	} else if month > 12 {
		month -= 12
		year++
	}

	formattedDateTime := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", year, month, day, hour, minute, second)
	return formattedDateTime, nil
}
