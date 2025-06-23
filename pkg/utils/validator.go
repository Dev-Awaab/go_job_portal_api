package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()


func ValidateRequest(req interface{}) error {
	if err := validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var messages []string
			for _, fieldErr := range ve {
				msg := fmt.Sprintf("Field '%s' failed on the '%s' rule", fieldErr.Field(), fieldErr.Tag())
				messages = append(messages, msg)
			}
			return fmt.Errorf("validation failed: %v", messages)
		}
		return err
	}
	return nil
}
