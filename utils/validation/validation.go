package validation

import (
	"errors"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func CheckEmpty(data ...interface{}) error {
	for _, d := range data {
		v := reflect.ValueOf(d)
		switch v.Kind() {
		case reflect.Ptr:
			if v.IsNil() {
				return errors.New("error : data cannot be empty")
			}
		case reflect.String:
			if d == "" {
				return errors.New("error : data cannot be empty")
			}
		case reflect.Slice:
			if v.Len() == 0 {
				return errors.New("error: slice cannot be empty")
			}
		}
	}
	return nil
}

func CheckAllEmpty(data ...interface{}) error {
	allEmpty := true
	for _, d := range data {
		v := reflect.ValueOf(d)
		switch v.Kind() {
		case reflect.String:
			if d != "" {
				allEmpty = false
			}
		case reflect.Struct:
			if !v.IsZero() {
				allEmpty = false
			}
		}
	}
	if allEmpty {
		return errors.New("error : all data cannot be empty")
	}
	return nil
}

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

func TimeAsia() *time.Location {
	timeAsia, _ := time.LoadLocation("Asia/Bangkok")
	return timeAsia
}

func DateAsia() string {
	timesAsia := TimeAsia()
	timepath := time.Now().In(timesAsia).Format("2006-01-02")
	return timepath
}

func CheckPagination(page, limit int) (int, int, error) {
	numericRegex := regexp.MustCompile(`^\d+$`)

	if page != 0 {
		if !numericRegex.MatchString(strconv.Itoa(page)) {
			return 0, 0, errors.New("error: page must be a numeric value")
		}
	}

	if limit != 0 {
		if !numericRegex.MatchString(strconv.Itoa(limit)) {
			return 0, 0, errors.New("error: limit must be a numeric value")
		}
	}

	if limit != 0 && limit != 5 && limit != 10 {
		return 0, 0, errors.New("error: limit must be either 5 or 10")
	}

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	return page, limit, nil
}

func EnvCheck(config interface{}, fileENV ...string) {
	godotenv.Load(fileENV...)

	configType := reflect.TypeOf(config).Elem()
	configValue := reflect.ValueOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		envKey := field.Tag.Get("env")

		if envKey == "" {
			log.Fatalf("Config: env tag not found for field %s", field.Name)
		}

		envValue, found := os.LookupEnv(envKey)
		if !found {
			log.Fatalf("Config: %s environment variable not found", envKey)
		}

		fieldValue := configValue.Field(i)
		kind := fieldValue.Kind()

		switch kind {
			case reflect.String :
				fieldValue.SetString(envValue)
			case reflect.Int :
				v, err := strconv.Atoi(envValue)
				if err != nil {
					logrus.Fatalf("Config: invalid %s value, %s", envKey, err.Error())
				}
				fieldValue.SetInt(int64(v))
			default:
				logrus.Fatalf("Config: unsupported field type %s for %s", fieldValue.Kind(), envKey)
		}
	}
}