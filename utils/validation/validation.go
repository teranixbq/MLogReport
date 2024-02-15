package validation

import (
	"errors"
	"strings"
	"time"
)

func CheckEqual(data string, validData []string) (string, error) {
	input := strings.ToLower(data)

	isValid := false
	for _, valid := range validData {
		if input == strings.ToLower(valid) {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", errors.New("error : data not match")
	}

	return input, nil
}

func CheckLength(data string) error {
	if len(data) <= 8 {
		return errors.New("error : password minimum 8 characters")
	}

	return nil
}

func LimitDescription(data string, limit int) error {
	clean := strings.ReplaceAll(data, " ", "")
	descLimit := strings.Count(clean, "")
	if descLimit > limit {
		return errors.New("error : karakter input melebihi batas")
	}

	return nil
}

func TimeUpdate(data time.Time) error {
	timeAsia, errTime := time.LoadLocation("Asia/Bangkok")
	if errTime != nil {
		return errTime
	}
	day := time.Now().In(timeAsia)

	if data.Before(day) {
		return errors.New("error : update time has expired")
	}

	return nil
}
