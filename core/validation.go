package core

import "encoding/json"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (er ValidationError) Error() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func NewError(field string, message string) ValidationError {
	return ValidationError{Field: field, Message: message}
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (ers ValidationErrors) Error() string {
	b, err := json.Marshal(ers)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
