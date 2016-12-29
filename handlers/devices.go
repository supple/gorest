package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/resources"
)

func (api *DeviceApi) Get(c *gin.Context) {
    var cc = core.GetCC(c)
    rp := resources.NewDeviceRP(cc)
    id := c.Param("id")

    ret, err := rp.FindOne(id)
    if (handleError(err, c)) {
        return
    }
    c.JSON(200, ret)
}

func (api *DeviceApi) Post(c *gin.Context) {
    var err error

    obj, err := resources.DeviceFromJson(c.Request.Body)
    if err != nil {
        c.JSON(422, err)
        return
    }

    dRp := resources.NewDeviceRP(core.GetCC(c))
    err = dRp.Create(obj)
    if (handleError(err, c)) {
        return
    }

    c.JSON(201, obj)
}

func (api *DeviceApi) Patch(c *gin.Context) {
    var err error

    obj, err := resources.MapFromJson(c.Request.Body)
    id := c.Param("id")
    if err != nil {
        c.JSON(422, err)
        return
    }
    dRp := resources.NewDeviceRP(core.GetCC(c))
    err = dRp.Update(id, obj)
    if (handleError(err, c)) {
        return
    }

    c.JSON(200, obj)
}


