package core
import (
    "github.com/supple/gorest/storage"
    "fmt"
)

type AppServices struct {
    Storage *storage.MemStorage
}

type AppError struct {
    Code string
    Message string
}

func (e *AppError) Error() string {
    return fmt.Sprintf("Object not found: %s, value: %s", e.Code, e.Message)
}

type CustomerContext struct {
    ApiKey string
    AppId string
    CustomerName string
}
