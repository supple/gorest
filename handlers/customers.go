package handlers

import (
    "github.com/gin-gonic/gin"
    "encoding/json"
    "github.com/supple/gorest/resources"
    "log"
)

type ApiController struct {

}

type CustomerApi struct {
    ApiController
}

//
func (api *CustomerApi) CampaignGet(c *gin.Context) {

    //var cc *core.CustomerContext = &core.CustomerContext{CustomerName: customerName}
    //cRp := r.NewCustomerRP(cc)

    id, _ := c.Params.Get("id")
    c.JSON(200, map[string]string{"id": id})

    //var id = c.URLParams["id"]
    //obj := app.Storage.Get(id)
    //jsonResponse(w, obj)
}
//
//func CampaignList(c web.C, w http.ResponseWriter, r *http.Request) {
//    names := app.Storage.GetByCriteria()
//    jsonResponse(w, names)
//}
//
func (api *CustomerApi) CampaignPost(c *gin.Context) {
    obj := resources.Customer{}
    decoder := json.NewDecoder(c.Request.Body)
    if err := decoder.Decode(&obj); err != nil {
        log.Print(err.Error())
        c.JSON(422, err)
        //http.Error(w, http.StatusText(422), 422)
        return
    }
    c.JSON(201, obj)
    //app.Storage.Set(obj.Id, &obj)
    //jsonResponse(w, obj)
}
//
//// https://github.com/quintans/goSQL#update
//func CampaignUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    defer r.Body.Close()
//    tmp := make(map[string]interface{})
//    if err := decoder.Decode(&tmp); err != nil {
//        http.Error(w, http.StatusText(422), 422)
//        return
//    }
//
//    var id = c.URLParams["id"]
//    obj := app.Storage.Get(id).(*resources.Device)
//    resources.UpdateModel(obj, tmp)
//
//    app.Storage.Set(id, obj)
//    jsonResponse(w, obj)
//}