package validation

import (
	"errors"
	"strings"
)

func CheckEqual(data string, validData []string) (string,error){
	input := strings.ToLower(data)
	
	isValid := false 
	for _,valid := range validData {
		if input == strings.ToLower(valid){
			isValid = true
			break
		}
	}

	if !isValid {
		return "",errors.New("error : data not match")
	}

	return input,nil
}

func CheckLength(data string) error{
	if len(data) <= 8 {
		return errors.New("error : password minimum 8 characters")
	}
	
	return nil
}