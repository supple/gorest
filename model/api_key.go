package model

type ApiKey struct {
    CustomerBased `bson:",inline"`
    AppId  string `json:"appId" bson:"appId"`
    ApiKey string `json:"apiKey" bson:"apiKey"`
}
