package main

import (
    "runtime"
    "log"
)

type APIErrors struct {
    Errors      []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
    return errors.Errors[0].Status
}

type APIError struct {
    Status      int         `json:"status"`
    Code        string      `json:"code"`
    Title       string      `json:"title"`
    Details     string      `json:"details"`
    Href        string      `json:"href"`
}

var errUnknown = &APIError{Title: "unknown", Code: 500}

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
            Errors: []*APIError{errUnknown},
        }
    }
    return apiErrors.Status(), apiErrors
}

