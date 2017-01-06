package model

type Device struct {
    CustomerBased `bson:",inline"`
    AppId      string        `json:"appId" bson:"appId" `
    AppToken   string        `json:"appToken" bson:"appToken"`
    AppVersion string         `json:"appVersion" bson:"appVersion"`
}

func (d *Device) IsValidForUpdate() {

}

func NewDevice() *Device {
    return &Device{}
}