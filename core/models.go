package core

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

type Storage interface {
    Get(id string) interface{}
    Set(id string, obj interface{}) bool
    Update(id string,obj interface{})
}

type AppServices struct {
    Storage Storage
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
