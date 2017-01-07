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

type ValidationErrors struct {
    errors   []ValidationError `json:"errors"`
}
func (ers ValidationErrors) Error() string {
    b, err := json.Marshal(ers)
    if err != nil {
        return ""
    } else {
        return string(b)
    }
}
