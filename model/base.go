package model

import "github.com/supple/gorest/core"

// Customer based object
type CustomerBased struct {
    Id           string `json:"id,omitempty" bson:"_id"`
    CustomerName string `json:"customerName" bson:"customerName,omitempty" validate:"required"`
    CreatedAt    string  `json:"createdAt" bson:"createdAt,omitempty"`
    UpdatedAt    string  `json:"updatedAt" bson:"updatedAt,omitempty"`
    DeletedAt    string  `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

// Set values on common model fields
func (model *CustomerBased) SetBasicFields() {
    model.CreatedAt = core.GetJodaTime()
    model.UpdatedAt = model.CreatedAt
}

