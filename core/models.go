package core
import (
    "github.com/supple/gorest/storage"
    "fmt"
    "github.com/gin-gonic/gin"
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
    DeviceId string
    CustomerName string
}

func GetCC(c *gin.Context) *CustomerContext {
    cc, ok := c.Get("cc")
    if ok {
        return cc.(*CustomerContext)
    }
    panic("Customer context not set")
}
