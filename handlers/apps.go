package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/resources"
)

type AppApi struct {
    ApiController
}

func (api *AppApi) Get(c *gin.Context) {
    var cc = core.GetCC(c)
    rp := resources.NewAppRP(cc)
    id := c.Param("id")
    ret, err := rp.FindOne(id)
    if (handleError(err, c)) {
        return
    }

    if (c.Query("fields") == "all") {
        c.JSON(200, ret)
        return
    }

    fields := []string{"id", "name", "createdAt", "updatedAt"}
    c.JSON(200, core.PartialStruct(*ret, fields...))
}

func (api *AppApi) Create(c *gin.Context) {
    var err error

    obj, err := resources.AppFromJson(c.Request.Body)
    if err != nil {
        c.JSON(422, err)
        return
    }

    // validate object
    errors := resources.ValidateApp(obj)
    if errors != nil {
        c.JSON(422, errors)
        return
    }

    dRp := resources.NewAppRP(core.GetCC(c))
    err = dRp.Create(obj)
    if (handleError(err, c)) {
        return
    }

    c.JSON(201, obj)
}

func (api *AppApi) Update(c *gin.Context) {
    var err error

    //obj, err := resources.MapFromJson(c.Request.Body)
    obj, err := resources.AppFromJson(c.Request.Body)
    id := c.Param("id")
    if err != nil {
        c.JSON(422, err)
        return
    }

    dRp := resources.NewAppRP(core.GetCC(c))
    err = dRp.Update(id, obj)
    if (handleError(err, c)) {
        return
    }

    c.JSON(200, obj)
}

func (api *AppApi) List(c *gin.Context) {

}

