package core

import (
    "runtime"
    "log"
)

type APIErrors struct {
    Errors      []*APIError `json:"errors"`
}

func (ers *APIErrors) Status() int {
    return ers.Errors[0].Status
}

type APIError struct {
    Status  int         `json:"status"`
    Code    string      `json:"code"`
    Message string      `json:"title"`
    Details string      `json:"details"`
    Href    string      `json:"href"`
}

func (e *APIError) Error() string {
    return e.Message
}

var (
    ErrDatabase         = newAPIError(500, "database_error", "Database Error", "An unknown error occurred.", "")
    ErrInvalidSet       = newAPIError(404, "invalid_set", "Invalid Set", "The set you requested does not exist.", "")
    ErrInvalidGroup     = newAPIError(404, "invalid_group", "Invalid Group", "The group you requested does not exist.", "")
    ErrUnknown     = newAPIError(500, "unknown_error", "Unknown error", "", "")
    ErrInvalidApiKey     = newAPIError(401, "invalid_api_key", "Invalid api key", "", "")

)

func newAPIError(status int, code string, message string, details string, href string) *APIError {
    return &APIError{
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
    case *APIError:
        apiError := error.(*APIError)
        apiErrors = &APIErrors{
            Errors: []*APIError{apiError},
        }
    case *APIErrors:
        apiErrors = error.(*APIErrors)
    default:
        apiErrors = &APIErrors{
            Errors: []*APIError{ErrUnknown},
        }
    }
    return apiErrors.Status(), apiErrors
}

