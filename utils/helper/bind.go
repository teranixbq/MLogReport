package helper

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)


func BindJSON(c *gin.Context, input interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(input); err != nil {
		if err.Error() == "EOF" {
			return errors.New("error: empty request")
		}
		return errors.New("error: input name property is incorrect")
	}

	err := convertToLowerCase(input)
	if err != nil {
		return err
	}

	return nil
}

func convertToLowerCase(input interface{}) error {
	val := reflect.ValueOf(input).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			val.Field(i).SetString(strings.ToLower(field.String()))
		}
	}
	return nil
}