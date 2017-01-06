package model

type App struct {
    CustomerBased `bson:",inline"`
    Os       string `json:"os" bson:"os"`
    Name     string `json:"name" bson:"name"`
    GcmToken string `json:"gcmToken" bson:"gcmToken"`
    ApnsAuthKey string `json:"apnsAuthKey" bson:"apnsAuthKey"`
    ApnsTeamId string `json:"apnsTeamId" bson:"apnsTeamId"`
    ApnsKeyId string `json:"apnsKeyId" bson:"apnsKeyId"`
}

