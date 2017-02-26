package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

type Order struct {
	gorm.Model
	OrderNo         int64
	UserID          uint
	User            User
	OrderType       string
	PaymentAmount   int64
	FromSource      string
	DepositAmout    int64
	PayFrom         string
	AbandonedReason string
	DiscountValue   uint
	TrackingNumber  *string
	//ShippedAt         *time.Time
	ReturnedAt  *time.Time
	CancelledAt *time.Time
	//ShippingAddressID uint
	//ShippingAddress   Address
	//BillingAddressID  uint
	//BillingAddress    Address
	//OrderItems        []OrderItem
	transition.Transition
}
