package model

type Properties map[string]string

type Location struct {
	Type string     `json:"type"`
	Data Properties `json:"data"`
}

type Device struct {
	CustomerBased `bson:",inline"`
	AppId         string      `json:"appId" bson:"appId" `
	AppToken      string      `json:"appToken" bson:"appToken"`
	AppVersion    string      `json:"appVersion" bson:"appVersion"`
	Properties    Properties  `json:"properties" bson:"properties"`
	Locations     []*Location `json:"locations" bson:"locations"`
}

func (d *Device) IsValidForUpdate() {

}

func (d *Device) GetProperty() {

}

func NewDevice() *Device {
	return &Device{}
}
