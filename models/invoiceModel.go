package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// why are some pointers and others not, because we need to refrence data from other models
type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id" validate="required"`
	Order_id         string             `json:"order_id" validate="required"`
	Payment_method   *string            `json:"payment_method" validate="eq=CARD|eq=CASH|eq="`
	payment_status   *string            `json:"payment_status" validate="required,eq=PENDING|eq=PAID"`
	payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"due_date"`
}
