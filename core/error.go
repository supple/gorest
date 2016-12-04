package core

import (
    "runtime"
    "log"
    "fmt"
)


type ErrObjectNotFound struct {
    Object string
    Value  string
}
func (e *ErrObjectNotFound) Error() string {
    return fmt.Sprintf("Object not found: %s, value: %s", e.Object, e.Value)
}

type APIErrors struct {
    Errors      []*ApiError `json:"errors"`
}

func (ers *APIErrors) Status() int {
    return ers.Errors[0].Status
}

type ApiError struct {
    Status  int         `json:"status"`
    Code    string      `json:"code"`
    Message string      `json:"title"`
    Details string      `json:"details"`
    Href    string      `json:"href"`
}

func (e *ApiError) Error() string {
    return e.Message
}

var (
    ErrDatabase = NewAPIError(503, "database_error", "Database Error", "Temporary server error.", "")
    Err404 = NewAPIError(404, "object_not_found", "Object not found", "Requested object not found.", "")
    ErrUnknown = NewAPIError(500, "unknown_error", "Unknown error", "", "")
    ErrInvalidApiKey = NewAPIError(401, "invalid_api_key", "Invalid api key", "", "")
)

func NewAPIError(status int, code string, message string, details string, href string) *ApiError {
    return &ApiError{
        Status:     status,
        Code:       code,
        Message:      message,
        Details:    details,
        Href:       href,
    }
}

func ErrorMessage(err error, error interface{}) (int, *APIErrors) {
    var apiErrors *APIErrors

    // This the best way to log?
    trace := make([]byte, 1024)
    runtime.Stack(trace, true)
    log.Printf("ERROR: %s\n%s", err, trace)

    switch error.(type) {
    case *ApiError:
        apiError := error.(*ApiError)
        apiErrors = &APIErrors{
            Errors: []*ApiError{apiError},
        }
    case *APIErrors:
        apiErrors = error.(*APIErrors)
    default:
        apiErrors = &APIErrors{
            Errors: []*ApiError{ErrUnknown},
        }
    }
    return apiErrors.Status(), apiErrors
}

