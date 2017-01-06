package model


type Customer struct {
    CustomerBased `bson:",inline"`
    Hash string `json:"hash" bson:"hash"`
}

